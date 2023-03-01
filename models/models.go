package models

import (
	"time"

	"gorm.io/gorm"
)

type Propertises struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"constraint:OnDelete:CASCADE;OnDelete:CASCADE"`
	Attribute string `gorm:"not null" json:"attribute"`
	Value     string `gorm:"not null" json:"value"`
}

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

type User struct {
	gorm.Model
	Email    string    `gorm:"unique;size:150" json:"email"`
	Password string    `gorm:"check:password <> '';not null" json:"password"`
	Products []Product `gorm:"foreignKey:UserEmail;references:Email"`
}

type Log struct {
	ID uint `gorm:"primaryKey"`
	UserEmail string
	Table     string
	EntityID  uint64
	OldValue  string
	NewValue  string
	Timestamp time.Time `gorm:"autoUpdateTime"`
}

func (p *Product) BeforeUpdate(db *gorm.DB) error {
	// var oldRecord Product

	// db.Model(Product{}).Where("id = ?", p.ID).First(&oldRecord)

	// oldRecordJSON, _ := json.Marshal(oldRecord)
	// newRecordJSON, _ := json.Marshal(p)

	// log := Log{
	// 	Table: db.Statement.Table,
	// 	EntityID: uint(db.Statement.ReflectValue.FieldByName("ID").Uint()),
	// 	OldValue: string(oldRecordJSON),
	// 	NewValue: string(newRecordJSON),
	// }
	// if err := db.Create(&log).Error; err != nil {
	// 	return err
	// }
	return nil
}
