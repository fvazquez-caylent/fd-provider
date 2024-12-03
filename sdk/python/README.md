# Pulumi Provider Boilerplate

This repository is a boilerplate showing how to create and locally test a native Pulumi provider for Alphonso.

## Authoring a Pulumi Native Provider

### Prerequisites

You will need to ensure the following tools are installed and present in your `$PATH`:

* [`pulumictl`](https://github.com/pulumi/pulumictl#installation)
* [Go 1.21](https://golang.org/dl/) or 1.latest
* [NodeJS](https://nodejs.org/en/) 14.x. or greater
* [Python](https://www.python.org/downloads/) (called as `python3`).  For recent versions of MacOS, the system-installed version is fine.

### Creating a new provider repository (One Time Setup)

1. Click "Use this template" on this repository (or ask DevOps to create a fork for you)
1. Edit the Makefile and modify the constants at the top. Commit the code after you are done with this step.
1. Run the following command:

    ```bash
    make prepare
    ```

   This will do the following:
   - rename folders in `provider/cmd` to `pulumi-resource-{NAME}`
   - replace dependencies in `provider/go.mod` to reflect your repository name
   - find and replace all instances of the boilerplate `fd-provider` with the `NAME` of your provider.
   - find and replace all instances of the boilerplate `abc` with the `ORG` of your provider.
   - replace all instances of the `github.com/fvazquez-caylent/fd-provider` repository with the `REPOSITORY` location

   Commit after you are done

#### Creating a Component Resource

1. Decide a package scope for your project. The `<scope>` here will be `index` if you want your resource to be available at the top level of your sdk, and will be the subpackage name otherwise.
   For example, if `scope` is `pqr` and the component name is Abc, in nodejs you will use the resource as:
   ```typescript
   import * as fdprovider from "@alphonsocode/pulumi-provider-fd-provider";
   const res = new fdprovider.pqr.Abc("name", {...args});
   ```
1. Run `make create_component SCOPE="<scope>" RESOURCE="resource package name" RESOURCE_STRUCT="resource struct name"`
1. Add your component resource to the registry at `provider/pkgs/core/registry`
   ```
   ProviderRegistryEntry{
      PackageName:       "<your component go package name>",
      Scope:             "<scope>",
      Kind:              ProviderKindComponent,
      InferredComponent: infer.Component[*package.ComponentName, package.Args, *package.State](),
   },
   ```
1. You will probably not have to change anything in `component.go` here.
1. Edit the schema for your component (inputs and state) in `schema.go`
1. Edit the code to create your component in `create.go`

#### Creating a Native Provider

1. Decide a package scope for your project. The `<scope>` here will be `index` if you want your resource to be available at the top level of your sdk, and will be the subpackage name otherwise.
   For example, if `scope` is `pqr` and the component name is Abc, in nodejs you will use the resource as:
   ```typescript
   import * as fdprovider from "@alphonsocode/pulumi-provider-fd-provider";
   const res = new fdprovider.pqr.Abc("name", {...args});
   ```
1. Run `make create_native SCOPE="<scope>" RESOURCE="resource package name" RESOURCE_STRUCT="resource struct name"`
1. Add your native provider to the registry at `provider/pkgs/core/registry`
   ```
   ProviderRegistryEntry{
      PackageName:       "<your resource go package name>",
      Scope:             "<scope>",
      Kind:              ProviderKindResource,
      InferredResource: infer.Resource[package.ResourceName, package.Args, package.State](),
   },
   ```
1. Implement the create method at the very least.


#### Build the provider and install the plugin

   ```bash
   $ make build install
   ```
   
This will:

1. Create the SDK codegen binary and place it in a `./bin` folder (gitignored)
2. Create the provider binary and place it in the `./bin` folder (gitignored)
3. Generate the Go, Node, and Python SDKs and place them in the `./sdk` folder
4. Install the provider on your machine (using npm link)

### Useful Shortcuts

The default make command automatically builds the program and examples. It is a good idea to do this pre-push of a version.

#### Test against the example
   
```bash
$ cd examples/simple #(or other nodejs project)
$ npm link @pulumi/fd-provider
$ npm install
$ pulumi stack init test
$ pulumi up
```

Now that you have completed all of the above steps, you have a working provider that generates a random string for you.

#### A brief repository overview

You now have:

1. A `provider/` folder containing the building and implementation logic
    1. `cmd/pulumi-resource-fd-provider/main.go` - holds the provider's sample implementation logic.
2. `deployment-templates` - a set of files to help you around deployment and publication
3. `sdk` - holds the generated code libraries created by `pulumi-gen-fd-provider/main.go`
4. `examples` a folder of Pulumi programs to try locally and/or use in CI.
5. A `Makefile` and this `README`.

### Build Examples

Create a YAML example program using the resources defined in your provider, and place it in the `examples/` folder.

Then run:
```bash
make gen_examples
```

You can now repeat the steps for [build, install, and test](#test-against-the-example).

## Configuring CI and releases

For now, please ignore `deployment-templates` and `.github.ignore`. These will be built out soon.

For now, releases are manual.

## References

Other resources/examples for implementing providers:
* [Pulumi Command provider](https://github.com/pulumi/pulumi-command/blob/master/provider/pkg/provider/provider.go)
* [Pulumi Go Provider repository](https://github.com/pulumi/pulumi-go-provider)
