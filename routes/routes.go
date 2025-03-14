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
	api.POST("/users/create", controllers.SoapCreateUser)
	api.GET("/users/:id", controllers.SoapGetUser)
	api.GET("/users", controllers.SoapGetUsers)
	api.PUT("/users/:id", controllers.SoapUpdateUser)
	api.DELETE("/users/:id", controllers.SoapDeleteUser)

	// Product routes
	api.GET("/products", controllers.GetProduct)
	api.POST("/products", controllers.CreateProduct)
	api.PUT("/products/:id", controllers.UpdateProduct)
	api.DELETE("/products/:id", controllers.DeleteProduct)

	// Category routes
	api.POST("/categories/create", controllers.SoapCreateCategory)
	api.GET("/categories/:id", controllers.SoapGetCategory)
	api.GET("/categories", controllers.SoapGetCategories)
	api.PUT("/categories/:id", controllers.SoapUpdateCategory)
	api.DELETE("/categories/:id", controllers.SoapDeleteCategory)

}
