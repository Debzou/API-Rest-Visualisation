package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)

func CreateUser(c *gin.Context) {
	var user LOGIN
    c.BindJSON($user)
	c.JSON(http.StatusOK, gin.H{"User": user.username})
}