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

package list

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	"github.com/radius-project/radius/pkg/cli/clients"
	types "github.com/radius-project/radius/pkg/cli/cmd/recipe"
	"github.com/radius-project/radius/pkg/cli/connections"
	"github.com/radius-project/radius/pkg/cli/framework"
	"github.com/radius-project/radius/pkg/cli/objectformats"
	"github.com/radius-project/radius/pkg/cli/output"
	"github.com/radius-project/radius/pkg/cli/workspaces"
	"github.com/radius-project/radius/pkg/corerp/api/v20231001preview"
	ds_ctrl "github.com/radius-project/radius/pkg/datastoresrp/frontend/controller"
	"github.com/radius-project/radius/pkg/recipes"
	"github.com/radius-project/radius/pkg/to"
	"github.com/radius-project/radius/test/radcli"
	"github.com/stretchr/testify/require"
)

func Test_CommandValidation(t *testing.T) {
	radcli.SharedCommandValidation(t, NewCommand)
}

func Test_Validate(t *testing.T) {
	configWithWorkspace := radcli.LoadConfigWithWorkspace(t)
	testcases := []radcli.ValidateInput{
		{
			Name:          "Valid List Command",
			Input:         []string{},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         configWithWorkspace,
			},
		},
		{
			Name:          "List Command with fallback workspace",
			Input:         []string{"-e", "my-env", "-g", "my-env"},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         radcli.LoadEmptyConfig(t),
			},
		},
		{
			Name:          "List Command with too many args",
			Input:         []string{"foo", "bar"},
			ExpectedValid: false,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         configWithWorkspace,
			},
		},
	}
	radcli.SharedValidateValidation(t, NewCommand, testcases)
}

func Test_Run(t *testing.T) {
	t.Run("List recipes linked to the environment - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		envResource := v20231001preview.EnvironmentResource{
			ID:       to.Ptr("/planes/radius/local/resourcegroups/kind-kind/providers/applications.core/environments/kind-kind"),
			Name:     to.Ptr("kind-kind"),
			Type:     to.Ptr("applications.core/environments"),
			Location: to.Ptr(v1.LocationGlobal),
			Properties: &v20231001preview.EnvironmentProperties{
				Recipes: map[string]map[string]v20231001preview.RecipePropertiesClassification{
					ds_ctrl.MongoDatabasesResourceType: {
						"cosmosDB": &v20231001preview.BicepRecipeProperties{
							TemplateKind: to.Ptr(recipes.TemplateKindBicep),
							TemplatePath: to.Ptr("ghcr.io/testpublicrecipe/bicep/modules/mongodatabases:v1"),
						},
						"cosmosDB-terraform": &v20231001preview.TerraformRecipeProperties{
							TemplateKind:    to.Ptr(recipes.TemplateKindTerraform),
							TemplatePath:    to.Ptr("Azure/cosmosdb/azurerm"),
							TemplateVersion: to.Ptr("1.1.0"),
						},
					},
				},
			},
		}
		recipes := []types.EnvironmentRecipe{
			{
				Name:         "cosmosDB",
				ResourceType: ds_ctrl.MongoDatabasesResourceType,
				TemplateKind: recipes.TemplateKindBicep,
				TemplatePath: "ghcr.io/testpublicrecipe/bicep/modules/mongodatabases:v1",
			},
			{
				Name:            "cosmosDB-terraform",
				ResourceType:    ds_ctrl.MongoDatabasesResourceType,
				TemplateKind:    recipes.TemplateKindTerraform,
				TemplatePath:    "Azure/cosmosdb/azurerm",
				TemplateVersion: "1.1.0",
			},
		}

		appManagementClient := clients.NewMockApplicationsManagementClient(ctrl)
		appManagementClient.EXPECT().
			GetEnvDetails(gomock.Any(), gomock.Any()).
			Return(envResource, nil).Times(1)

		outputSink := &output.MockOutput{}

		runner := &Runner{
			ConnectionFactory: &connections.MockFactory{ApplicationsManagementClient: appManagementClient},
			Output:            outputSink,
			Workspace:         &workspaces.Workspace{},
			Format:            "table",
		}

		err := runner.Run(context.Background())
		require.NoError(t, err)

		expected := []any{
			output.FormattedOutput{
				Format:  "table",
				Obj:     recipes,
				Options: objectformats.GetEnvironmentRecipesTableFormat(),
			},
		}
		require.Equal(t, expected, outputSink.Writes)
	})
}
