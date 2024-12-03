package component

import (
	providerconfig "github.com/AlphonsoCode/pulumi-provider-xyzqw/provider/pkgs/core/providerconfig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func (component *State) Create(ctx *pulumi.Context, name string, args Args, configs providerconfig.Config, opts pulumi.ResourceOption) error {

	// don't forget to assign the final state to the component this is defined on

	panic("not implemented")

	return nil
}
