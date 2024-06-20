package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/router"
	"github.com/gin-gonic/gin"
)

type travelRouter struct {
	travel     Interface.Router
	travelList Interface.Router
}

type TravelController struct {
	groups travelRouter
}

func (tc *TravelController) SetupRouters() {
	tc.groups.travel = router.NewREST("/api/travel/:id")
	{
		tc.groups.travel.SetHandler(
			tc.getTravel,    // GET
			tc.createTravel, // POST
			tc.updateTravel, // PUT
			tc.deleteTravel, // DELETE
		)
	}
	tc.groups.travelList = router.NewRouter("/api/travel", "POST")
	{
		tc.groups.travelList.SetHandler(tc.getTravelList)
	}
}

func (tc *TravelController) getTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func (tc *TravelController) createTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func (tc *TravelController) updateTravel(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
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
