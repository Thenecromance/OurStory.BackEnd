package Interface

import "github.com/gin-gonic/gin"

type RouterSetter interface {
	// SetPath Set the path of the router
	SetPath(path string)
	// SetMethod Set the method of the router if the router is RESTFUL, this method will not be used
	SetMethod(method string)
	// SetMiddleWare Set the middleware of the router, the middleware will be used before the handler
	SetMiddleWare(middleware gin.HandlersChain)
	// SetHandler Set the handler of the router, the handler will be used to handle the request
	// if the router is RESTFUL, this method will be followed by GET, POST, PUT, DELETE
	SetHandler(handler ...gin.HandlerFunc)
}
type RouterGetter interface {
	// GetPath Get the path of the router
	GetPath() string
	// GetMethod Get the method of the router
	GetMethod() string
	// GetMiddleWare Get the middleware of the router
	GetMiddleWare() gin.HandlersChain
	// GetHandler Get the handler of the router
	GetHandler() []gin.HandlerFunc // if the subject is RESTFUL, this should be 4 handlers
}

// RouterControl is the interface that wraps the basic methods of a router control.
type RouterControl interface {
	// Enable the router , the router will be used
	Enable()
	// Disable the router , the router will not be used
	Disable()
}

// Router is the interface that wraps the basic methods of a router.
type Router interface {
	IsRESTFUL() bool

	RouterControl

	RouterSetter
	RouterGetter
}
