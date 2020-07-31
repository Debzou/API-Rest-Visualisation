// define package name
package controllers

// import library
import(
	"time"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/Debzou/REST-API-GO/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// DATABASE INSTANCE
var collection2 *mongo.Collection

// define the collection
func ConsumptionCollection(c *mongo.Database) {
	collection2 = c.Collection("consumption")
}

// Post Ecolyo consumption in mongodb
func PostConsumption(c *gin.Context) {
	var json models.Consumption
	c.Bind(&json)	
	ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection2.InsertOne(ctx, json)
	// display message & httpstatus
	c.JSON(http.StatusOK, gin.H{"message": "An user consumption is posted"})
	return				 
}