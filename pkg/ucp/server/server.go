/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	hostOpts "github.com/radius-project/radius/pkg/armrpc/hostoptions"
	"github.com/radius-project/radius/pkg/kubeutil"
	metricsprovider "github.com/radius-project/radius/pkg/metrics/provider"
	metricsservice "github.com/radius-project/radius/pkg/metrics/service"
	profilerprovider "github.com/radius-project/radius/pkg/profiler/provider"
	profilerservice "github.com/radius-project/radius/pkg/profiler/service"
	"github.com/radius-project/radius/pkg/sdk"
	"github.com/radius-project/radius/pkg/trace"
	"github.com/radius-project/radius/pkg/ucp/backend"
	"github.com/radius-project/radius/pkg/ucp/config"
	"github.com/radius-project/radius/pkg/ucp/data"
	"github.com/radius-project/radius/pkg/ucp/dataprovider"
	"github.com/radius-project/radius/pkg/ucp/frontend/api"
	"github.com/radius-project/radius/pkg/ucp/hosting"
	"github.com/radius-project/radius/pkg/ucp/hostoptions"
	qprovider "github.com/radius-project/radius/pkg/ucp/queue/provider"
	"github.com/radius-project/radius/pkg/ucp/rest"
	"github.com/radius-project/radius/pkg/ucp/secret/provider"
	"github.com/radius-project/radius/pkg/ucp/ucplog"

	kube_rest "k8s.io/client-go/rest"
)

const (
	HTTPServerStopTimeout = time.Second * 10
	ServiceName           = "ucp"
)

type Options struct {
	Config                  *hostoptions.UCPConfig
	Port                    string
	StorageProviderOptions  dataprovider.StorageProviderOptions
	LoggingOptions          ucplog.LoggingOptions
	SecretProviderOptions   provider.SecretProviderOptions
	QueueProviderOptions    qprovider.QueueProviderOptions
	MetricsProviderOptions  metricsprovider.MetricsProviderOptions
	ProfilerProviderOptions profilerprovider.ProfilerProviderOptions
	TracerProviderOptions   trace.Options
	TLSCertDir              string
	PathBase                string
	InitialPlanes           []rest.Plane
	Identity                hostoptions.Identity
	UCPConnection           sdk.Connection
	Location                string
}

const UCPProviderName = "ucp"

// NewServerOptionsFromEnvironment creates a new Options struct from environment variables and returns it along with any errors.
func NewServerOptionsFromEnvironment() (Options, error) {
	basePath, ok := os.LookupEnv("BASE_PATH")
	if ok && len(basePath) > 0 && (!strings.HasPrefix(basePath, "/") || strings.HasSuffix(basePath, "/")) {
		return Options{}, errors.New("env: BASE_PATH must begin with '/' and must not end with '/'")
	}

	tlsCertDir := os.Getenv("TLS_CERT_DIR")
	ucpConfigFile := os.Getenv("UCP_CONFIG")

	port := os.Getenv("PORT")
	if port == "" {
		return Options{}, errors.New("UCP Port number must be set")
	}

	opts, err := hostoptions.NewHostOptionsFromEnvironment(ucpConfigFile)
	if err != nil {
		return Options{}, err
	}

	storeOpts := opts.Config.StorageProvider
	planes := opts.Config.Planes
	secretOpts := opts.Config.SecretProvider
	qproviderOpts := opts.Config.QueueProvider
	metricsOpts := opts.Config.MetricsProvider
	traceOpts := opts.Config.TracerProvider
	profilerOpts := opts.Config.ProfilerProvider
	loggingOpts := opts.Config.Logging
	identity := opts.Config.Identity
	// Set the default authentication method if AuthMethod is not set.
	if identity.AuthMethod == "" {
		identity.AuthMethod = hostoptions.AuthDefault
	}

	location := opts.Config.Location
	if location == "" {
		location = "global"
	}

	var cfg *kube_rest.Config
	if opts.Config.UCP.Kind == config.UCPConnectionKindKubernetes {
		cfg, err = kubeutil.NewClientConfig(&kubeutil.ConfigOptions{
			// TODO: Allow to use custom context via configuration. - https://github.com/radius-project/radius/issues/5433
			ContextName: "",
			QPS:         kubeutil.DefaultServerQPS,
			Burst:       kubeutil.DefaultServerBurst,
		})
		if err != nil {
			return Options{}, fmt.Errorf("failed to get kubernetes config: %w", err)
		}
	}

	ucpConn, err := config.NewConnectionFromUCPConfig(&opts.Config.UCP, cfg)
	if err != nil {
		return Options{}, err
	}

	return Options{
		Config:                  opts.Config,
		Port:                    port,
		TLSCertDir:              tlsCertDir,
		PathBase:                basePath,
		StorageProviderOptions:  storeOpts,
		SecretProviderOptions:   secretOpts,
		QueueProviderOptions:    qproviderOpts,
		MetricsProviderOptions:  metricsOpts,
		TracerProviderOptions:   traceOpts,
		ProfilerProviderOptions: profilerOpts,
		LoggingOptions:          loggingOpts,
		InitialPlanes:           planes,
		Identity:                identity,
		UCPConnection:           ucpConn,
		Location:                location,
	}, nil
}

