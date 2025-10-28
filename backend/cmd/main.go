package main

import (
	"log"
	"voice-app/config"
	auth2 "voice-app/internal/auth"
	"voice-app/internal/speech"
	"voice-app/internal/user"
	"voice-app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("backend/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

	router := gin.Default()

	userRepo := user.NewRepository()
	authService := auth2.NewService(userRepo)
	authHandler := auth2.NewHandler(authService)

	speechService := speech.NewService()
	speechHandler := speech.NewHandler(speechService)

	auth := router.Group("/auth")
	auth.POST("/register", authHandler.Register)

	api := router.Group("/api/v1")
	api.Use(middleware.JWTAuth())

	router.POST("/api/speech/recognize", speechHandler.Recognize)

	router.Run(":8080")
	if err != nil {
		log.Printf("Error starting server: %s", err)
	}
}
