package controllers

import(
	"time"
	"os"
	"log"

	"github.com/novaguinea/findor/models"

	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt/v2"
	"gorm.io/gorm"
)

var identityKey = "id"

type UserAuth struct{
	ID			string		`json: "id"`
	Email		string		`json: "email"`
	Name		string		`json: "name"`
	Password	string		`json: password`
}

type LoginData struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}



func Login(c *gin.Context)  {

	db := c.MustGet("db").(*gorm.DB)
	// r := gin.Default()
	var loginVals LoginData

	email := loginVals.Email
	password := loginVals.Password

	var user models.Users

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{

		Realm:       "Findor",
		Key:         []byte(os.Getenv("SECRET_KEY")),
		Timeout:     time.Hour*24,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			
			// if err := c.ShouldBindJSON(&loginVals); err != nil {
			// 	return "", jwt.ErrMissingLoginValues
			// }
			
			db.Where("email = ?", email).First(&user)

			if (email == user.Email && password == user.Password){
				return &UserAuth{
					Email:  user.Email,
					Name: user.Name,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*UserAuth); ok && v.Email == email {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	authMiddleware.LoginHandler(c)

}
