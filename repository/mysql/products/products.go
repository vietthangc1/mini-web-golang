package products

import (
	"log"
	"net/url"

	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/repository"
	"github.com/vietthangc1/mini-web-golang/utils"
	"gorm.io/gorm"
)

type ProductRepoImpl struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repository.ProductRepo {
	return &ProductRepoImpl{db: db}
}

func (r *ProductRepoImpl) GetProductByID(id uint) (models.Product, error) {
	var productQuery models.Product
	err := r.db.Preload("Propertises").Where("products.id = ?", id).First(&productQuery).Error

	if err != nil {
		return models.Product{}, err
	}
	return productQuery, nil
}

func (r *ProductRepoImpl) GetProducts(filter url.Values) ([]models.Product, error) {
	
	arrayProductFilter := []string{"cate1", "cate2", "cate3", "cate4"}
	productFilter := make(map[string]interface{})
	propertisesFilter := make(map[string]interface{})
	for k, v := range filter {
		if utils.Contains(arrayProductFilter, k) {
			productFilter[k] = v
		} else {
			propertisesFilter[k] = v
		}
	}

	log.Println(productFilter)
	log.Println(propertisesFilter)
	
	var lst_id []uint
	q := r.db.
		Distinct("products.id").
		Table("products").
		Joins("inner join propertises on products.id = propertises.product_id")

	// filter on propertises
	for key, element := range propertisesFilter {
		q = q.Where("propertises.Attribute = ? AND propertises.Value = ?", key, element)
	}	
	// filter on products
	q = q.Where(productFilter)
	_ = q.Find(&lst_id)
	log.Println(lst_id)

	// _q := r.db.
	// 	Distinct("products.id").
	// 	Table("products")
	// for key, element := range propertisesFilter {
	// 	_q = _q.Where("exist(?)", )
	// 	_q = _q.Where("propertises.Attribute = ? AND propertises.Value = ?", key, element)
	// }	

	var productsQuery []models.Product
	query := r.db.
		Preload("Propertises").
		Where("id in ?", lst_id)

	err := query.Find(&productsQuery).Error

	if err != nil {
		return []models.Product{}, err
	}
	return productsQuery, nil
}

func (r *ProductRepoImpl) AddProduct(newProduct models.Product) (models.Product, error) {
	err := r.db.Create(&newProduct).Error
	if err != nil {
		return models.Product{}, err
	}
	return newProduct, nil
}

func (r *ProductRepoImpl) UpdateProduct(updateProduct models.Product, id uint) (models.Product, error) {
	err := r.db.Model(&models.Product{}).Where("id = ?", id).Updates(updateProduct).Error
	if err != nil {
		log.Println(err)
		return models.Product{}, err
	}

	err = r.db.Model(&models.Propertises{}).Where("product_id = ?", id).Updates(updateProduct.Propertises).Error
	if err != nil {
		return models.Product{}, err
	}
	return updateProduct, nil
}

func (r *ProductRepoImpl) DeleteProduct(id uint) (models.Product, error) {
	var productDelete models.Product
	r.db.Where("id = ?", id).Delete(&productDelete)
	r.db.Where("product_id = ?", id).Delete(&models.Propertises{})
	return productDelete, nil
}
