package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasan-indra/learn-golang-api/internal/handlers"
	"github.com/hasan-indra/learn-golang-api/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// API v1 group
	// v1 := router.Group("/api/v1")
	// {
	// 	v1.GET("/users", handlers.GetUsers)
	// 	v1.POST("/users", handlers.CreateUser)
	// }

	return router
}
