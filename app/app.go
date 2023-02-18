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
	CacheInstance cache.CacheProducts
}

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

func NewCacheInstance() cache.CacheProducts {
	return cache.NewCache(os.Getenv("REDISHOST"), 0, 10*1000000000) // db 0, expire 10s
}



func (a *App) Initialize() error {
	db, err := models.ConnectDatabaseORM()
	if err != nil {
		return err
	}

	a.Handler = handlers.NewHandler(modules.NewProductRepository(db), modules.NewUserRepository(db))

	a.Router = a.InitializeRoutes()

	a.CacheInstance = NewCacheInstance()

	return nil
}

// wire.go

func NewApp(router *gin.Engine, handler handlers.Handler, cacheInstance cache.CacheProducts) App {
	return App{
		Router:        router,
		Handler:       handler,
		CacheInstance: cacheInstance,
	}
}
