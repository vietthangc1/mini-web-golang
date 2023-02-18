package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/modules"
	"github.com/vietthangc1/mini-web-golang/tokens"
)

func (a *App) HandlerGetProductByID(c *gin.Context) {
	id := c.Param("id")
	var productQuery models.Product

	val, err := a.CacheInstance.Get(id)
	if err != nil {
		_id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			log.Println(err.Error())
			c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
			return
		}
		productQuery, err = a.Handler.ProductRepo.GetProductByID(uint(_id))
		if err != nil {
			log.Println(err.Error())
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		err = a.CacheInstance.Set(id, productQuery)
		if err != nil {
			log.Println("Cannot update cache")
			log.Println(err.Error())
		} else {
			log.Println("Updated to cache")
		}
	} else {
		log.Println("Use cache")
		productQuery = val
	}
	c.IndentedJSON(http.StatusFound, productQuery)
}

func (a *App) HandlerGetProducts(c *gin.Context) {
	var (
		productsQuery []models.Product
	)

	filter := c.Request.URL.Query()

	arrayProductFilter := []string{"cate1", "cate2", "cate3", "cate4"}
	productFilter := make(map[string]interface{})
	for k, v := range filter {
		if modules.Contains(arrayProductFilter, k) {
			productFilter[k] = v
		}
	}

	arrayPropertisesFilter := []string{"color", "brand", "size"}
	propertisesFilter := make(map[string]interface{})
	for k, v := range filter {
		if modules.Contains(arrayPropertisesFilter, k) {
			propertisesFilter[k] = v
		}
	}

	log.Println(productFilter)
	log.Println(propertisesFilter)

	productsQuery, err := a.Handler.ProductRepo.GetProducts(productFilter, propertisesFilter)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusFound, productsQuery)
}

func (a *App) HandlerAddProduct(c *gin.Context) {
	var newProduct models.Product
	user_email, err := tokens.ExtractTokenEmail(c)
	if err != nil {
		c.IndentedJSON(http.StatusNonAuthoritativeInfo, gin.H{"error": err.Error()})
	}

	if err := c.BindJSON(&newProduct); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}
	newProduct.UserEmail = user_email

	newProduct, err = a.Handler.ProductRepo.AddProduct(newProduct)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}	

	c.IndentedJSON(http.StatusCreated, newProduct)
}

func (a *App) HandlerUpdateProduct(c *gin.Context) {
	var updateProduct models.Product
	user_email, err := tokens.ExtractTokenEmail(c)
	if err != nil {
		c.IndentedJSON(http.StatusNonAuthoritativeInfo, gin.H{"error": err.Error()})
	}

	id := c.Param("id")

	if err := c.BindJSON(&updateProduct); err != nil {
		return
	}
	updateProduct.UserEmail = user_email

	_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}
	updateProduct.ID = uint(_id)
	log.Println(updateProduct)

	updateProduct, err = a.Handler.ProductRepo.UpdateProduct(updateProduct, uint(_id))
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}

	err = a.CacheInstance.Set(id, updateProduct)
	if err != nil {
		log.Println("Cannot update cache")
		log.Println(err.Error())
	} else {
		log.Println("Updated to cache")
	}

	c.IndentedJSON(http.StatusOK, updateProduct)
}

func (a *App) HandlerDeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}
	
	deleteProduct, err := a.Handler.ProductRepo.DeleteProduct(uint(_id))
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}

	err = a.CacheInstance.Delete(id)
	if err != nil {
		log.Println("Cannot delete from cache")
		log.Println(err.Error())
	}
	c.IndentedJSON(http.StatusOK, gin.H{"product": deleteProduct, "message": "Deleted!"})
}
