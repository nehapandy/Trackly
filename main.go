package main

import (
	routes "TRACKLY/Router"
	"TRACKLY/database"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	database.DBInstance()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Change this to the port Postman is using
	}

	// Initialize Gin router
	router := gin.New()
	router.Use(gin.Logger()) // Use logging middleware

	// Register Routes
	routes.AuthRouter(router)
	routes.UserRouter(router)
	// Run the server
	router.Run(":" + port)
	expiryTime := time.Unix(1734503030488/1000, 0) // divide by 1000 to convert to seconds
	fmt.Println("Token expiry:", expiryTime)
}
