package controllers

import (
	"golang-postgresql/helpers"
	"golang-postgresql/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func RegisterUser(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	IDGroup, err := strconv.Atoi(c.FormValue("id_group"))

	hashPassword, _ := helpers.HashPassowrd(password)

	result, err := models.RegisterUser(username, hashPassword, email, IDGroup)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, IDGroup, err := models.LoginUser(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// create token jwt
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["id_group"] = IDGroup
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// generate encoded token
	tokenResult, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenResult,
	})
}

func GetAllUser(c echo.Context) error {
	result, err := models.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	username := c.FormValue("username")
	email := c.FormValue("email")

	convID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(convID, username, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {

	id := c.FormValue("id")

	convID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(convID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