// NewServer creates a new hosting.Host instance with services for API, EmbeddedETCD, Metrics, Profiler and Backend (if
// enabled) based on the given Options.
func NewServer(options *Options) (*hosting.Host, error) {
	hostingServices := []hosting.Service{
		api.NewService(api.ServiceOptions{
			ProviderName:           UCPProviderName,
			Address:                ":" + options.Port,
			PathBase:               options.PathBase,
			Config:                 options.Config,
			Location:               options.Location,
			TLSCertDir:             options.TLSCertDir,
			StorageProviderOptions: options.StorageProviderOptions,
			SecretProviderOptions:  options.SecretProviderOptions,
			QueueProviderOptions:   options.QueueProviderOptions,
			InitialPlanes:          options.InitialPlanes,
			Identity:               options.Identity,
			UCPConnection:          options.UCPConnection,
		}),
	}

	if options.StorageProviderOptions.Provider == dataprovider.TypeETCD &&
		options.StorageProviderOptions.ETCD.InMemory {
		hostingServices = append(hostingServices, data.NewEmbeddedETCDService(data.EmbeddedETCDServiceOptions{ClientConfigSink: options.StorageProviderOptions.ETCD.Client}))
	}

	options.MetricsProviderOptions.ServiceName = ServiceName
	if options.MetricsProviderOptions.Prometheus.Enabled {
		metricOptions := metricsservice.HostOptions{
			Config: &options.MetricsProviderOptions,
		}
		hostingServices = append(hostingServices, metricsservice.NewService(metricOptions))
	}

	if options.ProfilerProviderOptions.Enabled {
		profilerOptions := profilerservice.HostOptions{
			Config: &options.ProfilerProviderOptions,
		}
		hostingServices = append(hostingServices, profilerservice.NewService(profilerOptions))
	}

	backendServiceOptions := hostOpts.HostOptions{
		Config: &hostOpts.ProviderConfig{
			StorageProvider:  options.StorageProviderOptions,
			SecretProvider:   options.SecretProviderOptions,
			QueueProvider:    options.QueueProviderOptions,
			MetricsProvider:  options.MetricsProviderOptions,
			TracerProvider:   options.TracerProviderOptions,
			ProfilerProvider: options.ProfilerProviderOptions,
		},
	}
	hostingServices = append(hostingServices, backend.NewService(backendServiceOptions))

	options.TracerProviderOptions.ServiceName = "ucp"
	hostingServices = append(hostingServices, &trace.Service{Options: options.TracerProviderOptions})

	return &hosting.Host{
		Services: hostingServices,
	}, nil
}
