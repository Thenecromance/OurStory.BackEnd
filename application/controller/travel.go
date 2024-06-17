package controller

import "github.com/gin-gonic/gin"

type TravelController struct {
}

func (tc *TravelController) RegisterRoutes(engine *gin.Engine) {
	travelGroup := engine.Group("/travel")
	{
		travelGroup.POST("/create", tc.createTravel)
		travelGroup.GET("/list", tc.getTravelList)
	}
}
