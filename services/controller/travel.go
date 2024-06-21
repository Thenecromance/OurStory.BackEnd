package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/middleware/Authorization"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/services"
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

func (tc *TravelController) SetupRouters() {
	mw := JWT.Middleware()
	tc.groups.travel = route.NewREST("/api/travel/:id")
	{
		tc.groups.travel.SetMiddleWare(mw)
		tc.groups.travel.SetHandler(
			tc.getTravel,    // GET
			tc.createTravel, // POST
			tc.updateTravel, // PUT
			tc.deleteTravel, // DELETE
		)
	}
	tc.groups.travelList = route.NewRouter("/api/travels", "POST")
	{
		tc.groups.travelList.SetMiddleWare(mw)
		tc.groups.travelList.SetHandler(tc.getTravelList)
	}
}

func (tc *TravelController) GetRoutes() []Interface.IRoute {
	return []Interface.IRoute{tc.groups.travel, tc.groups.travelList}
}

func (tc *TravelController) collectParams(ctx *gin.Context) (travelId string, user models.UserClaim, success bool) {
	val, exists := ctx.Get(Authorization.AuthObject)
	if !exists {
		success = false
		return
	}

	success = true
	{
		// any to map
		m := val.(map[string]any)
		//user.Id =
		log.Info("user id: ", m["id"].(float64))
		user.Id = int(m["id"].(float64))
		user.UserName = m["username"].(string)
		/*user.Id = val.(map[string]string)["id"]*/
		//user = val.(models.UserClaim)
	}
	travelId = ctx.Param("id") // get id from url
	log.Info("user id: ", user.Id, " travel id: ", travelId, " user name: ", user)
	return

}

// getTravel get travel by id
func (tc *TravelController) getTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	tId, user, success := tc.collectParams(ctx)
	if !success {
		resp.Unauthorized("Unauthorized access")
		return
	}

	// get travel from db by id
	travel, err := tc.service.GetTravelByID(tId)
	if err != nil {
		log.Error(err)
		resp.Error(err.Error())
		return
	}

	// check if user is owner of travel
	// but this is not the best way to check if user is owner of travel
	// so this is a temp solution
	isOwner := func() bool {
		if travel.UserId == user.Id {
			return true
		}
		/*		for _, v := range travel.TogetherWith {
				if v == user.Id {
					return true
				}
			}*/
		return false

	}
	if !isOwner() {
		resp.Unauthorized("Unauthorized access")
		return
	}

	resp.Success(travel)

}

func (tc *TravelController) createTravel(ctx *gin.Context) {
	log.Debugf("Creating travel")
	resp := response.New()
	defer resp.Send(ctx)

	_, user, success := tc.collectParams(ctx)
	if !success {
		resp.Unauthorized("something goes wrong with your request")
		return
	}

	// get data from request
	var newTravel models.Travel
	err := ctx.ShouldBind(&newTravel)
	if err != nil {
		log.Error(err)
		resp.Error("something wrong with your request")
		return
	}
	newTravel.UserId = user.Id
	err = tc.service.CreateTravel(&newTravel)
	if err != nil {
		log.Error(err)
		resp.Error(err.Error())
		return
	}

	resp.Success("Create travel")
}

func (tc *TravelController) updateTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	_, user, success := tc.collectParams(ctx)
	if !success {
		resp.Unauthorized("Unauthorized access")
		return
	}
	var obj models.Travel
	if err := ctx.ShouldBind(&obj); err != nil {
		log.Error(err)
		resp.Error("something wrong with your request")
		return
	}

	if obj.UserId != user.Id {
		resp.Unauthorized("Unauthorized access")
		return
	}

	//tc.service.UpdateTravel(tId, user.Id, ctx)
}
func (tc *TravelController) deleteTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	_, user, success := tc.collectParams(ctx)
	if !success {
		resp.Unauthorized("Unauthorized access")
		return

	}

	tId := ctx.Param("id")
	travel, err := tc.service.GetTravelByID(tId)
	if err != nil {
		log.Error(err)
		resp.Error(err.Error())
		return
	}

	if travel.UserId != user.Id {
		resp.Unauthorized("Unauthorized access")
		return
	}

	err = tc.service.DeleteTravel(tId)
	if err != nil {
		log.Error(err)
		return
	}
	resp.Success("Delete travel")
}

//------------------------------------------------------------
//TravelList
//------------------------------------------------------------

func (tc *TravelController) getTravelList(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func NewTravelController(s services.TravelService) *TravelController {
	tc := &TravelController{
		service: s,
	}
	tc.SetupRouters()
	return tc
}
