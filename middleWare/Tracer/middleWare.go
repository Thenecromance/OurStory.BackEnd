package Tracer

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func createNewTraceId() string {

	return uuid.NewV4().String()
}

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("trace_id", createNewTraceId())
		ctx.Next()
	}
}
