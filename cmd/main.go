package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yagoayala/server/internal/config"
	"github.com/yagoayala/server/internal/models"
	"github.com/yagoayala/server/internal/routes"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	routes.SetupRoutes(router, db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
