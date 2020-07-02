package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)



func CreateUser(c *gin.Context) {
	todo := models.User{Username: c.PostForm("username")}
	c.JSON(http.StatusOK, gin.H{"User": todo.Username})
}