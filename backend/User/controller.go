package User

import (
	"errors"
	"github.com/Thenecromance/OurStories/backend"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SignTokenCallback = func(interface{}) (string, error)
type GetObjectFromTokenCallback = func(string) (interface{}, error)
type AuthorizeCallback = func(interface{}) bool
type Controller struct {
	Interface.ControllerBase
	model Model

	signedToken    SignTokenCallback
	getObject      GetObjectFromTokenCallback
	tokenAuthorize AuthorizeCallback
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
	c.model.init()
	c.POST("/login", c.login)
	c.POST("/register", c.register)
	c.GET("/profile", c.profile)
	c.ChildrenBuildRoutes()
}

//----------------------------Interface.Controller Implementation--------------------------------

// ------------------------------------------------------------

// signTokenToClient will sign the token and set it to the client if set the authorization token
func (c *Controller) signTokenToClient(ctx *gin.Context, usr any) error {
	if c.signedToken == nil {
		return nil
	}

	token, err := c.signedToken(usr)
	if err != nil {
		return errors.New("Failed to sign token")
	}
	ctx.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)
	return nil
}

func (c *Controller) login(ctx *gin.Context) {
	var user Info
	err := ctx.ShouldBind(&user) // both support form and json
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}
	user.Encrypt()
	err = c.model.authUser(&user)
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}

	err = c.signTokenToClient(ctx, user)
	if err != nil {
		backend.RespErr(ctx, "failed to authorize, please try again later")
		return
	}
	user.GetFromSQLByUserName()

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
	user.Encrypt()

	err = c.model.register(&user)
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}

	err = c.signTokenToClient(ctx, user)
	if err != nil {
		backend.RespErr(ctx, "failed to authorize, please try again later")
		return
	}

	backend.Resp(ctx, user)
}

func (c *Controller) profile(ctx *gin.Context) {
	var user Info
	err := ctx.ShouldBind(&user)
	if err != nil {
		backend.RespErr(ctx, err.Error())
		return
	}
	id := ctx.Query("id")
	logger.Get().Info(id)
	user.Id, err = strconv.Atoi(id)
	user.GetFromSQLById()
	backend.Resp(ctx, user)
}

//------------------------------------------------------------
