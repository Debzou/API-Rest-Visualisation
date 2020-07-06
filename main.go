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
	"github.com/Debzou/REST-API-GO/internal/models"
)

// define the key
var identityKey = "id"

// define client mongo
var client *mongo.Client


func main() {
	fmt.Println("Starting the application...")
	// mongodb context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// define the mongo client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	// defer client.Disconnect(ctx)
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
		port = "8000"
	}
	// ROUTE NOT PROTECTED
	r.POST("/signup", controllers.CreateUser)

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		// jwt option
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// define value in jwt
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.TokenInfoUser); ok {
				return jwt.MapClaims{
					identityKey: v.Status,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.TokenInfoUser{
				Status: claims[identityKey].(string),
			}
		},
		// authenticate , give a jwt
		// token info : status
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password
			condition,status := controllers.AuthUser(username,password)
			if (condition) {
				return &models.TokenInfoUser{
					Status:    status,
					UserName:  username,
					
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// authorization 
		// if the status is not admin, then you won't be able to use the protected routes.
		Authorizator: func(data interface{}, c *gin.Context) bool {
			v, ok := data.(*models.TokenInfoUser)
			if ok && v.Status == "admin" {
				return true
			}			
			return false
		},
		// else
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// route auth
	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// PROTECTED ROUTE
		auth.GET("/hello", controllers.HelloHandler)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}