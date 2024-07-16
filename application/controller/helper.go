package controller

import (
	"encoding/json"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/gin-gonic/gin"
)

func getAuthObject(ctx *gin.Context) *models.UserClaim {
	Jsonclaim, exists := ctx.Get(constants.AuthObject)
	if !exists {
		return nil
	}

	claim := &models.UserClaim{}
	if err := json.Unmarshal([]byte(Jsonclaim.(string)), claim); err != nil {
		return nil
	}
	return claim

}
