package controllers

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)



func CreateUser(c *gin.Context) {
	var user models.User
	// c.BindJSON(user)
	value, err := ioutil.ReadAll(body)
	if err != nil{
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"User": value.username})
}