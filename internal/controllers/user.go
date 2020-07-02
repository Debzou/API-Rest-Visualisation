package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)


type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
	Status   string             `json:"status,omitempty"`
}


func CreateUser(c *gin.Context) {
	var user User
    c.BindJSON(user)
	c.JSON(http.StatusOK, gin.H{"User": user.username})
}