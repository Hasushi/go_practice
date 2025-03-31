package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	// 普通のルート
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// パスパラメータ
	// curl -X GET "http://localhost:8080/Hasushi"
	e.GET("/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(200, "Hello, "+name+"!")
	})

	// クエリパラメータ
	// curl -X GET "http://localhost:8080/query?name=Hasushi&age=21"
	e.GET("/query", func(c echo.Context) error {
		name := c.QueryParam("name")
		age := c.QueryParam("age")
		return c.String(200, "Hello, "+name+"! You are "+age+" years old.")
	})
	
	e.Logger.Fatal(e.Start(":8080"))
}