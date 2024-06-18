package controller

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/server/router"
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
	tc.groups.travel = router.NewREST()
	{
		tc.groups.travel.SetPath("/api/travel/:id")
		tc.groups.travel.SetHandler(
			tc.getTravel,    // GET
			tc.createTravel, // POST
			tc.updateTravel, // PUT
			tc.deleteTravel, // DELETE
		)
	}
	tc.groups.travelList = router.NewRouter()
	{
		tc.groups.travelList.SetPath("/api/travel")
		tc.groups.travelList.SetHandler(tc.getTravelList)
	}
}

func (tc *TravelController) getTravel(ctx *gin.Context) {

}

func (tc *TravelController) createTravel(ctx *gin.Context) {

}

func (tc *TravelController) updateTravel(ctx *gin.Context) {

}
func (tc *TravelController) deleteTravel(ctx *gin.Context) {

}

//------------------------------------------------------------
//TravelList
//------------------------------------------------------------

func (tc *TravelController) getTravelList(ctx *gin.Context) {

}
