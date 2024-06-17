package router

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type REST struct {
	Path string

	MiddleWare gin.HandlersChain

	GETHandler    gin.HandlerFunc
	POSTHandler   gin.HandlerFunc
	PUTHandler    gin.HandlerFunc
	DELETEHandler gin.HandlerFunc
}

func (R *REST) GetPath() string {
	return R.Path
}

func (R *REST) GetMethod() string {
	return ""
}

func (R *REST) IsRESTFUL() bool {
	return true
}

func (R *REST) GetMiddleWare() gin.HandlersChain {
	return R.MiddleWare
}

func (R *REST) GetHandler() gin.HandlerFunc {
	return nil
}

func NewREST(path string) Interface.RouterProxy {
	return &REST{
		Path:       path,
		MiddleWare: nil,

		GETHandler:    defaultFunc,
		POSTHandler:   defaultFunc,
		PUTHandler:    defaultFunc,
		DELETEHandler: defaultFunc,
	}
}
