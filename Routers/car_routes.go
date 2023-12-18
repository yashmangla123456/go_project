package routers

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yashmangla123456/models"
	"github.com/yashmangla123456/utils"
	"gofr.dev/pkg/gofr"
)

func Car_Route(app *gofr.Gofr) {
	app.GET("/car:{person}:{number}", func(ctx *gofr.Context) (interface{}, error) {
		person := ctx.PathParam("person")
		carnumber := ctx.PathParam("number")
		result := utils.ExecQuery(ctx, "SELECT * FROM car WHERE person=? AND number=?", person, carnumber)
		if result != nil {
			var cars []models.Car
			for result.Next() {
				var car models.Car
				err := result.Scan(&car.CarNum, &car.CarName, &car.Color, &car.CarOwner)
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				cars = append(cars, car)
				fmt.Println(car.CarNum)
			}
			return cars, nil
		}
		return result, nil
	})
	app.GET("/car", func(ctx *gofr.Context) (interface{}, error) {
		result := utils.ExecQuery(ctx, "SELECT * FROM car")
		var cars []models.Car
		for result.Next() {
			var car models.Car
			err := result.Scan(&car.CarNum, &car.CarName, &car.Color, &car.CarOwner)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			cars = append(cars, car)
			fmt.Println(car.CarNum)
		}
		fmt.Println("Record Presented")
		return cars, nil

	})

	app.DELETE("/car/{carNum}", func(ctx *gofr.Context) (interface{}, error) {
		carNumber := ctx.PathParam("carNum")

		result := utils.ExecQuery(ctx, "DELETE FROM car WHERE number=?", carNumber)

		fmt.Println("Record Removed")
		return result, nil
	})

}
