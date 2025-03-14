package routes

import (
	"go-test/controllers"
	"go-test/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", controllers.OpenAPI)

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware())

	api.POST("/users/create", controllers.SoapCreateUser)
	api.GET("/users/:id", controllers.SoapGetUser)
	api.GET("/users", controllers.SoapGetUsers)
	api.PUT("/users/:id", controllers.SoapUpdateUser)
	api.DELETE("/users/:id", controllers.SoapDeleteUser)

	api.POST("/products/create", controllers.SoapCreateProduct)
	api.GET("/products/:id", controllers.SoapGetProduct)
	api.GET("/products", controllers.SoapGetProducts)
	api.PUT("/products/:id", controllers.SoapUpdateProduct)
	api.DELETE("/products/:id", controllers.SoapDeleteProduct)

	api.POST("/categories/create", controllers.SoapCreateCategory)
	api.GET("/categories/:id", controllers.SoapGetCategory)
	api.GET("/categories", controllers.SoapGetCategories)
	api.PUT("/categories/:id", controllers.SoapUpdateCategory)
	api.DELETE("/categories/:id", controllers.SoapDeleteCategory)

	api.GET("/products/soap/products.wsdl", func(c echo.Context) error {
		return c.File("product.wsdl")
	})
}
