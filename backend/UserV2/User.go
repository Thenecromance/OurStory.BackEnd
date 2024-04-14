package UserV2

import (
	response "github.com/Thenecromance/OurStories/backend/Response"
	"github.com/Thenecromance/OurStories/backend/UserV2/model"
	"github.com/Thenecromance/OurStories/base/log"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	group   *Interface.GroupNode
	auth    *model.Authorization
	profile *model.Profile
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController() *Controller {
	c := &Controller{
		auth:    model.NewAuth(),
		profile: model.NewProfile(),
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
	//c.group.Router.GET("/:username", c.Middleware(), c.profile)
	//c.group.Router.PUT("/:username", c.Middleware(), c.updateProfile)

}

func (c *Controller) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := ctx.Cookie("Authorization")
		log.Debug(token)
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
	log.Debug("login request")
	resp := response.New(ctx)
	defer resp.Send()

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

	log.Debug("login Info : ", info)
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
