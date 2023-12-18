package routers

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yashmangla123456/constants"
	"github.com/yashmangla123456/models"
	"github.com/yashmangla123456/utils"
	"gofr.dev/pkg/gofr"
	"io"
)

func Job_Summary(app *gofr.Gofr) {

	app.GET("/job_schedule/all", func(ctx *gofr.Context) (interface{}, error) {
		result := utils.ExecQuery(ctx, "SELECT * FROM job")
		fmt.Println(result.Err())
		var jobs []models.JobSummary
		for result.Next() {
			var job models.JobSummary
			err := result.Scan(&job.Id, &job.Car, &job.Fault, &job.Status)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			jobs = append(jobs, job)
			fmt.Println(job.Fault)
		}
		return jobs, nil
	})
	app.GET("/job_schedule:{status}", func(ctx *gofr.Context) (interface{}, error) {
		fmt.Println("Record Presented")
		status := ctx.PathParam("status")
		var result *sql.Rows
		if status == constants.StatusPending || status == constants.StatusWip || status == constants.StatusDone {
			result = utils.ExecQuery(ctx, "SELECT * FROM job WHERE status=?", status)
		} else {
			result = utils.ExecQuery(ctx, "SELECT * FROM job;")
		}

		fmt.Println("Record Presented")
		fmt.Println(result.Err())
		var jobs []models.JobSummary
		for result.Next() {
			var job models.JobSummary
			err := result.Scan(&job.Id, &job.Car, &job.Fault, &job.Status)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			jobs = append(jobs, job)
			fmt.Println(job.Fault)
		}
		return jobs, nil
	})

	app.GET("/job_schedule/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		result := utils.ExecQueryRow(ctx, "SELECT * FROM job WHERE id=?", id)
		if result == nil {
			fmt.Println("No Results")
		}
		return result, nil
	})

	app.POST("/job_schedule/", func(ctx *gofr.Context) (interface{}, error) {
		var person models.Person
		var job models.JobSummary
		var car models.Car
		reqData, err := io.ReadAll(ctx.Request().Body)
		err = json.Unmarshal(reqData, &person)
		err = json.Unmarshal(reqData, &job)
		err = json.Unmarshal(reqData, &car)

		result := utils.ExecQuery(ctx, "SELECT * FROM person WHERE contact=?", person.Contact)
		if result == nil {
			result = utils.ExecQuery(ctx, "INSERT INTO person(name,contact,address) VALUES(?,?,?) ", person.Name, person.Contact, person.Address)
		}

		result = utils.ExecQuery(ctx, "SELECT * FROM car WHERE number=?", car.CarNum)
		if result == nil {
			result = utils.ExecQuery(ctx, "INSERT INTO car(number,name,color,person) VALUES(?,?,?,?) ", car.CarNum, car.CarName, car.Color, person.Contact)
		}

		result = utils.ExecQuery(ctx, "INSERT INTO job(car,fault,status) VALUES(?,?,?) ", car.CarNum, job.Fault, constants.StatusPending)
		fmt.Println(result, err)
		return "hello", nil
	})

	app.DELETE("/job_schedule/{id}", func(ctx *gofr.Context) (interface{}, error) {
		jobId := ctx.PathParam("id")
		result := utils.ExecQuery(ctx, "DELETE FROM job WHERE id=?", jobId)
		fmt.Println(result)
		fmt.Println("Record Removed")
		return "Removed", nil
	})

	app.PUT("/job_schedule/{id}:{status}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		status := ctx.PathParam("status")
		rows := utils.ExecQuery(ctx, "UPDATE job SET status=? WHERE id=?", status, id)

		fmt.Println("Record Removed")
		fmt.Println(rows)
		return "Done", nil
	})
}
