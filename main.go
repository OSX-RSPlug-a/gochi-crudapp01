package main

import (
	"context"
	"fmt"

	"github.com/OSX-RSPlug-a/gochi-crudapp01/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed at", err)
	}
}