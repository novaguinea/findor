package main

import(
	"net/http"
	"time"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt/v2"
	// "gorm.io/gorm"
	"github.com/novaguinea/findor/database"
	"github.com/novaguinea/findor/controllers"
)

type User struct{
	ID			string		`json: "id"`
	Email		string		`json: "email"`
	Name		string		`json: "name"`
	Password	string		`json: password`
}

type login struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"ID":claims[identityKey],
		"email":user.(*User).Email,
		"name":"Mamangs",
	})
}

func main()  {
	
	r := gin.Default()

	db := database.SetupModels()

	// the jwt middleware
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
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			if (email == "admin" && password == "admin") || (email == "test" && password == "test") {
				return &User{
					Email:  email,
					Name: "Mamangs",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.Email == "admin" {
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
		auth.GET("/hello", func(c *gin.Context)  {
			c.JSON(http.StatusOK, gin.H{"status":"Successfully use JWT :D"})
		})
	}


	r.Use(func(c *gin.Context)  {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{"status":"hehe hello"})
	})

	//router for users
	r.GET("/users", controllers.GetUsers)
	r.POST("/user", controllers.AddUser)
	r.PUT("/user/:id", controllers.EditUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}