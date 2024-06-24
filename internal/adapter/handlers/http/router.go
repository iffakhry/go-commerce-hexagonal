package http

import (
	"github.com/iffakhry/go-commerce-hexagonal/internal/adapter/storages/postgres/repositories"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/pkg/encrypt"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] status=${status} method=${method} uri=${uri} latency=${latency_human} \n",
	}))

	//factory
	hashService := encrypt.NewHashService()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, hashService)
	userHandler := NewUserHandler(userService)
	api := e.Group("/api")
	apiV1 := api.Group("/v1")

	// apiV1.POST("/login", userHandler.Login)

	// apiV1.GET("/profile", userController.GetProfile, middlewares.JWTMiddleware())
	// apiV1.GET("/users", userController.GetAllUser, middlewares.JWTMiddleware())
	apiV1.POST("/users", userHandler.CreateUser)
	// apiV1.PUT("/users", userController.Update, middlewares.JWTMiddleware())
	// apiV1.DELETE("/users", userController.Delete, middlewares.JWTMiddleware())

	// productV1 := apiV1.Group("/products")

	// productRepo := repositories.NewProductRepository(db)
	// productController := controllers.NewProductController(productRepo)
	// productV1.POST("", productController.Create, middlewares.JWTMiddleware())
	// productV1.GET("", productController.GetAll)
	// productV1.GET("/:id", productController.GetById)
	// productV1.PUT("/:id", productController.Update, middlewares.JWTMiddleware())
	// productV1.DELETE("/:id", productController.Delete, middlewares.JWTMiddleware())
	// return e
}
