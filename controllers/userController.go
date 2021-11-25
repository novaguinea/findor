package controllers

import(
	"net/http"
	"github.com/novaguinea/findor/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct{
	Email		string		`json: "email"`
	Name		string		`json: "name"`
	Password	string		`json: password`
}

func GetUsers(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var user []models.Users

	db.Find(&user)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"User not found"})
	}

	c.JSON(http.StatusOK, gin.H{"data":user})
}

func AddUser(c *gin.Context)  {

	db := c.MustGet("db").(*gorm.DB)

	var data User

	if err:=c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
	}

	//input to db
	user := models.Users{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"Data":data})

}

func EditUser(c *gin.Context)  {

	db := c.MustGet("db").(*gorm.DB)

	//check user if exist
	var u models.Users

	if err := db.Where("id = ?", c.Param("id")).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"User not found"})
		return
	}

	var data User

	if err:=c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
	}

	//input to db
	user := models.Users{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}

	db.Model(&user).Where("id = ?", c.Param("id")).Updates(data)
	
	c.JSON(http.StatusOK, gin.H{"Data":data})

}