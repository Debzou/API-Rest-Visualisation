// define package name
package controllers

// import library
import(
	"time"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"net/http"
)


// DATABASE INSTANCE
var collection *mongo.Collection

// define the collection
func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

func CreateUser(c *gin.Context) {
	var json models.User
	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	// gather username and transform to lower case
	username := strings.ToLower(json.Username)
	//hash password
	hashpassword,_ := HashPassword(json.Password)
	log.Printf(hashpassword)
	// create with models an user
	user := models.User{Username: username,
		Password: hashpassword,
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
			c.JSON(http.StatusOK, gin.H{"message": "User is created"})
			return
		}		 
}

// return true if authenticate is true and the status value
func AuthUser(username string,password string) (bool,string){
	// init user structure
	user := models.User{}
	// define the context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// check if user exist
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return false,"no_status"
	}
	// check the password
	if !(CheckPasswordHash(password, user.Password)) {
		log.Printf("incorrect password")
		return false,"incorrect_pass"
	}
	return true,user.Status
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
		log.Printf("create an user")
		return false
	// the user not exist
	}else{
		return true
	}
}

// hash password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// compare hash password with password
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}