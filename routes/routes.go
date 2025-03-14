package routes

import (
	"go-test/controllers"
	"go-test/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", controllers.OpenAPI)

	// Auth routes (public)
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// API group with JWT middleware
	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware())

	// User routes
	api.GET("/users", controllers.GetUser)
	api.POST("/users", controllers.CreateUser)
	api.PUT("/users/:id", controllers.UpdateUser)
	api.DELETE("/users/:id", controllers.DeleteUser)

	// Product routes
	api.GET("/products", controllers.GetProduct)
	api.POST("/products", controllers.CreateProduct)
	api.PUT("/products/:id", controllers.UpdateProduct)
	api.DELETE("/products/:id", controllers.DeleteProduct)

	// Category routes
	api.GET("/categories", controllers.GetCategories)
	api.POST("/categories", controllers.CreateCategory)
	api.GET("/categories/:id", controllers.GetCategory)
	api.PUT("/categories/:id", controllers.UpdateCategory)
	api.DELETE("/categories/:id", controllers.DeleteCategory)
}
