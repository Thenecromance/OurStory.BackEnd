package MiddleWare

import (
	"fmt"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type controller struct {
	middleWareMap map[string]gin.HandlerFunc
}

func (c *controller) RegisterMiddleWare(name string, handler gin.HandlerFunc) {
	c.middleWareMap[name] = handler
}
func (c *controller) GetMiddleWare(name string) (gin.HandlerFunc, error) {
	handler, ok := c.middleWareMap[name]
	if !ok {
		return nil, fmt.Errorf("middleware %s not found", name)
	}
	return handler, nil
}

func NewController() Interface.MiddleWareController {
	return &controller{
		middleWareMap: make(map[string]gin.HandlerFunc),
	}
}
