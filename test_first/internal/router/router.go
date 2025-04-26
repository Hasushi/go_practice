package router

import (
	"todo/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todos", handler.GetTodos)

	e.POST("/todos", handler.CreateTodo)

	e.GET("/todos/:id", handler.GetTodo)

	e.PUT("/todos/:id", handler.UpdateTodo)

	e.DELETE("/todos/:id", handler.DeleteTodo)

	return e
}
