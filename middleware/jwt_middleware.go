package middleware

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SuccessHandler: func(c echo.Context) {
			log.Println("[AUTH SUCCESS]", c.Path())
		},
		ErrorHandler: func(c echo.Context, err error) error {
			log.Println("[AUTH ERROR]", err)
			return echo.NewHTTPError(401, "Unauthorized: "+err.Error())
		},
	})
}
