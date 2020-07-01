package main

import( 
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/controllers"
)

func main(){
	controllers.Welcome()
	router := gin.Default()
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/")
	router.Run(":8080")
	
}