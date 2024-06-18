package Interface

import "github.com/gin-gonic/gin"

type RouterSetter interface {
	SetPath(path string)
	SetMethod(method string)
	SetMiddleWare(middleware gin.HandlersChain)
	SetHandler(handler ...gin.HandlerFunc)
}
type RouterGetter interface {
	GetPath() string
	GetMethod() string
	GetMiddleWare() gin.HandlersChain
	GetHandler() []gin.HandlerFunc // if the subject is RESTFUL, this should be 4 handlers
}

// RouterControl is the interface that wraps the basic methods of a router control.
type RouterControl interface {
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
