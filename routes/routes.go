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

	router.POST("/user", handlers.HandlerAddUser)
	router.POST("/login", handlers.HandlerLogin)
	router.DELETE("/user/:id", handlers.HandlerDeleteUser)

	private := router.Group("/")
	private.Use(middlewares.JwtAuthMiddleware())
	private.GET("/user", handlers.HandlerGetUser)

	return router
}