package initializers

import (
	"fmt"
	"os"

	"github.com/Noush-012/web_jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// To connect database
func ConnToDB() {
	dsn := os.Getenv("DATABASE")

	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Faild to Connect Database")
		return
	}
	fmt.Println("Successfully Connected to database")
}

// Schema Migrate to database
func MigrateToDB() {

	if DB.AutoMigrate(&models.Admin{}, &models.User{}, &models.JwtSessionList{}); err != nil {
		fmt.Println("faild to sync database")
		return
	}
	fmt.Println("Successfully synced to database")
}
