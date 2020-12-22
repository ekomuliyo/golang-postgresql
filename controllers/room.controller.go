package controllers

import (
	"golang-postgresql/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func StoreRoom(c echo.Context) error {
	nameRoom := c.FormValue("name_room")
	price, err := strconv.Atoi(c.FormValue("price"))
	maxCapacity, err := strconv.Atoi(c.FormValue("max_capacity"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	result, err := models.StoreRoom(nameRoom, price, maxCapacity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// return c.JSON(http.StatusOK, "ok")
	return c.JSON(http.StatusInternalServerError, result)
}

func GetAllRoom(c echo.Context) error {
	result, err := models.GetAllRoom()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
