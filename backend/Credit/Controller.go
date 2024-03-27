package Credit

import (
	"github.com/Thenecromance/OurStories/backend"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase

	model Model

	jobRouter *gin.RouterGroup
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
	return "credit"
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
	c.ChildrenBuildRoutes()

	//using RESTful API to handle the credit
	c.Group.GET("/", c.getCreditCount) // get user credits count
	c.Group.POST("/", c.addCredit)
	c.Group.PUT("/", c.modifiedCredit)
	c.Group.DELETE("/", c.modifiedCredit)

	c.jobRouter = c.Group.Group("/job")
	c.jobRouter.GET("/")    //get the job list
	c.jobRouter.POST("/")   //add a new job
	c.jobRouter.PUT("/")    //modify a job
	c.jobRouter.DELETE("/") //delete a job

	//c.Group.Group("/job")
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getCreditCount(ctx *gin.Context) {
	//get the credit count of the user
	backend.ResponseUnImplemented(ctx)
}
func (c *Controller) addCredit(ctx *gin.Context) {
	//add credit to the user
	backend.ResponseUnImplemented(ctx)
}

func (c *Controller) modifiedCredit(ctx *gin.Context) {
	//modify the credit of the user
	backend.ResponseUnImplemented(ctx)
}
