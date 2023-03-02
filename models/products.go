package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	SKU         string        `gorm:"not null;size:150;check:sku <> ''" json:"sku"`
	Name        string        `gorm:"not null;size:150;check:name <> ''" json:"name"`
	Price       float64       `gorm:"not null;check:price > 0" json:"price"`
	Number      int64         `gorm:"not null;check:number > 0" json:"number"`
	Description string        `json:"description"`
	Cate1       string        `json:"cate1"`
	Cate2       string        `json:"cate2"`
	Cate3       string        `json:"cate3"`
	Cate4       string        `json:"cate4"`
	Propertises []Propertises `json:"propertises" gorm:"foreignKey:ProductID;references:ID"`
	UserEmail   string        `gorm:"size:150" json:"user_email"`
}