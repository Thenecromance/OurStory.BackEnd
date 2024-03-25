package helper

import (
	"github.com/Thenecromance/OurStories/base/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// RegisterMethodToGroup will register a method to a group, and the method name should be like "get_method" or "post_method"
func RegisterMethodToGroup(r *gin.RouterGroup, fnc gin.HandlerFunc) gin.IRoutes {
	//get the function name and split it
	funcName := utils.GetFunctionName(fnc)

	split := strings.Split(funcName, "_")

	// if the function name is not in the format of "method_name", then return nil
	if len(split) != 2 {
		return nil
	}

	//upper case the method name and join the rest of the split to form the path
	method := strings.ToUpper(split[0])

	path := "/" + strings.Join(split[1:], "/")
	return r.Handle(method, path, fnc)
}
