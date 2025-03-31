package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main(){
	e := echo.New()
	logger := e.Logger
	// ログレベルを設定
	// 設定したレベル以上のログが出力される
	logger.SetLevel(log.DEBUG)
	// ログの出力先を設定
	logger.SetOutput(os.Stdout)
	// ログのフォーマットを設定
	logger.SetHeader("${time_rfc3339} ${level} ${short_file}:${line} ${message}")

	e.GET("/", func(c echo.Context) error {
		// ログ出力
		logger.Debugf("Debug log")
		logger.Infof("Hello, World!")
		logger.Warnf("Warning log")
		logger.Errorf("Error log")
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}