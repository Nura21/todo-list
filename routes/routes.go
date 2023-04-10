package routes

import (
	"net/http"
	"todo-list-rest-go/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo")
	})

	e.GET("/activities", controllers.FetchAllActivity)
	e.POST("/activities/store", controllers.StoreActivity)
	e.PUT("/activities/update", controllers.UpdateActivity)
	e.DELETE("/activities/delete", controllers.DeleteActivity)

	e.GET("/todos", controllers.FetchAllTodo)
	e.POST("/todos/store", controllers.StoreTodo)
	e.PUT("/todos/update", controllers.UpdateTodo) 
	e.DELETE("/todos/delete", controllers.DeleteTodo)

	return e
}
