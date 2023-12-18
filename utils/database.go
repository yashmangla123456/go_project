package utils

import (
	"database/sql"
	"fmt"
	"gofr.dev/pkg/gofr"
)

func ExecQuery(ctx *gofr.Context, query string, args ...interface{}) *sql.Rows {
	result, err := ctx.DB().QueryContext(ctx, query, args...)
	if err != nil {
		fmt.Println("=========hwibw")
		fmt.Println(err.Error())
	}
	return result
}

func ExecQueryRow(ctx *gofr.Context, query string, args ...interface{}) *sql.Row {
	return ctx.DB().QueryRowContext(ctx, query, args...)
}
