package controllers

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var IdentityKey = "id"

// After auth, display a message
func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"text":     "Hello World.",
	})
}