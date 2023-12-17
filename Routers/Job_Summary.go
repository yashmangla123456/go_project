package Routers

import (
	"encoding/json"
	"fmt"
	"github.com/yashmangla123456/models"
	"gofr.dev/pkg/gofr"
)

func Job_Summary() {
	app := gofr.New()
	app.GET("/job_schedule/all", func(ctx *gofr.Context) (interface{}, error) {
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Job_schedule")
		if err != nil {
			return nil, err
		}
		fmt.Println("Record Presented")
		return rows, nil
	})
	app.POST("/job_schedule/", func(ctx *gofr.Context) (interface{}, error) {
		var job_dta map[string]interface{}
		err := json.Unmarshal([]byte(models.IncommingJobData), &job_dta)
		Car_Num := job_dta["Car_Num"]
		Car_Name := job_dta["Car_Name"]
		Owner := job_dta["Owner"]
		Color := job_dta["Color"]
		Contact := job_dta["Contact"]
		Address := job_dta["Address"]
		Fault := job_dta["Fault"]
		_, err1 := ctx.DB().QueryContext(ctx, "INSERT INTO Job_schedule(carNumber,carName,Fault,Status) VALUES(?,?,?,'NEW_ARRIVED') ", Car_Num, Car_Name, Fault)
		_, err2 := ctx.DB().QueryContext(ctx, "INSERT INTO Car(carNumber,carName,Fault,Color) VALUES(?,?,?,?) ", Car_Num, Car_Name, Fault, Color)
		q1, err3 := ctx.DB().QueryContext(ctx, "SELECT *FROM Owner WHERE Phone=?", Contact)
		if q1 == nil {
			q1, err3 = ctx.DB().QueryContext(ctx, "INSERT INTO Owner(ownerName,Contact,Address) VALUES(?,?,?) ", Owner, Contact, Address)
		}
		if err1 != nil && err2 != nil && err3 != nil {
			return nil, err
		}
		return "hello", nil

	})
	app.POST("/job_schedule/delete/{Job_id}", func(ctx *gofr.Context) (interface{}, error) {
		Job_Id := ctx.PathParam("Job_id")
		rows, err := ctx.DB().QueryContext(ctx, "DELETE FROM Job_schedule WHERE jobId=?", Job_Id)
		if err != nil {
			return nil, err
		}
		fmt.Println("Record Removed")
		return rows, nil
	})
	app.PUT("/job_schedule/update_status/JobId:{jobId}?{Status}", func(ctx *gofr.Context) (interface{}, error) {
		var update_status map[string]interface{}
		err := json.Unmarshal([]byte(models.UpdateStatus), &update_status)
		Job_Id := ctx.PathParam("Job_Id")
		Status := ctx.PathParam("Status")
		rows, err := ctx.DB().QueryContext(ctx, "UPDATE Job_schedule SET Status=? WHERE jobId=?", Status, Job_Id)
		if err != nil {
			return nil, err
		}
		fmt.Println("Record Removed")
		return rows, nil
	})
}
