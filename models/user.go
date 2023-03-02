package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `gorm:"unique;size:150" json:"email"`
	Password string    `gorm:"check:password <> '';not null" json:"password"`
	Products []Product `gorm:"foreignKey:UserEmail;references:Email"`
}