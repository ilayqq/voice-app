package main

import (
	"log"
	"voice-app/config"
	_ "voice-app/docs"
	auth2 "voice-app/internal/auth"
	"voice-app/internal/product"
	router2 "voice-app/internal/router"
	"voice-app/internal/user"

	"github.com/joho/godotenv"
)

//	@title			Voice-app API
//	@version		1.0
//	@description	Description
//	@host			localhost:8080
//	@BasePath		/api/v1

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

	userRepo := user.NewRepository()
	authService := auth2.NewService(userRepo)
	authHandler := auth2.NewHandler(authService)

	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	productRepo := product.NewRepository()
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	//speechService := speech.NewService()
	//speechHandler := speech.NewHandler(speechService)

	router := router2.NewRouter(authHandler, userHandler, productHandler)

	if err := router.Run(":8080"); err != nil {
		log.Printf("Error starting server: %s", err)
	}
}
