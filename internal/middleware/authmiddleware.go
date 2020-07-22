// package name
package middleware


// library imported
import (		
	"time"	
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"	
	"github.com/Debzou/REST-API-GO/internal/controllers"
	"github.com/Debzou/REST-API-GO/internal/models"
)
// define the key
var IdentityKey = "id"
// the jwt middleware
var AuthMiddleware, Err = jwt.New(&jwt.GinJWTMiddleware{
	// jwt option
	Realm:       "DebzouAPI",
	Key:         []byte("DebzouKeyDebzouCorp"),
	Timeout:     time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: IdentityKey,
	// define value in jwt
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*models.TokenInfoUser); ok {
			return jwt.MapClaims{
				IdentityKey: v.Status,
			}
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &models.TokenInfoUser{
			Status: claims[IdentityKey].(string),
		}
	},
	// authenticate , give a jwt
	// token info : status
	Authenticator: func(c *gin.Context) (interface{}, error) {
		var loginVals models.Login
		c.Bind(&loginVals) // This will infer what binder to use depending on the content-type header.
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