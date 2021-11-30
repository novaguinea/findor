package controllers

import(
	"net/http"
	"github.com/novaguinea/findor/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct{
	ID			uint32		`json: "id"`
	Email		string		`json: "email"`
	Name		string		`json: "name"`
	Password	string		`json: "password"`
	ConfirmPwd	string		`json: "confirmPwd"`
	Address		string		`json: "address"`
	Skill		string		`json: "skill"`
	Phone		string		`json: "phone"`
	Age			string			`json: "age"`
	IsAvailable	string		`json: "isAvailable"`
	AvatarURL	string		`json: "avatarURL"`
}

type UserShow struct{
	ID			uint32		`json: "id"`
	Email		string		`json: "email"`
	Name		string		`json: "name"`
	Address		string		`json: "address"`
	Skill		string		`json: "skill"`
	Phone		string		`json: "phone"`
	Age			string			`json: "age"`
	IsAvailable	string		`json: "isAvailable"`
	AvatarURL	string		`json: "avatarURL"`
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

func GetUser(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var user []models.Users

	db.Find(&user, c.Param("id"))

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
		return
	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message":"Email wajib diisi"})
		return
	}

	if data.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message":"Nama wajib diisi"})
		return
	}

	if data.Password != data.ConfirmPwd {
		c.JSON(http.StatusBadRequest, gin.H{"Message":"Password tidak sama"})
		return
	}

	//input to db
	user := models.Users{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message":"user successfully added"})

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
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		Address: data.Address,
		Skill: data.Skill,
		Phone: data.Phone,
		Age: data.Age,
		IsAvailable: data.IsAvailable,
		AvatarURL: data.AvatarURL,
	}

	db.Model(&user).Where("id = ?", c.Param("id")).Updates(data)

	c.JSON(http.StatusOK, gin.H{"Message":"Data successfully updated"})

}

func DeleteUser(c *gin.Context)  {

	db := c.MustGet("db").(*gorm.DB)

	//check user if exist
	var u models.Users

	if err := db.Where("id = ?", c.Param("id")).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"User not found"})
		return
	}

	db.Delete(&u)
	
	c.JSON(http.StatusOK, gin.H{"Status":true})

}