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
	"net/http"
	"log"
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

func GetAllConsumption(c *gin.Context) {
	// init table consumption structure
	var results []models.Consumption
	var consumption models.Consumption
	// define the context
	ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()  // releases resources if slowOperation completes before timeout elapses
	// find all consumption
	cursor, err := collection2.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "err"})
		panic(err)		
    } else {
		for cursor.Next(ctx) {
			// decode the document
			if err := cursor.Decode(&consumption); err != nil {
				log.Fatal(err)
			}
			// add in table
			results =append(results, consumption)
		}
		c.JSON(http.StatusOK, gin.H{"data": results})
	}

    		 
}

