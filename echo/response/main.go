package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type User struct {
	Name string `json:"name" query:"name"`
	Age  int    `json:"age" query:"age"`
}

func main() {
	e := echo.New()

	e.Static("/", "static")
	

	logger := e.Logger
	logger.SetLevel(log.DEBUG)
	logger.SetHeader("${time_rfc3339} ${level} ${short_file}:${line} ${message}")

	// JSONを返す
	// curl -X GET "http://localhost:8080/json"
	e.GET("/json", func(c echo.Context) error {
		user := User{
			Name: "Hasushi",
			Age:  21,
		}
		logger.Infof("Hello, %s! You are %d years old.", user.Name, user.Age)
		return c.JSON(200, user)
	})

	// JSON Pretty Print
	// curl -X GET "http://localhost:8080/json/pretty"
	e.GET("/json/pretty", func(c echo.Context) error {
		user := User{
			Name: "Hasushi",
			Age:  21,
		}
		logger.Infof("Hello, %s! You are %d years old.", user.Name, user.Age)
		return c.JSONPretty(200, user, "  ")
	})

	// んーここよくわからんなー
	// 結局attachmentとinlineを切り替えるのってフロントエンド側でやるんじゃねーの？
	// フロント側は、バックからattachmentかinlineかの指示があるからそれに従う感じっぽい
	// curl -X GET "http://localhost:8080/application"
	e.GET("/attach", func(c echo.Context) error {
		logger.Info("Downloading file")
		return c.Attachment("static/main.pdf", "download.pdf")
	})
	// curl -X GET "http://localhost:8080/inline"
	e.GET("/inline", func(c echo.Context) error {
		logger.Info("Downloading file")
		return c.Inline("static/main.pdf", "inline.pdf")
	})


	logger.Fatal(e.Start(":8080"))
}