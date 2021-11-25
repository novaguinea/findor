package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"github.com/novaguinea/findor/database"
	"github.com/novaguinea/findor/controllers"
)

func main()  {
	
	r := gin.Default()

	db := database.SetupModels()

	r.Use(func(c *gin.Context)  {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{"status":"hehe hello"})
	})

	r.GET("/user", controllers.GetUsers)
	r.POST("/user", controllers.AddUser)
	r.PUT("/user", controllers.EditUser)

	r.Run()
}