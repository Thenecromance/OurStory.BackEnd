package Interface

import "github.com/gin-gonic/gin"

type IRouterSetter interface {
	// SetPath Set the path of the route
	SetPath(path string)
	// SetMethod Set the method of the route if the route is RESTFUL, this method will not be used
	SetMethod(method string)
	// SetMiddleWare Set the middleware of the route, the middleware will be used before the handler
	SetMiddleWare(middleware gin.HandlerFunc)
	// SetHandler Set the handler of the route, the handler will be used to handle the request
	// if the route is RESTFUL, this method will be followed by GET, POST, PUT, DELETE
	SetHandler(handler ...gin.HandlerFunc)
}
type IRouterGetter interface {
	// GetPath Get the path of the route
	GetPath() string
	// GetMethod Get the method of the route
	GetMethod() string
	// GetMiddleWare Get the middleware of the route
	GetMiddleWare() gin.HandlersChain
	// GetHandler Get the handler of the route
	GetHandler() []gin.HandlerFunc // if the subject is RESTFUL, this should be 4 handlers
}

// IRouterControl is the interface that wraps the basic methods of a route control.
type IRouterControl interface {
	// Enable the route , the route will be used
	Enable()
	// Disable the route , the route will not be used
	Disable()
}

// IRoute is the interface that wraps the basic methods of a route.
type IRoute interface {
	IsRESTFUL() bool

	IRouterControl

	IRouterSetter
	IRouterGetter
}
