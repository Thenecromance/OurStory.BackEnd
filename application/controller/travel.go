package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type travelRouter struct {
	travel     Interface.IRoute
	travelList Interface.IRoute
}

type TravelController struct {
	groups travelRouter

	service services.TravelService
}

func (tc *TravelController) Name() string {
	return "TravelController"
}

func (tc *TravelController) Initialize() {
	tc.SetRoutes()
}

func (tc *TravelController) SetRoutes() {
	mw := JWT.Middleware()
	tc.groups.travel = route.NewREST("/api/travel/")
	{
		tc.groups.travel.SetMiddleWare(mw)
		tc.groups.travel.SetHandler(
			tc.getTravel,    // GET
			tc.createTravel, // POST
			tc.updateTravel, // PUT
			tc.deleteTravel, // DELETE
		)
	}
	tc.groups.travelList = route.NewRouter("/api/travel/list", "POST")
	{
		tc.groups.travelList.SetMiddleWare(mw)
		tc.groups.travelList.SetHandler(tc.getTravelList)
	}
}

func (tc *TravelController) GetRoutes() []Interface.IRoute {
	return []Interface.IRoute{tc.groups.travel, tc.groups.travelList}
}

//-----------------handlers begin-----------------

func (tc *TravelController) getTravel(ctx *gin.Context) {
	log.Debugf("getTravel")
	resp := response.New()
	defer resp.Send(ctx)

	obj, exists := ctx.Get(constants.AuthObject)
	if !exists {
		log.Warn("failed get auth object")
		resp.Unauthorized("please login first")
		return
	}

	travelID := ctx.Query("id")

	//parse user id
	uid := int(obj.(map[string]interface{})["id"].(float64))
	log.Debugf("uid: %d", uid)
	//get travel from database
	travel, err := tc.service.GetTravelByID(travelID, uid)
	if err != nil {
		resp.Error(err.Error())
		return
	}

	resp.Success(travel)
}

func (tc *TravelController) createTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	obj, exists := ctx.Get(constants.AuthObject)
	if !exists {
		log.Warn("failed get auth object")
		resp.Unauthorized("please login first")
		return
	}

	uid := int(obj.(map[string]interface{})["id"].(float64))
	//UserName := obj.(map[string]interface{})["name"].(string)

	dto := &models.TravelDTO{}
	if err := ctx.ShouldBindJSON(dto); err != nil {
		resp.Error("invalid request")
		return
	}

	// a simple precheck
	if dto.UserId != uid {
		resp.Error("invalid params")
		return
	}

	err := tc.service.CreateTravel(dto)
	if err != nil {
		log.Errorf("failed to create travel with error: %s", err.Error())
		resp.Error("failed to create travel")
		return
	}

	resp.Success("success")
}

func (tc *TravelController) updateTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	obj, exists := ctx.Get(constants.AuthObject)
	if !exists {
		log.Warn("failed get auth object")
		resp.Unauthorized("please login first")
		return
	}

	uid := int(obj.(map[string]interface{})["id"].(float64))
	//UserName := obj.(map[string]interface{})["name"].(string)

	dto := &models.TravelDTO{}
	if err := ctx.ShouldBindJSON(dto); err != nil {
		resp.Error("invalid request")
		return
	}

	// a simple precheck
	if dto.UserId != uid {
		resp.Error("invalid params")
		return
	}

	err := tc.service.Update(dto)
	if err != nil {
		log.Errorf("failed to update travel with error: %s", err.Error())
		resp.Error("failed to update travel")
		return
	}

	resp.Success("success")
}

func (tc *TravelController) deleteTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	obj, exists := ctx.Get(constants.AuthObject)
	if !exists {
		log.Warn("failed get auth object")
		resp.Unauthorized("please login first")
		return
	}

	uid := int(obj.(map[string]interface{})["id"].(float64))
	//UserName := obj.(map[string]interface{})["name"].(string)
	id := ctx.Query("id")

	err := tc.service.DeleteTravel(id, uid)
	if err != nil {
		resp.Error(err.Error())
		return
	}
	resp.Success("success")
}

func (tc *TravelController) getTravelList(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	obj, exists := ctx.Get(constants.AuthObject)
	if !exists {
		log.Warn("failed get auth object")
		resp.Unauthorized("please login first")
		return
	}

	uid := int(obj.(map[string]interface{})["id"].(float64))
	//UserName := obj.(map[string]interface{})["name"].(string)

	lists, err := tc.service.GetTravelList(uid)
	if err != nil {
		return
	}
	resp.Success(lists)
}

//-----------------handlers end-----------------

func NewTravelController(s services.TravelService) Interface.IController {
	tc := &TravelController{
		service: s,
	}
	tc.Initialize()
	return tc
}
