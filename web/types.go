package web

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// method type is defined in methods.go

type (
	// HandlersChain describes a map that stores
	// method chain info.
	HandlersChain map[method]MethodChain

	// MethodChain describes a map that stores handlers
	// for a path
	MethodChain map[string]func(*fasthttp.RequestCtx)

	// Instance describes a server instance
	Instance struct {
		address string
		Router  *fasthttprouter.Router
	}
)

// New web instance
func New(addr string) *Instance {
	return &Instance{
		address: addr,
		Router:  fasthttprouter.New(),
	}
}
