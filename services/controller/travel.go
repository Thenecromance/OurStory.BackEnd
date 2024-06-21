package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/middleware/Authorization"
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
	tc.groups.travel = route.NewREST("/api/travel/:id")
	{
		tc.groups.travel.SetHandler(
			tc.getTravel,    // GET
			tc.createTravel, // POST
			tc.updateTravel, // PUT
			tc.deleteTravel, // DELETE
		)
	}
	tc.groups.travelList = route.NewRouter("/api/travels", "POST")
	{
		tc.groups.travelList.SetHandler(tc.getTravelList)
	}
}

func (tc *TravelController) collectParams(ctx *gin.Context) (travelId string, user models.UserClaim, success bool) {
	val, exists := ctx.Get(Authorization.AuthObject)
	if !exists {
		success = false
		return
	}
	success = true
	user = val.(models.UserClaim)
	travelId = ctx.Param("id") // get id from url
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

	travel, err := tc.service.GetTravel(tId)
	if err != nil {
		log.Error(err)
		resp.Error(err.Error())
		return
	}

	isOwner := func() bool {
		if travel.UserId == user.Id {
			return true
		}
		for _, v := range travel.TogetherWith {
			if v == user.Id {
				return true
			}
		}
		return false

	}
	if !isOwner() {
		resp.Unauthorized("Unauthorized access")
		return
	}

	resp.Success(travel)

}

func (tc *TravelController) createTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	_, user, success := tc.collectParams(ctx)
	if !success {
		resp.Unauthorized("Unauthorized access")
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

	tId, user, success := tc.collectParams(ctx)
	if !success {
		resp.Unauthorized("Unauthorized access")
		return
	}

	tc.service.UpdateTravel(tId, user.Id, ctx)
}
func (tc *TravelController) deleteTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

//------------------------------------------------------------
//TravelList
//------------------------------------------------------------

func (tc *TravelController) getTravelList(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}
