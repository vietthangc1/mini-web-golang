package main

import "github.com/gin-gonic/gin"

func GenerateRoutes() *gin.Engine  {
	router := gin.Default()
	router.GET("/products", GetProducts)
	router.POST("/product", AddProduct)
	router.PUT("/product/:id", UpdateProduct)
	router.GET("/product/:id", GetProductByID)
	router.DELETE("/product/:id", DeleteProduct)

	return router
}