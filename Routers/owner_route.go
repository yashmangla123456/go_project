package Routers

import (
	"gofr.dev/pkg/gofr"
)

func Owner_Route() {
	app := gofr.New()
	app.GET("/Owner/all", func(ctx *gofr.Context) (interface{}, error) {
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Owner")
		if err != nil {
			return nil, err
		}
		return rows, nil
	})
	app.POST("/Owner/delete/{Phone}", func(ctx *gofr.Context) (interface{}, error) {
		Phone := ctx.PathParam("Phone")
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Owner WHERE Phone=?", Phone)
		if rows != nil {
			_, err2 := ctx.DB().QueryContext(ctx, "DELETE FROM Owner WHERE Phone=?", Phone)
			if err2 != nil {
				return nil, err2
			}
		}
		if err != nil {
			return nil, err
		}
		return rows, nil
	})
}
