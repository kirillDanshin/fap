package web

// method is a private type to safe RegisterChain args from invalid methods
type method string

const (
	// MethodDELETE is for RegisterChain func, it handles DELETE methods
	MethodDELETE method = "DELETE"

	// MethodGET is for RegisterChain func, it handles GET methods
	MethodGET method = "GET"

	// MethodHEAD is for RegisterChain func, it handles HEAD methods
	MethodHEAD method = "HEAD"

	// MethodOPTIONS is for RegisterChain func, it handles OPTIONS methods
	MethodOPTIONS method = "OPTIONS"

	// MethodPATCH is for RegisterChain func, it handles PATCH methods
	MethodPATCH method = "PATCH"

	// MethodPOST is for RegisterChain func, it handles POST methods
	MethodPOST method = "POST"

	// MethodPUT is for RegisterChain func, it handles PUT methods
	MethodPUT method = "PUT"
)
