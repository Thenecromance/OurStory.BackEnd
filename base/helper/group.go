package helper

import (
	"github.com/Thenecromance/OurStories/base/utils"
	"github.com/gin-gonic/gin"
)

func RegisterPackageAsGroup(r *gin.Engine, fnc interface{}) *gin.RouterGroup {
	return r.Group("/" + utils.GetPackageName(fnc))
}
