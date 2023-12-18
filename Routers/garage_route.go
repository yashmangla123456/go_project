package routers

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yashmangla123456/constants"
	"github.com/yashmangla123456/models"
	"github.com/yashmangla123456/utils"
	"gofr.dev/pkg/gofr"
)

func Garage_Route(app *gofr.Gofr) {
	app.GET("/garage", func(ctx *gofr.Context) (interface{}, error) {
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
		return cars, nil

	})
	app.GET("/garage/:{status}", func(ctx *gofr.Context) (interface{}, error) {
		status := ctx.PathParam("status")
		var result *sql.Rows
		if status == constants.StatusPending || status == constants.StatusDone || status == constants.StatusWip {
			result = utils.ExecQuery(ctx, "SELECT car,status FROM job WHERE status=?", status)
		} else {
			result = utils.ExecQuery(ctx, "SELECT car,status FROM job ")
		}
		var jobs []models.JobSummary
		for result.Next() {

			var job models.JobSummary
			err := result.Scan(&job.Car, &job.Status)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			jobs = append(jobs, job)
		}
		return jobs, nil

		return result, nil
	})
}
