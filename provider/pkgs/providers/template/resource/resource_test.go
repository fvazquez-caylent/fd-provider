package resource_test

import (
	"testing"

	provider "github.com/AlphonsoCode/pulumi-provider-xyzqw/provider"
	"github.com/blang/semver"
	integration "github.com/pulumi/pulumi-go-provider/integration"
	presource "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {
	server := integration.NewServer("xyzqw", semver.Version{Minor: 1}, provider.Provider())
	integration.LifeCycleTest{
		Resource: "xyzqw:index:Random",
		Create: integration.Operation{
			Inputs: presource.NewPropertyMapFromMap(map[string]interface{}{
				"length": 24,
			}),
			Hook: func(inputs, output presource.PropertyMap) {
				t.Logf("Outputs: %v", output)
				result := output["result"].StringValue()
				assert.Len(t, result, 24)
			},
		},
		Updates: []integration.Operation{
			{
				Inputs: presource.NewPropertyMapFromMap(map[string]interface{}{
					"length": 10,
				}),
				Hook: func(inputs, output presource.PropertyMap) {
					result := output["result"].StringValue()
					assert.Len(t, result, 10)
				},
			},
		},
	}.Run(t, server)
}
