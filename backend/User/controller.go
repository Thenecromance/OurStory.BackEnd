package User

import (
	"github.com/Thenecromance/OurStories/backend/api"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase
	model Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{
		model: Model{},
	}
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "user"
}

func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
	// parent group is  /api/
	c.ParentGroup = group
	//setup self group as /api/user
	c.Group = group.Group("/" + c.Name())
}

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the controller's group
func (c *Controller) Use(middleware ...gin.HandlerFunc) {
	c.Group.Use(middleware...)
}

func (c *Controller) BuildRoutes() {
	c.Group.POST("/login", c.login)
	c.Group.POST("/register", c.register)
	c.ChildrenBuildRoutes()
}

//----------------------------Interface.Controller Implementation--------------------------------

// ------------------------------------------------------------
func (c *Controller) login(ctx *gin.Context) {

	var user Info
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(200, api.Response{
			Code:   api.Failed,
			Result: err.Error(),
		})
		return
	}

	err = c.model.login(&user)
	if err != nil {
		ctx.JSON(200, api.Response{
			Code:   api.Failed,
			Result: err.Error(),
		})
		return
	}
	user.Password = ""
	ctx.JSON(200, api.Response{
		Code:   api.Success,
		Result: user,
	})
	return
}

func (c *Controller) register(ctx *gin.Context) {
	var user Info
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(200, api.Response{
			Code:   api.Failed,
			Result: err.Error(),
		})
		return
	}

	err = c.model.register(user)
	if err != nil {
		ctx.JSON(200, api.Response{
			Code:   api.Failed,
			Result: err.Error(),
		})
		return
	}
	ctx.JSON(200, api.Response{
		Code:   api.Success,
		Result: user,
	})
}

//------------------------------------------------------------
