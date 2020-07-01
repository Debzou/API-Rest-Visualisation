package main

import( 
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Debzou/REST-API-GO/internal/controllers"
)

func main(){
	// router := gin.Default()
	// // This handler will match /user/john but will not match /user/ or /user
	// router.GET("/", controllers.wellcome())
	// router.Run(":8080")
	controllers.wellcome()
}