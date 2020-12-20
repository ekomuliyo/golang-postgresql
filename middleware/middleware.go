package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// IsAuthenticated ...
var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

// AdminMiddleware ...
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hToken := c.Request().Header.Get("Authorization") // Bearer token12647
		token := strings.Split(hToken, " ")[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Unable to parse token",
			})
		}

		if int(claims["id_group"].(float64)) == 1 {
			return next(c)
		}

		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "Not authorized",
		})
	}
}
