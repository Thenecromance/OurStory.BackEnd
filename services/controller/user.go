package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"

	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

const (
	authKey = "AuthObject"
)

type userRouters struct {
	login    Interface.IRoute
	register Interface.IRoute
	logout   Interface.IRoute
	profile  Interface.IRoute
}

type UserController struct {
	service services.UserService
	routers userRouters
}

func (uc *UserController) GetRoutes() []Interface.IRoute {
	return []Interface.IRoute{uc.routers.login, uc.routers.register, uc.routers.logout, uc.routers.profile}
}

func (uc *UserController) setupRouters() {
	mw := JWT.Middleware()
	{
		uc.routers.login = route.NewDefaultRouter()
		{
			uc.routers.login.SetPath("/api/user/login")
			uc.routers.login.SetMethod("POST")
			//uc.routers.login.SetMiddleWare(mw)
			uc.routers.login.SetHandler(uc.login)
		}
	}
	{
		uc.routers.register = route.NewDefaultRouter()
		{
			uc.routers.register.SetPath("/api/user/register")
			uc.routers.register.SetMethod("POST")

			uc.routers.register.SetHandler(uc.register)
		}
	}
	{
		uc.routers.logout = route.NewRouter("/api/user/logout", "POST")
		{

			uc.routers.logout.SetMiddleWare(mw)
			uc.routers.logout.SetHandler(uc.logout)
		}
	}
	{
		uc.routers.profile = route.NewREST("/api/user/:username")
		{
			uc.routers.profile.SetMiddleWare(mw)
			uc.routers.profile.SetHandler(uc.getProfile, nil, uc.updateProfile)
		}
	}

}

//-----------------------------------------------------------
//User section
//-----------------------------------------------------------

func (uc *UserController) hasCredential(ctx *gin.Context) bool {
	obj, exists := ctx.Get(authKey)
	log.Info("User already Already login ", obj, exists)
	if exists {
		log.Info("User already Already login ", obj.(*models.UserClaim))
		return true
	}
	return false
}

func (uc *UserController) login(ctx *gin.Context) {

	resp := response.New()
	defer resp.Send(ctx)

	// TODO: add token auth here
	/*	auth, err := ctx.Cookie("Authorization")
		if err == nil {
			resp.SetCode(response.OK).AddData("Already login")
			return
		}*/

	if uc.hasCredential(ctx) {
		resp.SetCode(response.OK).AddData("Already login")
		return
	}

	info := models.UserLogin{}
	err := ctx.ShouldBind(&info)
	if err != nil {
		log.Error("params binding failed :", err)
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	log.Debugf("here should be add a precheck method for the user info incase some one might use the shit names")

	usr, err := uc.service.AuthorizeUser(&info)
	if err != nil {
		log.Error("authorize user failed :", err)
		resp.SetCode(response.BadRequest).AddData("Invalid username or password")
		return
	}

	claim_ := models.UserClaim{
		UserName: usr.UserName,
		Id:       usr.Id,
	}
	// generate the token to client and save it to the cookie
	token := uc.service.SignedTokenToUser(claim_)
	uc.signTokenToClient(ctx, token)

	// when login success, return the basic user info to the client
	resp.SetCode(response.OK).AddData(usr.UserBasicDTO)
}

func (uc *UserController) register(ctx *gin.Context) {
	// prebuild the response and use defer to send the response
	resp := response.New()
	defer resp.Send(ctx)

	if uc.hasCredential(ctx) {
		log.Infof("Already Logged in")
		resp.SetCode(response.OK).AddData("Already login")
		return
	}

	// get the user info from the request
	info := models.UserRegister{}
	if err := ctx.ShouldBind(&info); err != nil {
		log.Error("params binding failed :", err)
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	log.Debugf("here should be add a precheck method for the user info")

	// before added to the database, check if the user or email is already exist
	if uc.service.HasUserAndEmail(info.UserName, info.Email) {
		resp.SetCode(response.BadRequest).AddData("User or email already exist")
		return
	}

	// add the user to the database
	if err := uc.service.AddUser(&info); err != nil {
		log.Error("add user failed :", err)
		resp.SetCode(response.BadRequest).AddData("Register failed")
		return
	}

	// if the user is added successfully, login the user
	{
		uid, err := uc.service.GetUserIdByName(info.UserName)
		if err != nil {
			log.Error("get user failed :", err)
			return
		}

		claim_ := models.UserClaim{
			UserName: info.UserName,
			Id:       uid,
		}
		// generate the token to client and save it to the cookie
		token := uc.service.SignedTokenToUser(claim_)
		// generate the token to client and save it to the cookie
		uc.signTokenToClient(ctx, token)
	}

	resp.SetCode(response.OK).AddData("Register success")
}

func (uc *UserController) logout(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	if !uc.hasCredential(ctx) {
		log.Infof("Already Logged out")
		resp.SetCode(response.OK).AddData("Already logout")
		return
	}

	//delete the token
	uc.cleanUpClientToken(ctx)
	resp.SetCode(response.OK).AddData("Logout success")
}

//-----------------------------------------------------------
//Profile section
//-----------------------------------------------------------

// getProfile is the method to get the user profile
func (uc *UserController) getProfile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	obj, exist := ctx.Get(authKey)
	if !exist {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}
	obj_ := obj.(models.UserClaim)

	usrName := ctx.Param("username")
	if usrName != obj_.UserName {
		resp.SetCode(response.OK).AddData("Invalid request")
		return
	}

	usr, err := uc.service.GetUserByUsername(usrName)
	if err != nil {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	if usr.UserName != obj_.UserName || usr.Id != obj_.Id {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	resp.SetCode(response.OK).AddData(usr.UserAdvancedDTO)
}

func (uc *UserController) updateProfile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

// signTokenToClient is the method to sign the token to the client
func (uc *UserController) signTokenToClient(ctx *gin.Context, token string) {
	token = "Bearer " + token
	ctx.SetCookie("Authorization", token, 3600, "/", "", false, true)
}

func (uc *UserController) cleanUpClientToken(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", -1, "/", "", false, true)
}

func NewUserController(userService services.UserService) *UserController {

	uc := &UserController{
		service: userService,
	}
	uc.setupRouters()
	return uc
}
