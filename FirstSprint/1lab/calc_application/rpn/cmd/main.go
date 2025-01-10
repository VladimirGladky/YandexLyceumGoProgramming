package main

import (
	"github.com/VladimirGladky/rpn/rpn/internal/application"
)

func main() {
	app := application.New()
	err := app.Run()
	if err != nil {
		return
	}
}
