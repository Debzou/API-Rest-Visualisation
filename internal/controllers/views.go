// package name
package controllers

// import library
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func Getindex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	"title": "Main website",
	})
}