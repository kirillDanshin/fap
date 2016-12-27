package static

import "github.com/kirillDanshin/fap/lib/types"

const template = `package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", ":8080", "address to listen")
	dir  = flag.String("dir", "./static", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("Serving %s on %s\n", *dir, *addr)
	fasthttp.ListenAndServe(*addr, fasthttp.FSHandler(*dir, 0))
}`

// Generate returns a map with generated
// package structure
func Generate(packageName string) (types.PackageStructure, error) {
	structure := types.PackageStructure{
		"main.go": template,
	}
	return structure, nil
}
