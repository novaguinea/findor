package database

import(
	"fmt"
	"os"
	"github.com/novaguinea/findor/models"
	// "github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func SetupModels() *gorm.DB {

	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	db_name := os.Getenv("DB_NAME")
	var port = "8080"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Users{})
	
	return db

}