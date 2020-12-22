package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetAllRoom(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "ok")
}
