package models

import (
	"gorm.io/gorm"
)

type Propertises struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"constraint:OnDelete:CASCADE;OnDelete:CASCADE"`
	Attribute string `gorm:"not null" json:"attribute"`
	Value     string `gorm:"not null" json:"value"`
}