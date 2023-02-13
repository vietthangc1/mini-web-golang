package modules

import (
	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/gorm"
)

// ORM models

func GetProductByID (db *gorm.DB, productQuery *models.Product , id uint) (error) {
	err := db.Joins("Propertises", db.Select([]string{"brand","size","color"})).Where("products.id = ?", id).First(productQuery).Error

	if err != nil {
		return err
	}
	return nil
}

func GetProducts(db *gorm.DB, productsQuery *[]models.Product, productFilter, propertisesFilter map[string]interface{}) (error) {
	// err := db.Preload("Propertises").Where(productFilter).Find(productsQuery).Error
	err := db.
		Joins("Propertises", db.Where(propertisesFilter)).
		Where(productFilter).
		Find(productsQuery).Error
	if err != nil {
		return err
	}
	return nil
}

func AddProduct(db *gorm.DB, newProduct *models.Product) (error) {
	err := db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, updateProduct *models.Product) (error) {
	err := db.Save(updateProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(db *gorm.DB, productDelete *models.Product, id uint) (error) {
	db.Where("id = ?", id).Delete(productDelete)
	return nil
}