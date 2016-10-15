package web

// RegisterChain adds path handlers for specified method
func (i *Instance) RegisterChain(handlersChain HandlersChain) {
	for m, mHandlers := range handlersChain {
		switch m {
		case MethodPOST:
			for path, handler := range mHandlers {
				i.Router.POST(path, handler)
			}
		case MethodDELETE:
			for path, handler := range mHandlers {
				i.Router.DELETE(path, handler)
			}
		case MethodGET:
			for path, handler := range mHandlers {
				i.Router.GET(path, handler)
			}
		case MethodHEAD:
			for path, handler := range mHandlers {
				i.Router.HEAD(path, handler)
			}
		case MethodOPTIONS:
			for path, handler := range mHandlers {
				i.Router.OPTIONS(path, handler)
			}
		case MethodPATCH:
			for path, handler := range mHandlers {
				i.Router.PATCH(path, handler)
			}
		case MethodPUT:
			for path, handler := range mHandlers {
				i.Router.PUT(path, handler)
			}
		}
	}
}
