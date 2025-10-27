package main

import (
	"log"
	"voice-app/backend/config"
	speech2 "voice-app/backend/internal/speech"
	"voice-app/backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

	router := gin.Default()

	speechService := speech2.NewService()
	speechHandler := speech2.NewHandler(speechService)

	api := router.Group("/api/v1")
	api.Use(middleware.JWTAuth())

	router.GET("/test", func(c *gin.Context) { c.JSON(200, "test") })
	router.POST("/api/speech/recognize", speechHandler.Recognize)

	router.Run(":8080")
	if err != nil {
		log.Printf("Error starting server: %s", err)
	}
}
