package Travel

import (
	"github.com/Thenecromance/OurStories/backend"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Travel section still has a lot of shit to do , user auth, monkey created bugs ,

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

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	//c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the controller's group
func (c *Controller) AddMiddleWare(middleware ...gin.HandlerFunc) {
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
	var received ClientData

	err := ctx.ShouldBind(&received)
	if err != nil {
		logger.Get().Info(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"info": err.Error(),
		})
		return
	}

	err = c.model.AddToSQL(&received)
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

	//iid, err := strconv.ParseInt(id, 10, 64)
	//logger.Get().Info(iid)
	//if err != nil {
	//	backend.RespErr(ctx, "id is not a number")
	//	return
	//}
	err := c.model.RemoveTravel(id)
	if err != nil {
		logger.Get().Error(err.Error())
		backend.RespErr(ctx,
			"failed to remove travel",
		)
		return
	}

	backend.Resp(ctx, "success")
}

//1711952398
//1712016000000

func (c *Controller) checkState(item *ClientData) {
	current := time.Now().Unix() * 1000

	if item.StartTime > current {
		item.State = Prepare
	} else if item.EndTime < current {
		item.State = Finished
	} else {
		item.State = Ongoing
	}

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
	data, err := c.model.GetTravelListByUser(iid)

	for index := range data {
		c.checkState(&data[index])
	}

	if err != nil {
		logger.Get().Error(err.Error())
		backend.RespErr(ctx, "failed to get travel")
		return
	}
	backend.RespWithCount(ctx, data, len(data))
}

func (c *Controller) updateTravel(ctx *gin.Context) {
	backend.ResponseUnImplemented(ctx)
	//data := UpdateData{}
	//
	//err := ctx.ShouldBind(&data)
	//if err != nil {
	//	backend.RespErr(ctx, "wrong params")
	//	return
	//}
	//
	//err = c.model.updateToDatabase(data)
	//if err != nil {
	//	backend.RespErr(ctx, "failed to update travel")
	//
	//	return
	//}
	//
	//backend.Resp(ctx, "success")

}
