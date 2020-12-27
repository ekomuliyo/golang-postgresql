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

func DeleteRoom(c echo.Context) error {

	id := c.FormValue("id")
	convID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteRoom(convID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateRoom(c echo.Context) error {

	id := c.FormValue("id")
	price := c.FormValue("price")
	maxCapacity := c.FormValue("max_capacity")
	convID, err := strconv.Atoi(id)
	convPrice, err := strconv.Atoi(price)
	convMaxCapacity, err := strconv.Atoi(maxCapacity)

	nameRoom := c.FormValue("name_room")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateRoom(convID, nameRoom, convPrice, convMaxCapacity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
