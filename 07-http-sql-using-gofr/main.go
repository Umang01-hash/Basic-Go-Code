package main

import (
	"strings"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource"
)

type data struct {
	EnrollmentNumber int    `json:"enrollmentNumber"`
	Name             string `json:"name"`
}

func main() {
	// Creating GoFr's instance
	app := gofr.New()

	// POST endpoint using GoFr
	app.POST("/students", func(ctx *gofr.Context) (interface{}, error) {
		var d data

		// Reading and converting request data to go struct type
		err := ctx.Bind(&d)
		if err != nil {
			return nil, err
		}

		// Insert operation for database
		_, dbErr := ctx.SQL.Exec("INSERT INTO students (enrollment_no,name) values(?,?)", d.EnrollmentNumber, d.Name)
		if dbErr != nil {
			return nil, datasource.ErrorDB{Err: dbErr}
		}

		return "Student added successfully!", nil
	})

	// GET endpoint using GoFr
	app.GET("/students", func(ctx *gofr.Context) (interface{}, error) {
		// Query returns all matching rows
		rows, err := ctx.SQL.Query("SELECT name from students;")
		if err != nil {
			return nil, datasource.ErrorDB{Err: err}
		}
		defer rows.Close()

		var responseText string

		// Next prepares the next result row for reading with the Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it
		for rows.Next() {
			var name string

			// Scan copies the rows fetched
			err := rows.Scan(&name)
			if err != nil {
				return nil, datasource.ErrorDB{Err: err}
			}

			responseText += name + ","
		}

		// Trim the last ,
		res := strings.TrimRight(responseText, ",")

		return res, nil
	})

	// Starting the server
	app.Run()
}
