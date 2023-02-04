package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/handlers"
	"github.com/vietthangc1/mini-web-golang/middlewares"
)

func GenerateRoutes() *gin.Engine  {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.GET("/products", handlers.HandlerGetProducts)
	router.POST("/product", handlers.HandlerAddProduct)
	router.PUT("/product/:id", handlers.HandlerUpdateProduct)
	router.GET("/product/:id", handlers.HandlerGetProductByID)
	router.DELETE("/product/:id", handlers.HandlerDeleteProduct)

	return router
}