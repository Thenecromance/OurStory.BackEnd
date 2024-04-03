package User

import (
	response "github.com/Thenecromance/OurStories/backend/Response"
	"github.com/Thenecromance/OurStories/backend/User/Data"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignTokenCallback = func(*gin.Context, interface{}) (string, error)
type GetObjectFromTokenCallback = func(string) (interface{}, error)
type AuthorizeCallback = func(interface{}) bool

func emptySignTokenCallback(*gin.Context, interface{}) (string, error) { return "", nil }
func emptyGetObjectFromTokenCallback(string) (interface{}, error)      { return nil, nil }
func emptyAuthorizeCallback(interface{}) bool                          { return false }

type Controller struct {
	group *Interface.GroupNode
	model Model

	signedToken    SignTokenCallback
	getObject      GetObjectFromTokenCallback
	tokenAuthorize AuthorizeCallback
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController() Interface.Controller {
	c := &Controller{
		model: Model{},
		//signedToken:    emptySignTokenCallback,
		getObject:      emptyGetObjectFromTokenCallback,
		tokenAuthorize: emptyAuthorizeCallback,
	}

	return c
}

func (c *Controller) Name() string {
	return "user"
}

func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "api")
}

func (c *Controller) BuildRoutes() {
	c.model.init()
	c.group.Router.POST("/", c.login)
	c.group.Router.POST("/register", c.register)

	c.group.Router.GET("/:username", c.profile)
	//c.group.Router.POST("/:username", c.logout)
	c.group.Router.PUT("/:username", c.updateProfile)
	//c.group.Router.DELETE("/:username", c.deleteProfile)
}

//----------------------------Interface.Controller Implementation--------------------------------

// ------------------------------------------------------------
func (c *Controller) login(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	var authInfo Data.AuthorizationInfo
	if err := ctx.ShouldBind(&authInfo); err != nil {
		resp.SetCode(http.StatusBadRequest).AddData("Invalid request")
		return
	}

	//check if the user's credentials are valid
	info, err := c.model.Login(&authInfo)
	if err != nil {
		resp.SetCode(http.StatusUnauthorized).AddData(err.Error())
		return
	}
	var res struct {
		Data.CommonInfo
		UserName string `json:"username"`
	}
	res.CommonInfo = info.CommonInfo
	res.UserName = info.UserName

	resp.SetCode(response.SUCCESS).AddData(res)

	//generate a token
	if c.signedToken == nil {
		return
	}
	token, err := c.signedToken(ctx, info)
	if err != nil {
		resp.SetCode(http.StatusInternalServerError).AddData(err.Error())
		return
	}
	logger.Get().Infof("Signed token for user %s with token %s", info.UserName, token)
}

func (c *Controller) register(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	var userReg struct {
		Data.AuthorizationInfo
		Email string `json:"email" form:"email" binding:"required"`
	}
	if err := ctx.ShouldBind(&userReg); err != nil {
		resp.SetCode(http.StatusBadRequest).AddData("Invalid request")
		return
	}

	//register the user
	err := c.model.Register(&userReg.AuthorizationInfo, userReg.Email)
	if err != nil {
		resp.SetCode(response.FAIL).AddData(err.Error())
		return
	}

	resp.SetCode(response.SUCCESS).AddData("User created successfully")

	if c.signedToken == nil {
		return
	}
	token, err := c.signedToken(ctx, userReg)
	if err != nil {
		//http.StatusInternalServerError
		resp.SetCode(response.FAIL).AddData(err.Error())
		return
	}

	logger.Get().Infof("Signed token for user %s with token %s", userReg.UserName, token)
}

func (c *Controller) profile(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	// user profile id
	username := ctx.Param("username")
	if username == "" {
		return
	}

	usrInfo, err := c.model.Profile(username)

	if err != nil {
		resp.SetCode(http.StatusUnauthorized).AddData(err.Error())
		return
	}

	resp.SetCode(response.SUCCESS).AddData(usrInfo.CommonInfo)
	return
}

func (c *Controller) updateProfile(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	newInfo := &Data.CommonInfo{}
	if err := ctx.ShouldBind(newInfo); err != nil {
		resp.SetCode(http.StatusBadRequest).AddData("Invalid request")
		return
	}

	username := ctx.Param("username")
	if username == "" {
		return
	}

	profile, err := c.model.UpdateProfile(username, newInfo)
	if err != nil {
		resp.SetCode(response.FAIL).AddData(err.Error())
		return
	}

	resp.SetCode(response.SUCCESS).AddData(profile.CommonInfo)

	return

}

func (c *Controller) logout(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	// user profile id
	username := ctx.Param("username")
	if username == "" {
		return
	}

	c.model.LogoutUser(username)
	resp.SetCode(response.SUCCESS).AddData("User logged out successfully")
}

//------------------------------------------------------------
