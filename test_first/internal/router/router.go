package router

import (
	"todo/internal/handler"

	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.GET("/todos", handler.GetTodos)

	e.POST("/todos", handler.CreateTodo)

	e.GET("/todos/:id", handler.GetTodo)

	e.PUT("/todos/:id", handler.UpdateTodo)

	e.DELETE("/todos/:id", handler.DeleteTodo)

	return e
}
