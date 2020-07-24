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
	fmt.Println("Starting the application ...<3")
	// mongodb context
	ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// define the mongo client
	// without docker you must be use this url --> mongodb://127.0.0.1:27017
	// with docker you must be use this url --> mongodb://mongo:27017/
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
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
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// view 
	router.LoadHTMLGlob("internal/views/*.tmpl")
	// servers other static files
	router.Static("/static", "./static")
	// if port is not define
	if port == "" {
		port = "8080"
	} 
	// ROUTE NOT PROTECTED
	router.POST("/signup", controllers.CreateUser)
	router.GET("/datavis/index",controllers.Getindex)
	router.GET("/datavis/geovis",controllers.GetGeoVis)
	if middleware.Err != nil {
		log.Fatal("JWT Error:" + middleware.Err.Error())
	}
	// route auth
	router.POST("/login", middleware.AuthMiddleware.LoginHandler)

	router.NoRoute(middleware.AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		// PROTECTED ROUTE
		auth.GET("/hello", controllers.HelloHandler)
	}

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}