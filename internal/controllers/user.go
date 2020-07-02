package controllers

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)



func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"User": ""})
}