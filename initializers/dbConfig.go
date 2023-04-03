package initializers

import (
	"fmt"
	"os"

	"github.com/Noush-012/web_jwt/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	if DB.AutoMigrate(&models.Admin{}, &models.User{}); err != nil {
		fmt.Println("faild to sync database")
		return
	}
	fmt.Println("Successfully synced to database")
}

func CreateAdmin() {
	ADMINID := os.Getenv("ADMINID")
	ADMINPASS := os.Getenv("ADMINPASS")

	//hash the password and if no error create admin
	if hashPass, err := bcrypt.GenerateFromPassword([]byte(ADMINPASS), 10); err == nil {

		DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Admin{
			Email:    ADMINID,
			Password: string(hashPass),
		})

		fmt.Println("Admin created successful")
	}
}
