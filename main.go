package main

import (
	"github.com/yashmangla123456/Routers"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	//Database.Connection()

	// register route greet
	routers.Car_Route(app)
	routers.Garage_Route(app)
	routers.Job_Summary(app)
	routers.Person_Route(app)
	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello", nil
	})

	app.Start()
}
