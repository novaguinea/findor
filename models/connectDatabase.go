package models

import(
	// "fmt"
	"os"
	"github.com/novaguinea/findor/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

func SetupModels() *gorm.DB {

	errEnv := godotenv.Load()

	driver := os.Getenv("DRIVER")
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	db_name := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", "root:@(localhost)/findor_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Users)
	
	return db

}