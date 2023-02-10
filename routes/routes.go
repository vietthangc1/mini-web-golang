package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/handlers"
	"github.com/vietthangc1/mini-web-golang/middlewares"
	"github.com/vietthangc1/mini-web-golang/models"
)

func GenerateRoutes() *gin.Engine {
	db, err := models.ConnectDatabaseORM()
	if err != nil {
		log.Fatal(err)
	}
	h := handlers.NewBaseHandler(db)

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.GET("/products", h.HandlerGetProducts)
	router.GET("/product/:id", h.HandlerGetProductByID)
	router.POST("/user", h.HandlerAddUser)
	router.POST("/login", h.HandlerLogin)

	private := router.Group("/")
	private.Use(middlewares.JwtAuthMiddleware())
	private.GET("/user", h.HandlerGetUser)
	private.POST("/product", h.HandlerAddProduct)
	private.PUT("/product/:id", h.HandlerUpdateProduct)
	private.DELETE("/product/:id", h.HandlerDeleteProduct)
	private.DELETE("/user/:id", h.HandlerDeleteUser)

	return router
}
