package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)

func CreateUser(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil{
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"User": string(value)})
}