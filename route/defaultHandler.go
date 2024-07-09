package route

import (
	"github.com/Thenecromance/OurStories/response"
	"github.com/gin-gonic/gin"
)

func _defaultHandler(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
	resp.Code = response.NotAcceptable
	resp.Meta.Count = 0
	resp.Data = gin.H{
		"system": "service not found",
	}
}

func DefaultRESTHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		_defaultHandler,
		_defaultHandler,
		_defaultHandler,
		_defaultHandler,
	}
}

// DefaultHandler will be used to set the default handler of the route
func DefaultHandler() gin.HandlerFunc {
	return _defaultHandler
}

//--------------------------------------------
// DefaultMiddleware
//--------------------------------------------

func _defaultMiddleware(ctx *gin.Context) {
	ctx.Next() // just use ctx.Next() to skip the middleware
}

// DefaultMiddleware is the default middleware of the route
func DefaultMiddleware() gin.HandlersChain {
	return gin.HandlersChain{
		_defaultMiddleware,
	}
}
