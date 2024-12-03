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

package resource

import (
	"context"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type ResourceName struct{}

// Each resource has an input struct, defining what arguments it accepts.
type Args struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but it's generally a
	// good idea.
	Length int `pulumi:"length"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type State struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	Args
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minimum.
func (ResourceName) Create(ctx context.Context, name string, input Args, preview bool) (string, State, error) {
	state := State{Args: input, Result: "some text to override later"}

	if preview {
		// don't do any actual API calls for preview
		return name, state, nil
	}

	panic("not implemented")

	return name, state, nil
}

// Don't forget to add this to the registry:
/*
providerRegistryEntry{
		PackageName:      "resource",
		// index if the SDK root should expose this, the name of the scope otherwise.
		// if scope is pqr, the import will be:
		// import * as fdprovider from @alphonsocode/pulumi-provider-fd-provider
		// fdprovider.pqr.Resource
		Scope:            "index",
		Kind:             ProviderKindResource,
		InferredResource: infer.Resource[resource.Resource, resource.Args, resource.State](),
	},
*/

/*
func (AccountReady) Update(ctx context.Context, name string, old State, new Args, preview bool) (State, error) {
	state := State{Args: new, Ready: true}

	if preview {
		return state, nil
	}

	panic("Not Implemented")

	return state, nil
}
*/
