package main

import (
	"github.com/yashmangla123456/Routers"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// register route greet
	Routers.Car_Route()
	Routers.Garage_Route()
	Routers.Job_Summary()
	Routers.Owner_Route()
	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello", nil
	})

	app.Start()
}
