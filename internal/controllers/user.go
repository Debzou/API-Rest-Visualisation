// define package name
package controllers

// import library
import(
	"time"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)


// DATABASE INSTANCE
var collection *mongo.Collection

// define the collection
func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

func CreateUser(c *gin.Context) {
	// create with models an user
	user := models.User{Username: c.PostForm("username"),
	Password: c.PostForm("password"),
	Status: "normal_user"}
	// post user in mongodb
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection.InsertOne(ctx, user)
	// display message & httpstatus
	c.JSON(http.StatusOK, gin.H{"message": "A user was been created with status : " + user.Status})
	return
}

func AuthUser(username string,password string){
	// init user structure
	user := models.User{}
	// define the context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// find the user with username 
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return
	}
	// check if password is good
	return
}

func isExist(username string) bool{
	// init user structure
	user := models.User{}
	// define the context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// find an user 
	err := collection.FindOne(ctx,bson.M{"username": username}).Decode(&user)
	// the user exist
	if err != nil {
		return true
	// the user not exist
	}else{
		return false
	}
}