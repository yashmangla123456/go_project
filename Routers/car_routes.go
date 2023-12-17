package Routers

import (
	"fmt"
	"gofr.dev/pkg/gofr"
)

func Car_Route() {
	app := gofr.New()
	app.GET("/car", func(ctx *gofr.Context) (interface{}, error) {
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Car")
		if err != nil {
			return nil, err
		}
		fmt.Println("Record Presented")
		return rows, nil
	})

	app.POST("/car/delete/{Car_Num}", func(ctx *gofr.Context) (interface{}, error) {
		Car_Num := ctx.PathParam("Car_Num")
		Car_Num = string.ToUpper(Car_Num)
		rows, err := ctx.DB().QueryContext(ctx, "DELETE FROM Car WHERE carNumber=?", Car_Num)
		if err != nil {
			return nil, err
		}
		fmt.Println("Record Removed")
		return rows, nil
	})

}
