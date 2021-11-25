package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"github.com/novaguinea/findor/database"
)

func main()  {
	
	r := gin.Default()

	db := database.SetupModels()

	r.GET("/", func(c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{"status":"hehe hello"})
	})

	r.Run()
}