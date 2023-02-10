package modules

import (
	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/gorm"
)

// ORM models

func GetProductByID (db *gorm.DB, productQuery *models.Product , id uint) (error) {
	err := db.Preload("Propertises").Where("id = ?", id).First(productQuery).Error

	if err != nil {
		return err
	}
	return nil
}

func GetProducts(db *gorm.DB, productsQuery *[]models.Product, productFilter, propertisesFilter map[string]interface{}) (error) {
	// err := db.Preload("Propertises").Where(productFilter).Find(productsQuery).Error
	err := db.
		Preload("Propertises").
		Model(&models.Product{}).
		Select("*").
		Joins("left join propertises on products.id = propertises.product_id").
		Where(productFilter).
		Where(propertisesFilter).
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