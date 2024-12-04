package provider

import (
	"fmt"

	"github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/metadata"
	"github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/providerconfig"
	"github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/registry"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/resources"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "fd-provider"

func Provider() p.Provider {
	inferOptions := infer.Options{
		Resources:  []infer.InferredResource{},
		Components: []infer.InferredComponent{},
		Config:     infer.Config[providerconfig.Config](),
		ModuleMap:  map[tokens.ModuleName]tokens.ModuleName{},
		Functions:  []infer.InferredFunction{},
		Metadata:   metadata.Metadata,
	}

	// Register S3 resource
	inferOptions.Resources = append(inferOptions.Resources, infer.InferredResource{
		Name:    "s3.Bucket",
		Factory: resources.NewS3Bucket,
	})

	for _, entry := range registry.Registry {
		switch entry.Kind {
		case registry.ProviderKindComponent:
			{
				inferOptions.Components = append(inferOptions.Components, entry.InferredComponent)
			}
		case registry.ProviderKindFunction:
			{
				inferOptions.Functions = append(inferOptions.Functions, entry.InferredFunction)
			}
		case registry.ProviderKindResource:
			{
				inferOptions.Resources = append(inferOptions.Resources, entry.InferredResource)
			}
		default:
			{
				panic(fmt.Errorf("error: unknown kind %s", entry.Kind))
			}
		}

		if entry.PackageName != entry.Scope {
			inferOptions.ModuleMap[tokens.ModuleName(entry.PackageName)] = tokens.ModuleName(entry.Scope)
		}
	}
	return infer.Provider(inferOptions)
}
