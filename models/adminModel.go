package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Email    string `gorm:"type:VARCHAR(100);unique"`
	Password string
}
