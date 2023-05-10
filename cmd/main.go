package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vinilunni/pogr_api_golang_vinil/internal/handlers"
)

func main() {
	// Load env vars
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	// v0.1
	v0_1 := router.Group("/api/v0.1")

	// user
	user := v0_1.Group("/user")

	user.GET("/:id", handlers.GetUserByID)

	if err := router.Run(os.Getenv("HTTP_PORT")); err != nil {
		log.Fatalf("Fail to run router: %v", err)
	}

}
