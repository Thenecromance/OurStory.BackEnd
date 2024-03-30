package User

import (
	"github.com/Thenecromance/OurStories/backend"
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
	c.RouteNode = Interface.NewNode("api", c.Name())
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "user"
}

/*
	func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
		// parent group is  /api/
		c.ParentGroup = group
		//setup self group as /api/user
		c.Group = group.Group("/" + c.Name())
	}
*/

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	//c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the controller's group
func (c *Controller) AddMiddleWare(middleware ...gin.HandlerFunc) {
	c.AddMiddleWare(middleware...)
}

func (c *Controller) BuildRoutes() {
	c.POST("/login", c.login)
	c.POST("/register", c.register)
	c.ChildrenBuildRoutes()
}

//----------------------------Interface.Controller Implementation--------------------------------

// ------------------------------------------------------------

func (c *Controller) login(ctx *gin.Context) {

	var user Info
	err := ctx.ShouldBind(&user) // both support form and json
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}

	err = c.model.login(&user)
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}
	user.Password = ""

	backend.Resp(ctx, user)
	return
}

func (c *Controller) register(ctx *gin.Context) {
	var user Info
	err := ctx.ShouldBind(&user)
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}

	err = c.model.register(user)
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}
	backend.Resp(ctx, user)
}

//------------------------------------------------------------
