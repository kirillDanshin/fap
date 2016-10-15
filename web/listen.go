package web

import (
	"errors"

	"github.com/valyala/fasthttp"
)

// ListenAndServe on instance address
func (i *Instance) ListenAndServe() error {
	return fasthttp.ListenAndServe(i.address, i.Router.Handler)
}

// ListenAndServeGZip on instance address with default compress middleware
func (i *Instance) ListenAndServeGZip() error {
	return fasthttp.ListenAndServe(i.address, fasthttp.CompressHandler(i.Router.Handler))
}

// ListenAndServeGZipLevel on instance address with selected compress level
// NOTE: level should be greater or equal to 0 and lower or equal to 9
func (i *Instance) ListenAndServeGZipLevel(level int) error {
	if level > 9 || level < 0 {
		return errors.New("Invalid gzip level")
	}
	return fasthttp.ListenAndServe(i.address, fasthttp.CompressHandlerLevel(i.Router.Handler, level))
}
