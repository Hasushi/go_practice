package handler

import (
	"net/http"
	"todo/internal/entity"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
	var todo entity.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bind error"})
	}
	if err := c.Validate(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "validation error"})
	}

	

	panic("not implemented")
}

func GetTodo(c echo.Context) error {
	panic("not implemented")
}

func GetTodos(c echo.Context) error {
	panic("not implemented")
}

func UpdateTodo(c echo.Context) error {
	panic("not implemented")
}

func DeleteTodo(c echo.Context) error {
	panic("not implemented")
}
