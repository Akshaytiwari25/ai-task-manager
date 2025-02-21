package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Akshaytiwari25/ai-task-manager-backend/database"
	"github.com/Akshaytiwari25/ai-task-manager-backend/handlers"
	"github.com/Akshaytiwari25/ai-task-manager-backend/middleware"
	ws "github.com/Akshaytiwari25/ai-task-manager-backend/websocket"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	// Connect to the database.
	database.ConnectDatabase()

	// Initialize Gin router.
	r := gin.Default()

	// Root route.
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the AI Task Manager API!"})
	})

	// Test route.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Public routes.
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.POST("/ai/suggestions", handlers.AISuggestions)

	// Protected routes.
	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.POST("/tasks", handlers.CreateTask)
		protected.GET("/tasks", handlers.GetTasks)
	}

	// WebSocket endpoint.
	r.GET("/ws", ws.HandleWebSocket)

	// Determine port.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
