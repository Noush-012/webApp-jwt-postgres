package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	Email    string `gorm:"type:VARCHAR(100);unique"`
	Password string
}

// CheckPassword checks if the given password matches the user's hashed password
func (A *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(A.Password), []byte(password))
	return err == nil
}
