package controllers

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	
)

// jwt user intentification
var IdentityKey = "id"

// After auth, display a message
func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"status":   claims[IdentityKey],
		"text":     "Welcome is api. You are authenticate",
	})
}
