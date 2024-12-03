// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"

	"github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/metadata"
	"github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/providerconfig"
	"github.com/fvazquez-caylent/fd-provider/provider/pkgs/core/registry"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
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
