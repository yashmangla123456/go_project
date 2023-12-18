package routers

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yashmangla123456/models"
	"github.com/yashmangla123456/utils"
	"gofr.dev/pkg/gofr"
)

func Person_Route(app *gofr.Gofr) {
	app.GET("/person", func(ctx *gofr.Context) (interface{}, error) {
		result := utils.ExecQuery(ctx, "SELECT * FROM person")

		fmt.Println(result.Err())
		var persons []models.Person
		for result.Next() {
			var person models.Person
			err := result.Scan(&person.Name, &person.Contact, &person.Address)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			persons = append(persons, person)
			fmt.Println(person.Name)
		}
		return persons, nil

	})

	app.GET("/person/{phone}:{data}", func(ctx *gofr.Context) (interface{}, error) {
		phone := ctx.PathParam("phone")
		data := ctx.PathParam("data")
		var result *sql.Rows
		if data == "true" {
			result := utils.ExecQuery(ctx, "SELECT job.*, car.*, person.* FROM person INNER JOIN car ON person.contact = car.person INNER JOIN job ON job.car = car.number WHERE person.contact=?", phone)
			var person models.PersonData
		LOOP:
			for result.Next() {
				var car models.CarData
				var job models.JobSummaryData
				err := result.Scan(&job.Id, &car.CarNum, &job.Fault, &job.Status, nil, &car.CarName, &car.Color, nil, nil, nil, nil)
				if err != nil {
					fmt.Println(err.Error())
				}
				for _, c := range person.Cars {
					if c.CarNum == car.CarNum {
						c.Jobs = append(c.Jobs, job)
						goto LOOP
					}
				}
				car.Jobs = append(car.Jobs, job)
				person.Cars = append(person.Cars, car)
			}
			return person, nil
		} else {
			result = utils.ExecQuery(ctx, "SELECT * FROM person WHERE contact=?", phone)
			var persons []models.Person
			for result.Next() {
				fmt.Println("======kjbkbj====")
				var person models.Person
				err := result.Scan(&person.Name, &person.Contact, &person.Address)
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				persons = append(persons, person)
				fmt.Println(person.Name)
			}
			return persons, nil
		}
		return result, nil
	})

	app.DELETE("/person/{phone}", func(ctx *gofr.Context) (interface{}, error) {
		phone := ctx.PathParam("phone")
		result := utils.ExecQuery(ctx, "SELECT * FROM person WHERE contact=?", phone)
		if result != nil {
			result = utils.ExecQuery(ctx, "DELETE FROM Owner WHERE contact=?", phone)
			return "Deleted", nil
		}
		return "No Reocrd Found", nil
	})
}
