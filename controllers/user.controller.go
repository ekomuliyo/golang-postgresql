package controllers

import (
	"golang-postgresql/helpers"
	"golang-postgresql/models"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func RegisterUser(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	IDGroup, err := strconv.Atoi(c.FormValue("id_group"))

	photo, err := c.FormFile("photo")
	if err != nil {
		return err
	}

	srcPhoto, err := photo.Open()
	if err != nil {
		return err
	}
	defer srcPhoto.Close()

	var randNumber string = strconv.Itoa(rand.Intn(10000))
	var nameFile string = "images/photo-" + randNumber + "." + strings.Split(photo.Filename, ".")[1]
	dstPhoto, err := os.Create(nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	defer dstPhoto.Close()

	// copy file
	if _, err = io.Copy(dstPhoto, srcPhoto); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	hashPassword, _ := helpers.HashPassowrd(password)

	result, err := models.RegisterUser(username, hashPassword, email, nameFile, IDGroup)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// return c.JSON(http.StatusOK, "ok")
	return c.JSON(http.StatusInternalServerError, result)
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
