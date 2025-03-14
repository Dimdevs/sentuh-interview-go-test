package main

import (
	"go-test/config"
	"go-test/routes"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		e.Logger.Fatal("Gagal membuka file log:", err)
	}
	defer logFile.Close()

	e.Logger.SetOutput(logFile)

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	config.InitDB()
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":5050"))
}
