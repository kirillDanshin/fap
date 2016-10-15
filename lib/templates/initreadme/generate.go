package initreadme

import (
	"github.com/kirillDanshin/fap/lib/templates/initreadme/qtpl"
	"github.com/kirillDanshin/fap/lib/types"
)

// Generate returns PackageStruct with README.md
func Generate(name, template string) (types.PackageStructure, error) {
	return types.PackageStructure{
		"README.md": qtpl.Readme(name, template),
	}, nil
}
