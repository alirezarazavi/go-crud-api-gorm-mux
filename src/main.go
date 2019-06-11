package main

import (
	"github.com/alirezarazavi/go-crud-api-gorm-mux/app"
	"github.com/alirezarazavi/go-crud-api-gorm-mux/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
