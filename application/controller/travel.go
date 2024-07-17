package controller

import (
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	Interface2 "github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/server/response"
	route2 "github.com/Thenecromance/OurStories/server/route"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type travelRouter struct {
	travel     Interface2.IRoute
	travelList Interface2.IRoute
}

type TravelController struct {
	groups travelRouter

	service services.TravelService
}

func (tc *TravelController) Name() string {
	return "TravelController"
}

func (tc *TravelController) Initialize() {
}

func (tc *TravelController) SetupRoutes() {
	mw := JWT.Middleware()
	tc.groups.travel = route2.NewREST("/api/travel/")
	{
		tc.groups.travel.SetMiddleWare(mw)
		tc.groups.travel.SetHandler(
			tc.getTravel,    // GET
			tc.createTravel, // POST
			tc.updateTravel, // PUT
			tc.deleteTravel, // DELETE
		)
	}
	tc.groups.travelList = route2.NewRouter("/api/travel/list", "POST")
	{
		tc.groups.travelList.SetMiddleWare(mw)
		tc.groups.travelList.SetHandler(tc.getTravelList)
	}
}

func (tc *TravelController) GetRoutes() []Interface2.IRoute {
	return []Interface2.IRoute{tc.groups.travel, tc.groups.travelList}
}

//-----------------handlers begin-----------------

func (tc *TravelController) getTravel(ctx *gin.Context) {
	log.Debugf("getTravel")
	resp := response.New()
	defer resp.Send(ctx)
	travelID := ctx.Query("id")

	claim := getAuthObject(ctx)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	log.Debugf("uid: %d", claim.Id)
	//get travel from database
	travel, err := tc.service.GetTravelByID(travelID, claim.Id)
	if err != nil {
		resp.Error(err.Error())
		return
	}

	resp.Success(travel)
}

func (tc *TravelController) createTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	claim := getAuthObject(ctx)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	log.Debugf("uid: %d", claim.Id)
	//UserName := obj.(map[string]interface{})["name"].(string)

	dto := &models.Travel{}
	if err := ctx.ShouldBindJSON(dto); err != nil {
		resp.Error("invalid request")
		return
	}

	// a simple precheck
	if dto.UserId != claim.Id {
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

	claim := getAuthObject(ctx)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	log.Debugf("uid: %d", claim.Id)
	//UserName := obj.(map[string]interface{})["name"].(string)

	dto := &models.Travel{}
	if err := ctx.ShouldBindJSON(dto); err != nil {
		resp.Error("invalid request")
		return
	}

	// a simple precheck
	if dto.UserId != claim.Id {
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

	claim := getAuthObject(ctx)

	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}
	log.Debugf("uid: %d", claim.Id)
	//UserName := obj.(map[string]interface{})["name"].(string)
	id := ctx.Query("id")

	err := tc.service.DeleteTravel(id, claim.Id)
	if err != nil {
		resp.Error(err.Error())
		return
	}
	resp.Success("success")
}

func (tc *TravelController) getTravelList(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	claim := getAuthObject(ctx)

	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}
	log.Debugf("uid: %d", claim.Id)

	//UserName := obj.(map[string]interface{})["name"].(string)

	lists, err := tc.service.GetTravelList(claim.Id)
	if err != nil {
		return
	}
	resp.Success(lists)
}

//-----------------handlers end-----------------

func NewTravelController(s services.TravelService) Interface2.IController {
	tc := &TravelController{
		service: s,
	}
	return tc
}
