package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
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

func (tc *TravelController) getTravel(ctx *gin.Context) {
	log.Debugf("getTravel")
	resp := response.New()
	defer resp.Send(ctx)

	obj, exists := ctx.Get(constants.AuthObject)
	if !exists {
		log.Warn("failed get auth object")
		resp.Error("please login first")
		return
	}

	travelID := ctx.Query("id")

	//parse user id
	uid := int(obj.(map[string]interface{})["id"].(float64))
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
		resp.Error("please login first")
		return
	}

}

func NewTravelController(s services.TravelService) Interface.IController {
	tc := &TravelController{
		service: s,
	}
	tc.Initialize()
	return tc
}
