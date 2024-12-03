package component

// Usually you don't have to edit this

import (
	"github.com/pkg/errors"

	providerconfig "github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/providerconfig"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// You should not have to change this file
// unless you need low level modifications to how a provider is handled.

type ComponentName struct{}

func (s *ComponentName) Construct(
	ctx *pulumi.Context,
	name string,
	typ string,
	args Args,
	options pulumi.ResourceOption,
) (*State, error) {
	resource := &State{
		Args: args,
	}
	err := ctx.RegisterComponentResource(typ, name, resource, options)
	if err != nil {
		return nil, errors.Wrapf(err, "error registering resource")
	}

	config := infer.GetConfig[providerconfig.Config](ctx.Context())

	err = resource.Create(ctx, name, args, config, options)
	if err != nil {
		return nil, errors.Wrapf(err, "error initializing resource")
	}

	return resource, nil
}
