// This file is automatically generated by qtc from "main.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line lib/templates/webrdb/qtpl/main.qtpl:1
package qtpl

//line lib/templates/webrdb/qtpl/main.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/templates/webrdb/qtpl/main.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/templates/webrdb/qtpl/main.qtpl:1
func StreamFastHTTPRouter(qw422016 *qt422016.Writer, packageName string) {
	//line lib/templates/webrdb/qtpl/main.qtpl:1
	if packageName == "main" {
		//line lib/templates/webrdb/qtpl/main.qtpl:1
		qw422016.N().S(`package main`)
		//line lib/templates/webrdb/qtpl/main.qtpl:1
	} else {
		//line lib/templates/webrdb/qtpl/main.qtpl:1
		qw422016.N().S(`package `)
		//line lib/templates/webrdb/qtpl/main.qtpl:1
		qw422016.E().S(packageName)
		//line lib/templates/webrdb/qtpl/main.qtpl:1
	}
	//line lib/templates/webrdb/qtpl/main.qtpl:1
	qw422016.N().S(`

import (`)
	//line lib/templates/webrdb/qtpl/main.qtpl:3
	if packageName == "main" {
		//line lib/templates/webrdb/qtpl/main.qtpl:3
		qw422016.N().S(`
	"flag"

`)
		//line lib/templates/webrdb/qtpl/main.qtpl:6
	}
	//line lib/templates/webrdb/qtpl/main.qtpl:6
	qw422016.N().S(`	"github.com/kirillDanshin/dlog"
	"github.com/kirillDanshin/fap/web"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
)
`)
	//line lib/templates/webrdb/qtpl/main.qtpl:11
	if packageName == "main" {
		//line lib/templates/webrdb/qtpl/main.qtpl:11
		qw422016.N().S(`
var (
	address = flag.String("address", "127.0.0.1:3270", "address to bind")
)`)
		//line lib/templates/webrdb/qtpl/main.qtpl:14
	}
	//line lib/templates/webrdb/qtpl/main.qtpl:14
	qw422016.N().S(`

`)
	//line lib/templates/webrdb/qtpl/main.qtpl:16
	if packageName == "main" {
		//line lib/templates/webrdb/qtpl/main.qtpl:16
		qw422016.N().S(`func main() {
	flag.Parse()
	addr := *address
`)
		//line lib/templates/webrdb/qtpl/main.qtpl:19
	} else {
		//line lib/templates/webrdb/qtpl/main.qtpl:19
		qw422016.N().S(`func Run(addr string) {`)
		//line lib/templates/webrdb/qtpl/main.qtpl:19
	}
	//line lib/templates/webrdb/qtpl/main.qtpl:19
	qw422016.N().S(`
	dlog.F("Registering handlers")

	web.RegisterChain(web.HandlersChain{
		web.MethodGET: web.MethodChain{
			"/": homeHandler,
		}
	})

	dlog.F("Listening on %%s", addr)
	myutils.LogFatalError(web.ListenAndServe(addr))
}
`)
//line lib/templates/webrdb/qtpl/main.qtpl:31
}

//line lib/templates/webrdb/qtpl/main.qtpl:31
func WriteFastHTTPRouter(qq422016 qtio422016.Writer, packageName string) {
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	StreamFastHTTPRouter(qw422016, packageName)
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line lib/templates/webrdb/qtpl/main.qtpl:31
}

//line lib/templates/webrdb/qtpl/main.qtpl:31
func FastHTTPRouter(packageName string) string {
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	WriteFastHTTPRouter(qb422016, packageName)
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	qs422016 := string(qb422016.B)
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
	//line lib/templates/webrdb/qtpl/main.qtpl:31
	return qs422016
//line lib/templates/webrdb/qtpl/main.qtpl:31
}
