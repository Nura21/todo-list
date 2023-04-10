package controllers

import (
	"net/http"
	"strconv"
	"todo-list-rest-go/models"

	"github.com/labstack/echo"
)

func FetchAllTodo(c echo.Context) error {
	result, err := models.FetchAllTodo()

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreTodo(c echo.Context) error {
	activity_group_id := c.FormValue("activity_group_id")
	title := c.FormValue("title")
	priority := c.FormValue("priority")

	conv_activity_group_id, err := strconv.Atoi(activity_group_id)

	result, err := models.StoreTodo(conv_activity_group_id, title, priority)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTodo(c echo.Context) error {
	id := c.FormValue("id")
	activity_group_id := c.FormValue("activity_group_id")
	title := c.FormValue("title")
	priority := c.FormValue("priority")

	conv_activity_group_id, err := strconv.Atoi(activity_group_id)
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateTodo(conv_activity_group_id, conv_id, title, priority)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTodo(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteTodo(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
