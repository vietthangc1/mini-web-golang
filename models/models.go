package models

import (
	"gorm.io/gorm"
)

type Propertises struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"constraint:OnDelete:CASCADE"`
	Color     string `json:"color"`
	Brand     string `json:"brand"`
	Size      string `json:"size"`
}

type Product struct {
	gorm.Model
	SKU         string      `gorm:"size:150" json:"sku"`
	Name        string      `gorm:"size:150" json:"name"`
	Price       float64     `json:"price"`
	Number      int64       `json:"number"`
	Description string      `json:"description"`
	Cate1       string      `json:"cate1"`
	Cate2       string      `json:"cate2"`
	Cate3       string      `json:"cate3"`
	Cate4       string      `json:"cate4"`
	Propertises Propertises `json:"propertises"`
	UserEmail   string      `gorm:"size:150" json:"user_email"`
}

type User struct {
	gorm.Model
	Email    string    `gorm:"unique;size:150" json:"email"`
	Password string    `json:"password"`
	Products []Product `gorm:"foreignKey:UserEmail;references:Email"`
}
