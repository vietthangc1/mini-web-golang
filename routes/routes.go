package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/handlers"
	"github.com/vietthangc1/mini-web-golang/middlewares"
	"github.com/vietthangc1/mini-web-golang/models"
)

func GenerateRoutes() *gin.Engine  {
	db, err := models.ConnectDatabase()
	if (err != nil) {
		log.Fatal(err)
	}
	h := handlers.NewBaseHandler(db)
	
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.GET("/products", h.HandlerGetProducts)
	router.POST("/product", h.HandlerAddProduct)
	router.PUT("/product/:id", h.HandlerUpdateProduct)
	router.GET("/product/:id", h.HandlerGetProductByID)
	router.DELETE("/product/:id", h.HandlerDeleteProduct)

	router.POST("/user", h.HandlerAddUser)
	router.POST("/login", h.HandlerLogin)
	router.DELETE("/user/:id", h.HandlerDeleteUser)

	private := router.Group("/")
	private.Use(middlewares.JwtAuthMiddleware())
	private.GET("/user", h.HandlerGetUser)

	return router
}