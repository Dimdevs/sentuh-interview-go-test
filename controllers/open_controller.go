package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func OpenAPI(c echo.Context) error {
	response := Response{
		Message: "OPEN API by Dimdevs",
	}
	return c.JSON(http.StatusOK, response)
}
