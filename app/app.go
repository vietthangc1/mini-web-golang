package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/cache"
	"github.com/vietthangc1/mini-web-golang/handlers"
	"github.com/vietthangc1/mini-web-golang/middlewares"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/modules"
)

type App struct {
	Router        *gin.Engine
	Handler       handlers.Handler
}

// for dependency injection not using google wire

func (a *App) InitializeRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	router.GET("/products", a.HandlerGetProducts)
	router.GET("/product/:id", a.HandlerGetProductByID)
	router.POST("/user", a.HandlerAddUser)
	router.POST("/login", a.HandlerLogin)

	private := router.Group("/")
	private.Use(middlewares.JwtAuthMiddleware())
	private.GET("/user", a.HandlerGetUser)
	private.POST("/product", a.HandlerAddProduct)
	private.PUT("/product/:id", a.HandlerUpdateProduct)
	private.DELETE("/product/:id", a.HandlerDeleteProduct)
	private.DELETE("/user/:id", a.HandlerDeleteUser)

	return router
}

func (a *App) Run() error {
	return a.Router.Run(os.Getenv("PORT"))
}


func (a *App) Initialize() error {
	db, err := models.ConnectDatabaseORM()
	if err != nil {
		return err
	}

	a.Handler = handlers.NewHandler(modules.NewProductRepository(db), modules.NewUserRepository(db), cache.NewCacheInstance())

	a.Router = a.InitializeRoutes()

	return nil
}

// for wire

func NewApp(router *gin.Engine, handler handlers.Handler) App {
	return App{
		Router:        router,
		Handler:       handler,
	}
}

func NewRouter(h handlers.Handler) *gin.Engine {
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
