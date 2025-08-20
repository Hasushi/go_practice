package handler

import (
	"net/http"
	"todo/internal/entity"
	"todo/internal/schema"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
	var todo entity.Todo
	if err := c.Bind(&todo); err != nil {
		errRes := schema.ErrorResponse{
			ErrorCode:   "BAD_REQUEST",
			ErrorMessage: "bind error",
		}
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(&todo); err != nil {
		errRes := schema.ErrorResponse{
			ErrorCode:   "BAD_REQUEST",
			ErrorMessage: "title is required",
		}
		return c.JSON(http.StatusBadRequest, errRes)
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
