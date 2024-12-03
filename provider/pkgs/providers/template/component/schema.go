package component

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Args struct {
	// use pulumi.Input type, not normal types
	Input pulumi.StringInput `pulumi:"input"`

	// use PtrInput type for optional inputs
	// NOTE: Here, OptionalInput itself may be nil
	// and the resolved output may also be nil
	// if you are trying to apply this resource, be careful to check both times!
	OptionalInput pulumi.StringPtrInput `pulumi:"optionalInput,optional"`
}

type State struct {
	pulumi.ResourceState

	// Adding the inputs as an output is generally a good idea
	// but not necessary
	Args
}
