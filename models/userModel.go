package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string `gorm:"type:VARCHAR(100);unique"`
	Password  string
	Status    bool
}

type JwtSessionList struct {
	ID          uint    `gorm:"primarykey"`
	TokenString string  `gorm:"not null"`
	EndTime     float64 `gorm:"not null"`
}
