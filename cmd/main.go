package main

import (
	"fmt"
	"goweb/internal/application"
)

func main() {
	// env
	// ...

	// app
	// - config
	app := application.NewDefaultHTTP(":8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
