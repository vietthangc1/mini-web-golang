package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type properties struct {
	Color string  `json:"color"`
	Size  float64 `json:"size"`
	Brand string  `json:"brand"`
}

type product struct {
	ID          string  `json:"id"`
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Number      int64   `json:"number"`
	Description string  `json:"description"`
	Cate1       string  `json:"cate1"`
	Cate2       string  `json:"cate2"`
	Properties properties `json:"properties"`
}

var products = []product{
	{
		ID:          "12",
		SKU:         "1234",
		Name:        "Giày da Nam",
		Price:       10000,
		Number:      10,
		Description: "Mô tả nè",
		Cate1:       "Thời trang",
		Cate2:       "Thời trang Nam",
		Properties: properties{
			Color: "Red",
		},
	},
	{
		ID:          "21",
		SKU:         "123",
		Name:        "Giày da Nữ",
		Price:       12000,
		Number:      10,
		Description: "Mô tả giày da Nữ nè",
		Cate1:       "Thời trang",
		Cate2:       "Thời trang Nữ",
		Properties: properties{
			Size: 12.3,
		},
	},
	{
		ID:          "159",
		SKU:         "123",
		Name:        "Quần da Nữ",
		Price:       12000,
		Number:      10,
		Description: "Mô tả quần da Nữ nè",
		Cate1:       "Thời trang",
		Cate2:       "Thời trang Nữ",
		Properties: properties{
			Size: 12,
		},
	},
	{
		ID:          "21569",
		SKU:         "123",
		Name:        "Quần da Nam",
		Price:       15000,
		Number:      10,
		Description: "Mô tả quần da Nam nè",
		Cate1:       "Thời trang",
		Cate2:       "Thời trang Nam",
		Properties: properties{
			Size: 12,
		},
	},
}

func getProducts(c *gin.Context) {
	filterProducts := c.Request.URL.Query()
	log.Print(filterProducts)

	keys := make([]string, 0, len(filterProducts))
	for k := range filterProducts {
		keys = append(keys, k)
	}

	// filter cate
	productCate1 := products
	cate1 := filterProducts.Get("cate1")
	if (cate1 != "") {
		productCate1 = []product{}
		for _, item := range products {
			if (item.Cate1 == cate1) {
				productCate1 = append(productCate1, item)
			}
		}
	}

	productCate2 := productCate1
	cate2 := filterProducts.Get("cate2")
	if (cate2 != "") {
		productCate2 = []product{}
		for _, item := range products {
			if (item.Cate2 == cate2) {
				productCate2 = append(productCate2, item)
			}
		}
	}

	// Pagination
	productPerPage := 2
	pageNum := 1
	
	page := filterProducts.Get("page")	
	if (page != "") {
		pageNum, _ = strconv.Atoi(page) 
	}
	fmt.Println(pageNum)

	positionStart := (pageNum-1)*productPerPage
	positionEnd := int(math.Min(float64(pageNum*productPerPage), float64(len(productCate2))))

	productsPagination := []product{}
	if (positionStart < positionEnd) {
		productsPagination = productCate2[positionStart:positionEnd]
	}
	c.IndentedJSON(http.StatusOK, productsPagination)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range products {
		if (item.ID == id) {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found product"})
}

func addProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	id := rand.Intn(1000000)
	newProduct.ID = strconv.Itoa(id)

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	var newProduct product

	id := c.Param("id")
	var indexRemove = -1

	for index, item := range products {
		if (item.ID == id) {
			indexRemove = index
			break
		}
	}

	if (indexRemove == -1) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found product"})
		return
	}

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	var indexRemove = -1
	var itemRemove product

	for index, item := range products {
		if (item.ID == id) {
			indexRemove = index
			itemRemove = item
			return
		}
	}

	if (indexRemove == -1) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found product"})
		return
	}

	products = append(products[:indexRemove], products[indexRemove+1:]...)
	c.IndentedJSON(http.StatusCreated, itemRemove)
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/product", addProduct)
	router.PUT("/product/:id", updateProduct)
	router.GET("/product/:id", getProductByID)
	router.DELETE("/product/:id", deleteProduct)

	fmt.Println("Running at http://127.0.0.1:8080")
	router.Run("localhost:8080")
}
