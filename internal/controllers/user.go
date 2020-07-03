package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)



func CreateUser(c *gin.Context,database *mongo.Database) {
	// create with models an user
	user := models.User{Username: c.PostForm("username"),
	Password: c.PostForm("password"),
	Status: c.PostForm("status")}
	// post data in mongodb
	collection := database.Collection("podcastsv2")
	collection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})
	c.JSON(http.StatusOK, gin.H{"User": user.Username})
}