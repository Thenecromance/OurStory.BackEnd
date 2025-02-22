package controller

import (
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/server/response"
	"github.com/Thenecromance/OurStories/server/route"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
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

func (uc *UserController) Name() string {
	return "UserController"
}

func (uc *UserController) Initialize() {
}

func (uc *UserController) SetupRoutes() {
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

func (uc *UserController) GetRoutes() []Interface.IRoute {
	return []Interface.IRoute{uc.routers.login, uc.routers.register, uc.routers.logout, uc.routers.profile}
}

//-----------------------------------------------------------
//User section
//-----------------------------------------------------------

func (uc *UserController) hasCredential(ctx *gin.Context) bool {
	obj, exists := ctx.Get(constants.AuthObject)
	log.Info("User already Already login ", obj, exists)
	if exists {
		log.Info("User already Already login ", obj)
		return true
	}
	return false
}

func (uc *UserController) preCheckByToken(ctx *gin.Context, resp *response.Response) bool {
	token, err := ctx.Cookie("Authorization")
	log.Debugf(token)
	if token != "" && err == nil {
		// check the token is valid
		ok, err := JWT.TokenValid(ctx)
		if !ok || err != nil {
			return false
		}
		// get the user info from the token
		claim, err := JWT.Instance().GetUserClaimFromToken(token)
		if err != nil {
			return false
		}

		if userClaim, ok := claim.(*models.UserClaim); ok {
			user, err := uc.service.GetUser(userClaim.Id)
			if user == nil || err != nil {
				return false
			}
			resp.SetCode(response.OK).AddData(user.UserBasicDTO)
			return true
		} else {
			id := int64(claim.(map[string]interface{})["id"].(float64))
			user, err := uc.service.GetUser(id)
			if user == nil || err != nil {
				return false
			}
			resp.SetCode(response.OK).AddData(user.UserBasicDTO)

			return true
		}

	}
	return false
}

// login  337.7µs about ~3000 QPS
func (uc *UserController) login(ctx *gin.Context) {

	resp := response.New()
	defer resp.Send(ctx)
	// due to login does not Attached middle ware, so the user claim will not be attached to the context
	/*if uc.hasCredential(ctx) {
		resp.SetCode(response.OK).AddData("Already login")
		return
	}*/

	// use the token to check if the user is already login
	// if the user is already login, return the user info to the client
	if uc.preCheckByToken(ctx, resp) {
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

	loginSuccess, err := uc.service.AuthorizeUser(&info)
	if err != nil {
		log.Info("authorize user failed :", err)
		resp.SetCode(response.BadRequest).AddData("Invalid username or password")
		return
	}
	if !loginSuccess {
		resp.SetCode(response.OK).AddData("Login failed, please check username or password")
		return
	}

	usr, err := uc.service.GetUserByUsername(info.UserName)
	if err != nil {
		log.Error("something goes wrong with  uc.service.GetUserByUsername(info.UserName) please check ", err)
		return
	}

	//set up claim for sign token to client
	claim_ := models.UserClaim{
		UserName: usr.UserName,
		Id:       usr.UserId,
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

	// check this client already has the token to identify the user is already login
	// but seems like this is not necessary
	_, exists := ctx.Get("Authorization")
	if exists {
		ok, err := JWT.TokenValid(ctx)
		if err != nil || ok {
			log.Error("token valid failed :", err)
			resp.SetCode(response.BadRequest).AddData("Invalid request")
			return
		}
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
	// sign the token to the client
	uc.signTokenToClient(ctx, token)

	usr, err := uc.service.GetUser(uid)
	if err != nil {
		log.Error("get user failed :", err)
		return
	}

	resp.SetCode(response.OK).AddData(usr.UserBasicDTO)
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

	claim := getAuthObject(ctx)

	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}
	log.Debugf("uid: %d", claim.Id)

	usrName := ctx.Param("username")
	if usrName != claim.UserName {
		resp.SetCode(response.OK).AddData("Invalid request")
		return
	}

	usr, err := uc.service.GetUserByUsername(usrName)
	if err != nil {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	if usr.UserName != claim.UserName || usr.UserId != claim.Id {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	resp.SetCode(response.OK).AddData(usr.UserAdvancedDTO)
}

func (uc *UserController) updateProfile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
	panic("implement me")
}

// signTokenToClient is the method to sign the token to the client
func (uc *UserController) signTokenToClient(ctx *gin.Context, token string) {
	//token = "Bearer " + token
	ctx.SetCookie("Authorization", token, 3600, "/", "", false, true)
}

func (uc *UserController) cleanUpClientToken(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", -1, "/", "", false, true)
}

func NewUserController(userService services.UserService) Interface.IController {

	uc := &UserController{
		service: userService,
	}
	return uc
}
