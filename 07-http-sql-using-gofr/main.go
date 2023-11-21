package main

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"strings"
)

type data struct {
	EnrollmentNumber int    `json:"enrollmentNumber"`
	Name             string `json:"name"`
}

func main() {
	// Creating GoFr's instance
	app := gofr.New()

	// POST endpoint using GoFr
	app.POST("/post", func(ctx *gofr.Context) (interface{}, error) {
		var d data

		// Reading and converting request data to go struct type
		err := ctx.Bind(&d)
		if err != nil {
			return nil, err
		}

		// Insert operation for database
		_, dbErr := ctx.DB().Exec("INSERT INTO students (enrollment_number,name) values(?,?)", d.EnrollmentNumber, d.Name)
		if dbErr != nil {
			return nil, &errors.DB{Err: err}
		}

		return "Student added successfully!", nil
	})

	// GET endpoint using GoFr
	app.GET("/get", func(ctx *gofr.Context) (interface{}, error) {
		// Query returns all matching rows
		rows, err := ctx.DB().Query("SELECT name from students;")
		if err != nil {
			return nil, errors.DB{Err: err}
		}
		defer rows.Close()

		var responseText string

		// Next prepares the next result row for reading with the Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it
		for rows.Next() {
			var name string

			// Scan copies the rows fetched
			err := rows.Scan(&name)
			if err != nil {
				return nil, errors.DB{Err: err}
			}

			responseText += name + ","
		}

		// Trim the last ,
		res := strings.TrimRight(responseText, ",")

		return res, nil
	})

	// Starting the server
	app.Start()
}
