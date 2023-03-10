package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/tokens"
)

func (h *Handler) HandlerGetProductByID(c *gin.Context) {
	id := c.Param("id")
	var productQuery models.Product

	val, err := h.CacheInstance.Get(id)
	if err != nil {
		_id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			log.Println(err.Error())
			c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
			return
		}
		productQuery, err = h.ProductRepo.GetProductByID(uint(_id))
		if err != nil {
			log.Println(err.Error())
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		err = h.CacheInstance.Set(id, productQuery)
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

func (h *Handler) HandlerGetProducts(c *gin.Context) {
	var (
		productsQuery []models.Product
	)

	filter := c.Request.URL.Query()

	productsQuery, err := h.ProductRepo.GetProducts(filter)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusFound, productsQuery)
}

func (h *Handler) HandlerAddProduct(c *gin.Context) {
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

	newProduct, err = h.ProductRepo.AddProduct(newProduct)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newProduct)
}

func (h *Handler) HandlerUpdateProduct(c *gin.Context) {
	var updateProduct models.Product
	user_email, err := tokens.ExtractTokenEmail(c)
	if err != nil {
		
		c.IndentedJSON(401, gin.H{"error": err.Error()})
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

	oldRecord, err := h.ProductRepo.GetProductByID(uint(_id))
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	updateProduct.ID = uint(_id)

	updateProduct, err = h.ProductRepo.UpdateProduct(updateProduct, uint(_id), user_email)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}

	productQuery, err := h.ProductRepo.GetProductByID(uint(_id))
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	oldRecordJSON, _ := json.Marshal(oldRecord)
	newRecordJSON, _ := json.Marshal(productQuery)

	if string(oldRecordJSON) != string(newRecordJSON) {
		newLog := models.Log{
			UserEmail: user_email,
			TableModel: "Products",
			EntityID: _id,
			OldValue: string(oldRecordJSON),
			NewValue: string(newRecordJSON),
		}
		
		_, err = h.LogRepo.AddLog(newLog)
		if err != nil {
			log.Println(err.Error())
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	}

	err = h.CacheInstance.Set(id, productQuery)
	if err != nil {
		log.Println("Cannot update cache")
		log.Println(err.Error())
	} else {
		log.Println("Updated to cache")
	}

	c.IndentedJSON(http.StatusOK, productQuery)
}

func (h *Handler) HandlerDeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}

	_, err = h.ProductRepo.DeleteProduct(uint(_id))
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": err.Error()})
		return
	}

	err = h.CacheInstance.Delete(id)
	if err != nil {
		log.Println("Cannot delete from cache")
		log.Println(err.Error())
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted!"})
}
