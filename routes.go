package main

import "github.com/gin-gonic/gin"

func GenerateRoutes() *gin.Engine  {
	router := gin.Default()
	router.GET("/products", HandlerGetProducts)
	router.POST("/product", HandlerAddProduct)
	router.PUT("/product/:id", HandlerUpdateProduct)
	router.GET("/product/:id", HandlerGetProductByID)
	router.DELETE("/product/:id", HandlerDeleteProduct)

	return router
}