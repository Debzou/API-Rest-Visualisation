// package name
package controllers

// import library
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func Getindex(c *gin.Context) {
	c.HTML(http.StatusOK, "Index.tmpl", gin.H{
	"Page": "Home",
	})
}

func GetGeoVis(c *gin.Context) {
	c.HTML(http.StatusOK, "Index.tmpl", gin.H{
	"Page": "GeoVis",
	})
}