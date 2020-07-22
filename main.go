// the main package
package main

// library imported
import (
	"log"
	"net/http"
	"os"
	"time"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"	
	"github.com/Debzou/REST-API-GO/internal/controllers"
	"github.com/Debzou/REST-API-GO/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


// define client mongo
var client *mongo.Client
var errMongo *mongo.Client


func main() {
	fmt.Println("Starting the application...")
	// mongodb context
	ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// define the mongo client
	// URL without docker : mongodb://127.0.0.1:27017
	// URL with docker : mongodb://mongo:27017/
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/")
	client, errMongo := mongo.Connect(ctx, clientOptions)
	// errMongo
	if errMongo != nil {
		panic(errMongo)
	}
	defer func() {
		if errMongo = client.Disconnect(ctx); errMongo != nil {
			panic(errMongo)
		}
	}()
	// Ping the primary
	if errMongo := client.Ping(ctx, readpref.Primary()); errMongo != nil {
		log.Fatal("err ping mongo")
		panic(errMongo)
	}
	database := client.Database("RESTapi")
	// define collection
	controllers.UserCollection(database)
	// Start Gin
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// if port is not define
	if port == "" {
		port = "8080"
	} 
	// ROUTE NOT PROTECTED
	r.POST("/signup", controllers.CreateUser)

	if middleware.Err != nil {
		log.Fatal("JWT Error:" + middleware.Err.Error())
	}
	// route auth
	r.POST("/login", middleware.AuthMiddleware.LoginHandler)

	r.NoRoute(middleware.AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		// PROTECTED ROUTE
		auth.GET("/hello", controllers.HelloHandler)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}