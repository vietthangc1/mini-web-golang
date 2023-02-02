package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HandlerGetProductByID(c *gin.Context) {
	id := c.Param("id")
	var (
		productQuery product
	)
	query := `
	SELECT id, sku, name, price, number, description, cate1, cate2, coalesce(cate3, '') as cate3, coalesce(cate4, '') as cate4, propertises
	FROM products 
	WHERE id = ?
	`
	productQuery, err := QueryGetProductByID(query, id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusFound, productQuery)
}

func HandlerGetProducts(c *gin.Context) {
	var (
		productsQuery []product
	)

	filterProducts := c.Request.URL.Query()

	cate1 := filterProducts.Get("cate1")
	if cate1 == "" {
		cate1 = "%%"
	}
	cate2 := filterProducts.Get("cate2")
	if cate2 == "" {
		cate2 = "%%"
	}
	cate3 := filterProducts.Get("cate3")
	if cate3 == "" {
		cate3 = "%%"
	}
	cate4 := filterProducts.Get("cate4")
	if cate4 == "" {
		cate4 = "%%"
	}

	query := `
	SELECT id, sku, name, price, number, description, cate1, cate2, coalesce(cate3, '') as cate3, coalesce(cate4, '') as cate4, propertises
	FROM products 
	WHERE 1=1
	AND cate1 like ?
	AND cate2 like ?
	AND cate3 like ?
	AND cate4 like ?
	`
	productsQuery, err := QueryGetProducts(query, cate1, cate2, cate3, cate4)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Pagination
	
	productPerPage := 2
	pageNum := 1
	
	page := filterProducts.Get("page")	
	if (page != "") {
		pageNum, _ = strconv.Atoi(page) 
	}

	positionStart := (pageNum-1)*productPerPage
	positionEnd := int(math.Min(float64(pageNum*productPerPage), float64(len(productsQuery))))

	productsPagination := []product{}
	if (positionStart < positionEnd) {
		productsPagination = productsQuery[positionStart:positionEnd]
	}

	c.IndentedJSON(http.StatusFound, productsPagination)
}

func HandlerAddProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	id := time.Now().UnixMilli()
	newProduct.ID = strconv.Itoa(int(id))

	query := "INSERT INTO products ( id, sku, name, price, number, description, cate1, cate2, cate3, cate4, propertises) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	
	newProduct, err := QueryAddProduct(query, newProduct)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func HandlerUpdateProduct(c *gin.Context) {
	var updateProduct product

	id := c.Param("id")

	if err := c.BindJSON(&updateProduct); err != nil {
		return
	}

	query := `
	UPDATE products 
	SET sku = ?, name = ?, price = ?, number = ?, description = ?, cate1 = ?, cate2 = ?, cate3 = ?, cate4 = ?, propertises = ?
	WHERE id = ?
	`

	updateProduct, err := QueryUpdateProduct(query, id, updateProduct)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusCreated, updateProduct)
}

func HandlerDeleteProduct(c *gin.Context) {
	id := c.Param("id")

	query := `
	DELETE FROM products 
	WHERE id = ?
	`

	err := QueryDeleteProduct(query, id)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Deleted!"})
}