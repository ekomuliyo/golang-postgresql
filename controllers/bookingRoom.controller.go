package controllers

import (
	"golang-postgresql/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllBookingRoom(c echo.Context) error {

	result, err := models.GetAllBookingRoom()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
