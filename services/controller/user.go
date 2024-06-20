package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/router"

	"github.com/Thenecromance/OurStories/services/models"
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

func (uc *UserController) GetRoutes() []Interface.Router {
	return []Interface.Router{uc.routers.login, uc.routers.register, uc.routers.logout, uc.routers.profile}
}

func (uc *UserController) setupRouters() {

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

//-----------------------------------------------------------
//User section
//-----------------------------------------------------------

func (uc *UserController) login(ctx *gin.Context) {

	resp := response.New()
	defer resp.Send(ctx)

	// TODO: add token auth here
	/*	auth, err := ctx.Cookie("Authorization")
		if err == nil {
			resp.SetCode(response.OK).AddData("Already login")
			return
		}*/

	info := models.UserLogin{}
	err := ctx.ShouldBind(&info)
	if err != nil {
		log.Error("params binding failed :", err)
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	log.Debugf("here should be add a precheck method for the user info incase some one might use the shit names")

	usr, err := uc.userService.AuthorizeUser(&info)
	if err != nil {
		log.Error("authorize user failed :", err)
		resp.SetCode(response.BadRequest).AddData("Invalid username or password")
		return
	}

	// generate the token to client and save it to the cookie
	token := uc.userService.SignedTokenToUser(usr.UserName)
	uc.signTokenToClient(ctx, token)

	// when login success, return the basic user info to the client
	resp.SetCode(response.OK).AddData(usr.UserBasicDTO)
}

func (uc *UserController) register(ctx *gin.Context) {
	// prebuild the response and use defer to send the response
	resp := response.New()
	defer resp.Send(ctx)

	// get the user info from the request
	info := models.UserRegister{}
	if err := ctx.ShouldBind(&info); err != nil {
		log.Error("params binding failed :", err)
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	log.Debugf("here should be add a precheck method for the user info")

	// before added to the database, check if the user or email is already exist
	if uc.userService.HasUserAndEmail(info.UserName, info.Email) {
		resp.SetCode(response.BadRequest).AddData("User or email already exist")
		return
	}

	// add the user to the database
	if err := uc.userService.AddUser(&info); err != nil {
		log.Error("add user failed :", err)
		resp.SetCode(response.BadRequest).AddData("Register failed")
		return
	}

	// generate the token to client and save it to the cookie
	uc.signTokenToClient(ctx, info.UserName)
	resp.SetCode(response.OK).AddData("Register success")
}

func (uc *UserController) logout(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	//delete the token
	ctx.SetCookie("Authorization", "", -1, "/", "", false, true)
	resp.SetCode(response.OK).AddData("Logout success")
}

//-----------------------------------------------------------
//Profile section
//-----------------------------------------------------------

// getProfile is the method to get the user profile
func (uc *UserController) getProfile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

func (uc *UserController) updateProfile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

func (uc *UserController) signTokenToClient(ctx *gin.Context, token string) {
	token = "Bearer " + token
	ctx.SetCookie("Authorization", token, 3600, "/", "", false, true)
}

func NewUserController(userService *services.UserService) *UserController {

	uc := &UserController{
		userService: userService,
	}
	uc.setupRouters()
	return uc
}
