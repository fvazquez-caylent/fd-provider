package scope

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Similar to resources, components have an input struct, defining what arguments it accepts.
type RandomFunction struct {
}

type RandomFunctionArgs struct {
	Length pulumi.IntInput `pulumi:"length"`
}

type RandomFunctionOutput struct {
	Length pulumi.Output `pulumi:"out"`
}

func (f *RandomFunction) Call(ctx context.Context, input RandomFunctionArgs) (*RandomFunctionOutput, error) {
	panic("not implemented")
}
