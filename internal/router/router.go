package router

import (
	"voice-app/internal/auth"
	"voice-app/internal/product"
	"voice-app/internal/user"
	"voice-app/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	authHandler *auth.Handler,
	userHandler *user.Handler,
	productHandler *product.Handler,
) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	api := r.Group("/api/v1")
	api.Use(middleware.JWTAuth())
	{
		user := api.Group("/users")
		{
			user.GET("", userHandler.GetUsers)
		}
		product := api.Group("/products")
		{
			product.GET("", productHandler.GetAll)
			product.POST("", productHandler.AddProduct)
		}
	}

	return r
}
