package main

import (
	"log"

	"github.com/alberto-valdez/lazy-taxes/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userHandler := handlers.NewUserHandler()

	api := router.Group("/api/v1")
	{
		userHandler.RegisterRoutes(api)
	}

	log.Println("Servidor en http://localhost:8080")
	router.Run(":8080")
}
