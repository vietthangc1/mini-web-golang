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
	Router *gin.Engine
	Handler handlers.Handler
	CacheInstance cache.CacheProducts
}

func (a *App) InitializeRoutes() *gin.Engine {
	a.Router.Use(middlewares.CORSMiddleware())

	a.Router.GET("/products", a.HandlerGetProducts)
	a.Router.GET("/product/:id", a.HandlerGetProductByID)
	a.Router.POST("/user", a.HandlerAddUser)
	a.Router.POST("/login", a.HandlerLogin)

	private := a.Router.Group("/")
	private.Use(middlewares.JwtAuthMiddleware())
	private.GET("/user", a.HandlerGetUser)
	private.POST("/product", a.HandlerAddProduct)
	private.PUT("/product/:id", a.HandlerUpdateProduct)
	private.DELETE("/product/:id", a.HandlerDeleteProduct)
	private.DELETE("/user/:id", a.HandlerDeleteUser)

	return nil
}

func (a *App) Run() error {
    return a.Router.Run(os.Getenv("PORT"))
}

func (a *App) Initialize() (error) {
	db, err := models.ConnectDatabaseORM()
	if err != nil {
		return err
	}

	a.Handler.ProductRepo = *modules.NewProductRepository(db)
	a.Handler.UserRepo = *modules.NewUserRepository(db)

	a.Router = gin.Default()
	a.InitializeRoutes()

	a.CacheInstance = cache.CreateCache(os.Getenv("REDISHOST"), 0, 0.5 *1000000000) // db 0, expire 10s

	return nil
}

