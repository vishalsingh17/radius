// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package datamodel

// ExtensionKind
type ExtensionKind string

const (
	ManualScaling                ExtensionKind = "manualScaling"
	DaprSidecar                  ExtensionKind = "daprSidecar"
	KubernetesMetadata           ExtensionKind = "kubernetesMetadata"
	KubernetesNamespaceExtension ExtensionKind = "kubernetesNamespace"
)

// Extension of a resource.
type Extension struct {
	Kind                ExtensionKind           `json:"kind,omitempty"`
	ManualScaling       *ManualScalingExtension `json:"manualScaling,omitempty"`
	DaprSidecar         *DaprSidecarExtension   `json:"daprSidecar,omitempty"`
	KubernetesMetadata  *KubeMetadataExtension  `json:"kubernetesMetadata,omitempty"`
	KubernetesNamespace *KubeNamespaceExtension `json:"kubernetesNamespace,omitempty"`
}

// KubeMetadataExtension represents the extension of kubernetes resource.
type KubeMetadataExtension struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// KubeNamespaceOverrideExtension represents the extension to override kubernetes namespace.
type KubeNamespaceExtension struct {
	Namespace string `json:"namespace,omitempty"`
}

// FindExtension finds the extension.
func FindExtension(exts []Extension, kind ExtensionKind) *Extension {
	for _, ext := range exts {
		if ext.Kind == kind {
			return &ext
		}
	}
	return nil
}
