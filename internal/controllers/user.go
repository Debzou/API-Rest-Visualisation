package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)



func CreateUser(c *gin.Context) {
	// create with models an user
	todo := models.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Status: c.PostForm("status")
	}
	c.JSON(http.StatusOK, gin.H{"User": todo.Username})
}