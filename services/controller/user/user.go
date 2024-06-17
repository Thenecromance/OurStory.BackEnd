package user

import (
	"github.com/Thenecromance/OurStories/application/model/user"
	response "github.com/Thenecromance/OurStories/backend/Response"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	group   *Interface.GroupNode
	auth    *user.Authorization
	profile *user.Profile
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController() *Controller {
	c := &Controller{
		auth:    user.NewAuth(),
		profile: user.NewProfile(),
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
	//---------------------------------Auth---------------------------------
	c.group.Router.POST("/", c.login)
	c.group.Router.POST("/register", c.register)
	c.group.Router.POST("/logout", c.logout)

	//---------------------------------Profile CRUD---------------------------------
	c.group.Router.GET("/:username", c.Middleware(), c.getProfile)
	c.group.Router.PUT("/:username", c.Middleware(), c.updateProfile)

}

// Middle ware for gin framework
func (c *Controller) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := ctx.Cookie("Authorization")
		if err != nil {
			log.Error("get token failed :", err)
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if !c.auth.ValidByToken(token) {
			log.Error("token invalid")
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) login(ctx *gin.Context) {

	resp := response.New(ctx)
	defer resp.Send()

	usrToken, _ := ctx.Cookie("Authorization")
	if usrToken != "" {
		if c.auth.ValidByToken(usrToken) {
			resp.SetCode(response.SUCCESS).AddData("Already login")
			return
		}
	}

	type loginInfo struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}
	info := loginInfo{}

	err := ctx.ShouldBind(&info)
	if err != nil {
		log.Error("params binding failed :", err)
		resp.AddData("Invalid request")
		return
	}

	usr := c.auth.ValidByUserName(info.Username, info.Password)
	if usr == nil {
		resp.AddData("Invalid username or password")
		return
	}

	resp.SetCode(response.SUCCESS).AddData(usr.ToUserResponse())
	token := c.auth.SignedToken(usr.ToUserClaim())
	c.setTokenCookie(ctx, token)
}

func (c *Controller) setTokenCookie(ctx *gin.Context, token string) {
	token = "Bearer " + token
	ctx.SetCookie("Authorization", token, 3600, "/", "", false, true)
}

func (c *Controller) register(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

}

func (c *Controller) logout(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	//delete the token
	ctx.SetCookie("Authorization", "", -1, "/", "", false, true)
	resp.SetCode(response.SUCCESS).AddData("Logout success")
}

func (c *Controller) getProfile(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	username := ctx.Param("username")
	auth, err := ctx.Cookie("Authorization")
	usrFromToken, err := c.auth.ParseToken(auth)
	if err != nil {
		log.Error("parse token failed :", err)
		resp.AddData("Invalid request")
		return
	}
	if usrFromToken.UserName != username {
		log.Error("username not match")
		resp.AddData("Invalid request")
		return
	}

	usrProfile := c.profile.GetProfile(username)
	resp.SetCode(response.SUCCESS).AddData(usrProfile)
}

func (c *Controller) updateProfile(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
}
