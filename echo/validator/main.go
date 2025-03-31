package main

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// validatorのバリデーションを使うために、構造体にタグを付ける
type User struct {
	Name string `json:"name" query:"name" validate:"required"`
	Age  *int    `json:"age" query:"age" validate:"required"`
}

type CustomValidator struct {
	// validatorは、非ポインタ変数に関しては0値、または空文字を許可しない
	// ポインタ変数に関してはnilを許可しない
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	// カスタムバリデーションを追加
	// カスタムしなかったらageが0の場合もエラーになる
	// if user, ok := i.(*User); ok {
	// 	if user.Age < 0 {
	// 		return echo.NewHTTPError(http.StatusBadRequest, "Age must be positive")
	// 	}
	// }

	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}


func main(){
	e := echo.New()

	logger := e.Logger
	logger.SetLevel(log.DEBUG)
	logger.SetHeader("${time_rfc3339} ${level} ${short_file}:${line} ${message}")

	e.Validator = &CustomValidator{validator: validator.New()}
	
	// JSONパラメータのバリデーション
	// curl -X POST -H "Content-Type: application/json" -d '{"name":"Hasushi","age":21}' "http://localhost:8080/json"
	e.POST("/json", func(c echo.Context) error {
		user := User{}
		if err := c.Bind(&user); err != nil {
			logger.Errorf("Bind error: %v", err)
			logger.Debug(user)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		if err := c.Validate(user); err != nil {
			logger.Errorf("Validation error: %v", err)
			logger.Debug(user)
			return err
		}
		logger.Infof("Hello, %s! You are %d years old.", user.Name, user.Age)
		return c.String(http.StatusOK, "Hello, "+user.Name+"! You are "+strconv.Itoa(*user.Age)+" years old.")
	})

	logger.Fatal(e.Start(":8080"))
}