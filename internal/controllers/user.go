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
	// gather username
	username := c.PostForm("username")
	// create with models an user
	user := models.User{Username: username,
	Password: c.PostForm("password"),
	Status: "normal_user"}
	// check if username exist
	if (isExist(username)){ 
		// username already exist
		c.JSON(http.StatusOK, gin.H{"message": "User already exist"})
		return
	}else{
		// post user in mongodb (username no exist)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		collection.InsertOne(ctx, user)
		// display message & httpstatus
		c.JSON(http.StatusOK, gin.H{"message": "An user was been created with status : " + user.Status})
		return
	}		
}

func AuthUser(username string,password string) bool{
	// init user structure
	user := models.User{}
	// define the context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// find the user with username 
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return false
	}
	// check if password is good
	return true
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
		log.Printf("Error, Reason: %v\n", err)
		return false
	// the user not exist
	}else{
		return true
	}
}