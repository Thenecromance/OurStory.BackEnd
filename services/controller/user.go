package controller

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/server/response"
	"github.com/Thenecromance/OurStories/server/router"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type userRouters struct {
	login    Interface.Router
	register Interface.Router
	logout   Interface.Router
	profile  Interface.Router
}

type UserController struct {
	userService *services.UserService
	routers     userRouters
}

func (uc *UserController) RegisterRoutes(engine *gin.Engine) {
	/*	userGroup := engine.Group("/user")
		{
			userGroup.POST("/login", uc.login)
			userGroup.POST("/register", uc.register)
			userGroup.POST("/logout", uc.logout)
		}
		{
			userGroup.GET("/:username", uc.getProfile)
			userGroup.PUT("/:username", uc.updateProfile)
		}*/

	{
		uc.routers.login = router.NewRouter()
		{
			uc.routers.login.SetPath("/api/user/login")
			uc.routers.login.SetMethod("POST")
			uc.routers.login.SetHandler(uc.login)
		}
	}
	{
		uc.routers.register = router.NewRouter()
		{
			uc.routers.register.SetPath("/api/user/register")
			uc.routers.register.SetMethod("POST")
			uc.routers.register.SetHandler(uc.register)
		}
	}
	{
		uc.routers.logout = router.NewRouter()
		{
			uc.routers.logout.SetPath("/api/user/logout")
			uc.routers.logout.SetMethod("POST")
			uc.routers.logout.SetHandler(uc.logout)
		}
	}
	{
		uc.routers.profile = router.NewREST()
		{
			uc.routers.profile.SetPath("/api/user/:username")
			uc.routers.profile.SetHandler(uc.getProfile, nil, uc.updateProfile)
		}
	}

}

func (uc *UserController) login(ctx *gin.Context) {

	resp := response.New()
	defer resp.Send(ctx)

	usrToken, _ := ctx.Cookie("Authorization")
	if usrToken == "" {
		log.Error("get token failed")
		return
	}

	/*if c.auth.ValidByToken(usrToken) {
		resp.SetCode(response.SUCCESS).AddData("Already login")
		return
	}*/

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

	resp.SetCode(response.Accepted).AddData(usr.ToUserResponse())
	token := c.auth.SignedToken(usr.ToUserClaim())
	uc.setTokenCookie(ctx, token)
}

func (uc *UserController) setTokenCookie(ctx *gin.Context, token string) {
	token = "Bearer " + token
	ctx.SetCookie("Authorization", token, 3600, "/", "", false, true)
}

func (uc *UserController) register(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

func (uc *UserController) logout(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	//delete the token
	ctx.SetCookie("Authorization", "", -1, "/", "", false, true)
	resp.SetCode(response.Accepted).AddData("Logout success")
}

func (uc *UserController) getProfile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

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

func (uc *UserController) updateProfile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
