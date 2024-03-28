package Travel

import (
	"github.com/Thenecromance/OurStories/backend"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	Interface.ControllerBase
	model Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) Name() string {
	return "travel"
}

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{
		model: Model{},
	}
	c.RouteNode = Interface.NewNode("api", c.Name())
	c.LoadChildren(i...)
	return c
}

/*func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
	// parent group is  /api/
	c.ParentGroup = group
	//setup self group as /api/user
	c.Group = group.Group("/" + c.Name())
}*/

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	//c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the controller's group
func (c *Controller) Use(middleware ...gin.HandlerFunc) {
	c.Use(middleware...)
}

func (c *Controller) BuildRoutes() {
	/*	c.Group.POST("/addTravel", c.addTravel)
		c.Group.POST("/removeTravel", c.removeTravel)*/
	// I want to use the RESTful API
	c.GET("/", c.getTravels)
	c.POST("/", c.addTravel)
	c.PUT("/", c.updateTravel)
	c.DELETE("/", c.removeTravel)

	c.ChildrenBuildRoutes()
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) addTravel(ctx *gin.Context) {
	var received Data

	err := ctx.ShouldBind(&received)
	if err != nil {
		logger.Get().Info(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"info": err.Error(),
		})
		return
	}

	err = c.model.AddTravel(received)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"info": "failed to add travel",
		})
		return
	}

	backend.Resp(ctx, "add travel complete!")
}

func (c *Controller) removeTravel(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		backend.RespErr(ctx, "id is empty")
		return
	}

	iid, err := strconv.Atoi(id)
	if err != nil {
		backend.RespErr(ctx, "id is not a number")
		return
	}

	err = c.model.RemoveTravel(iid)
	if err != nil {
		logger.Get().Error(err.Error())
		backend.RespErr(ctx,
			"failed to remove travel",
		)
		return
	}

	backend.Resp(ctx, "success")
}

func (c *Controller) getTravels(ctx *gin.Context) {
	usrId := ctx.Query("user")
	if usrId == "" {
		backend.RespErr(ctx, "usrId is empty")
		return
	}
	iid, err := strconv.Atoi(usrId)
	if err != nil {
		backend.RespErr(ctx, "usrId is not a number")
		return

	}
	data, err := c.model.GetTravelByUserId(iid)
	if err != nil {
		logger.Get().Error(err.Error())
		backend.RespErr(ctx, "failed to get travel")
		return
	}
	backend.RespWithCount(ctx, data, len(data))
}

func (c *Controller) updateTravel(ctx *gin.Context) {
	data := UpdateData{}

	err := ctx.ShouldBind(&data)
	if err != nil {
		backend.RespErr(ctx, "wrong params")
		return
	}

	err = c.model.updateToDatabase(data)
	if err != nil {
		backend.RespErr(ctx, "failed to update travel")

		return
	}

	backend.Resp(ctx, "success")

}
