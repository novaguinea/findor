package database

import(
	"fmt"
	// "os"
	"github.com/novaguinea/findor/models"
	// "github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func SetupModels() *gorm.DB {

	dsn := fmt.Sprintf("root:@(localhost:3306)/findor_db?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Users{})
	
	return db

}