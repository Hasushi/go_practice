package main

import "todo/internal/router"

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}