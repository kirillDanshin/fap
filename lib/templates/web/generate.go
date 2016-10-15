package web

import (
	"github.com/kirillDanshin/fap/lib/templates/web/qtpl"
	"github.com/kirillDanshin/fap/lib/types"
	"github.com/kirillDanshin/myutils"
)

// Generate returns a map with generated
// package structure
func Generate(packageName string) (types.PackageStructure, error) {
	if packageName == "" {
		packageName = "main"
	}
	path := "main.go"
	if packageName != "main" {
		path = myutils.Concat(packageName, "/web.go")
	}
	structure := types.PackageStructure{
		"homeHandler.go":   qtpl.HomeHandler(packageName),
		"getRouteChain.go": qtpl.GetRouteChain(packageName),
		path:               qtpl.FastHTTPRouter(packageName),
	}
	return structure, nil
}
