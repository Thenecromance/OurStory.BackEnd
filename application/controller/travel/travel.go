package travel

import (
	response "github.com/Thenecromance/OurStories/backend/Response"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Travel section still has a lot of shit to do , user auth, monkey created bugs ,

type Controller struct {
	group *Interface.GroupNode
	model Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) Name() string {
	return "travel"
}

func NewController() Interface.Controller {
	c := &Controller{
		model: Model{},
	}

	return c
}
func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "api")
}

func (c *Controller) BuildRoutes() {

	// I want to use the RESTFUL API
	c.group.Router.GET("/", c.getTravels)
	c.group.Router.POST("/", c.addTravel)
	c.group.Router.PUT("/", c.updateTravel)
	c.group.Router.DELETE("/", c.removeTravel)
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) addTravel(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	var received ClientData

	err := ctx.ShouldBind(&received)
	if err != nil {
		log.Info(err.Error())
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

	resp.SetCode(response.SUCCESS).AddData("add travel complete!")
}

func (c *Controller) removeTravel(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	id := ctx.Query("id")
	if id == "" {
		resp.AddData("id is empty")
		//backend.RespErr(ctx, "id is empty")
		return
	}

	//iid, err := strconv.ParseInt(id, 10, 64)
	//log.Info(iid)
	//if err != nil {
	//	backend.RespErr(ctx, "id is not a number")
	//	return
	//}
	err := c.model.RemoveTravel(id)
	if err != nil {
		log.Error(err.Error())
		resp.AddData(
			"failed to remove travel",
		)
		return
	}
	resp.SetCode(response.SUCCESS).AddData("success")
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
	resp := response.New(ctx)
	defer resp.Send()

	usrId := ctx.Query("user")
	if usrId == "" {
		resp.AddData("usrId is empty")
		return
	}
	iid, err := strconv.Atoi(usrId)
	if err != nil {
		resp.AddData("usrId is not a number")
		return

	}
	data, err := c.model.GetTravelListByUser(iid)

	for index := range data {
		c.checkState(&data[index])
	}

	if err != nil {
		log.Error(err.Error())
		resp.AddData("failed to get travel")
		return
	}
	resp.AddData(data)
}

func (c *Controller) updateTravel(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

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
