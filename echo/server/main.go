package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

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

	// 普通のサーバ起動
	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal(err)
	}

	// HTTP/2.0のサーバ起動
	// h2s := &http2.Server{
	// 	MaxConcurrentStreams: 100,
	// 	MaxReadFrameSize:   1024 * 1024,
	// 	IdleTimeout: 10 * time.Second,
	// }
	// s := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: h2c.NewHandler(e, h2s),
	// }
	// if err := s.ListenAndServe(); err != http.ErrServerClosed {
	// 	logger.Fatal(err)
	// }

	// TLSのサーバ起動
	// s := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: e,
	// 	TLSConfig: &tls.Config{},
	// }
	// if err := s.ListenAndServeTLS("server.crt", "server.key"); err != http.ErrServerClosed {
	// 	logger.Fatal(err)
	// }

}