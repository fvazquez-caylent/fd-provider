package metadata

import (
	"strings"

	gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
)

const PROVIDER_NAME = "xyzqw"
const DESCRIPTION = ""
const README = ``

var PROVIDER_NAME_SNAKE = strings.ReplaceAll(PROVIDER_NAME, "-", "_")

var Metadata = schema.Metadata{
	DisplayName: PROVIDER_NAME,
	LanguageMap: map[string]any{
		"nodejs": map[string]any{
			"packageName":        "@alphonsocode/pulumi-provider-" + PROVIDER_NAME,
			"packageDescription": DESCRIPTION,
			"readme":             README,
		},
		"python": map[string]any{
			"packageName": "pulumi_" + PROVIDER_NAME_SNAKE,
			"readme":      README,
		},
		"go": gen.GoPackageInfo{
			ImportBasePath: "github.com/AlphonsoCode/pulumi-provider-" + PROVIDER_NAME + "/sdk/go/" + PROVIDER_NAME,
		},
	},
}
