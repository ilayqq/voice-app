package main

import (
	"log"
	"voice-app/config"
	_ "voice-app/docs"
	auth2 "voice-app/internal/auth"
	"voice-app/internal/speech"
	"voice-app/internal/user"
	"voice-app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Voice-app API
// @version 1.0
// @description Description
// @host localhost:8080
// @BasePath /api/v1
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

	api := router.Group("/api/v1")
	//api.Use(middleware.JWTAuth())

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
	api.GET("/test", middleware.JWTAuth(), func(c *gin.Context) { c.JSON(200, gin.H{"test": "test"}) })

	router.POST("/api/speech/recognize", speechHandler.Recognize)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
	if err != nil {
		log.Printf("Error starting server: %s", err)
	}
}
