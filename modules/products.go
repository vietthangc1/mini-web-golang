package modules

import (

	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/gorm"
)

// ORM models

type ProductRepository struct{
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProductByID (id uint) (models.Product, error) {
	var productQuery models.Product
	err := r.db.Joins("Propertises").Where("products.id = ?", id).First(&productQuery).Error

	if err != nil {
		return models.Product{}, err
	}
	return productQuery, nil
}

func (r *ProductRepository) GetProducts(productFilter, propertisesFilter map[string]interface{}) ([]models.Product, error) {
	var productsQuery []models.Product
	err := r.db.
		Joins("Propertises", r.db.Where(propertisesFilter)).
		Where(productFilter).
		Find(&productsQuery).Error
	if err != nil {
		return []models.Product{}, err
	}
	return productsQuery, nil
}

func (r *ProductRepository) AddProduct(newProduct models.Product) (models.Product, error) {
	err := r.db.Create(&newProduct).Error
	if err != nil {
		return models.Product{}, err
	}
	return newProduct, nil
}

func (r *ProductRepository) UpdateProduct(updateProduct models.Product, id uint) (models.Product, error) {
	err := r.db.Model(&models.Product{}).Where("id = ?", id).Updates(updateProduct).Error
	if err != nil {
		return models.Product{}, err
	}
	err = r.db.Model(&models.Propertises{}).Where("product_id = ?", id).Updates(updateProduct.Propertises).Error
	if err != nil {
		return models.Product{}, err
	}
	return updateProduct, nil
}

func (r *ProductRepository) DeleteProduct(id uint) (models.Product, error) {
	var productDelete models.Product
	r.db.Where("id = ?", id).Delete(&productDelete)
	r.db.Where("product_id = ?", id).Delete(&models.Propertises{})
	return productDelete, nil
}