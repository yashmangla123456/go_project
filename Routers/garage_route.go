package Routers

import (
	"gofr.dev/pkg/gofr"
)

func Garage_Route() {
	app := gofr.New()
	app.GET("/garage/all", func(ctx *gofr.Context) (interface{}, error) {
		rows, err := ctx.DB().QueryContext(ctx, "SELECT Car_Num,Owner,Status FROM car")
		if err != nil {
			return nil, err
		}
		return rows, nil
	})
	app.POST("/garage/status:{status}", func(ctx *gofr.Context) (interface{}, error) {
		Status := ctx.PathParam("status")
		Status = string.ToUpper(Status)
		rows, err := ctx.DB().QueryContext(ctx, "SELECT Car_Num,Owner FROM car WHERE Status=?", Status)
		if err != nil {
			return nil, err
		}
		return rows, nil
	})
}
