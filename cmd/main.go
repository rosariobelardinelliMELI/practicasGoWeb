package main

import (
	"fmt"
	"goweb/internal/application"
	"os"
)

func main() {
	// env
	token := os.Getenv("TOKEN")

	// app
	// - config
	app := application.NewDefaultHTTP(":8080", token)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
