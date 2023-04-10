package controllers

import (
	"net/http"
	"strconv"
	"todo-list-rest-go/models"

	"github.com/labstack/echo"
)

func FetchAllActivity(c echo.Context) error {
	result, err := models.FetchAllActivity()

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreActivity(c echo.Context) error {
	title := c.FormValue("title")
	email := c.FormValue("email")

	result, err := models.StoreActivity(title, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateActivity(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	email := c.FormValue("email")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateActivity(conv_id, title, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteActivity(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteActivity(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
