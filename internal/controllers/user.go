package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
)

func CreateUser(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"User": json.Username})
}