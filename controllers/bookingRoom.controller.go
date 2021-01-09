package controllers

import (
	"golang-postgresql/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllBookingRoom(c echo.Context) error {

	result, err := models.GetAllBookingRoom()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreBookingRoom(c echo.Context) error {

	dateBooking := c.FormValue("date_booking")
	IDRoom, err := strconv.Atoi(c.FormValue("id_room"))
	IDUser, err := strconv.Atoi(c.FormValue("id_user"))
	amountCapacity, err := strconv.Atoi(c.FormValue("amount_capacity"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.StoreBookingRoom(IDRoom, IDUser, dateBooking, amountCapacity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
