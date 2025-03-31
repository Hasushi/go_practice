package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// jsonやqueryのパラメータを複数設定できる
type User struct {
	Name string `json:"name" query:"name"`
	Age  int    `json:"age" query:"age"`
}

func main(){
	e := echo.New()
	logger := e.Logger
	logger.SetLevel(log.INFO)
	logger.SetHeader("${time_rfc3339} ${level} ${short_file}:${line} ${message}")

	// 普通のルート
	e.GET("/", func(c echo.Context) error {
		logger.Infof("Hello, World!")
		return c.String(200, "Hello, World!")
	})

	// クエリパラメータ
	// curl -X GET "http://localhost:8080/query?name=Hasushi&age=21"
	e.GET("/query", func(c echo.Context) error {
		user := User{}
		if err := c.Bind(&user); err != nil {
			return c.String(400, "Bad Request")
		}
		logger.Infof("Hello, %s! You are %d years old.", user.Name, user.Age)
		return c.String(200, "Hello, "+user.Name+"! You are "+string(rune(user.Age))+" years old.")
	})
	// JSONパラメータ
	// curl -X POST -H "Content-Type: application/json" -d '{"name":"Hasushi","age":21}' "http://localhost:8080/json"
	e.POST("/json", func(c echo.Context) error {
		user := User{}
		if err := c.Bind(&user); err != nil {
			return c.String(400, "Bad Request")
		}
		logger.Infof("Hello, %s! You are %d years old.", user.Name, user.Age)
		return c.String(200, "Hello, "+user.Name+"! You are "+string(rune(user.Age))+" years old.")
	})

	logger.Fatal(e.Start(":8080"))
	
}