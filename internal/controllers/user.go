package controllers

import(
	"time"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)


// DATABASE INSTANCE
var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

func CreateUser(c *gin.Context) {
	// create with models an user
	user := models.User{Username: c.PostForm("username"),
	Password: c.PostForm("password"),
	Status: c.PostForm("status")}
	// post data in mongodb
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection.InsertOne(ctx, user)
	c.JSON(http.StatusOK, gin.H{"user status": user.Username})
}