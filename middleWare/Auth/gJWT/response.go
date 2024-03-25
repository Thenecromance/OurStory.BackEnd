package gJWT

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func responseUnauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
	c.Abort()
}

func responseSuccess(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"message": msg})
	c.Next()
}
