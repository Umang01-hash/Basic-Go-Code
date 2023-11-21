package main

import (
	"fmt"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type data struct {
	EnrollmentNumber int    `json:"enrollmentNumber"`
	Name             string `json:"name"`
}

func main() {
	app := gofr.New()

	app.POST("/post", func(ctx *gofr.Context) (interface{}, error) {
		var d data

		err := ctx.Bind(&d)
		if err != nil {
			return nil, err
		}

		_, dbErr := ctx.DB().Exec("INSERT INTO students (enrollment_number,name) values(?,?)", d.EnrollmentNumber, d.Name)
		if dbErr != nil {
			return nil, &errors.DB{Err: err}
		}

		return "Student added successfully!", nil
	})

	app.GET("/get", func(ctx *gofr.Context) (interface{}, error) {
		// Query returns all matching rows as a Rows struct your code can loop over
		rows, err := ctx.DB().Query("SELECT name from students;")
		if err != nil {
			return nil, errors.DB{Err: err}
		}
		defer rows.Close()

		// Generate a dynamic response based on the retrieved data
		responseText := "Hello from Go Web Server!"

		// Next prepares the next result row for reading with the Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it
		for rows.Next() {
			var name string

			// Scan copies the columns from the matched row into the values pointed at by dest
			err := rows.Scan(&name)
			if err != nil {
				return nil, errors.DB{Err: err}
			}

			responseText += fmt.Sprintf("\nUser: %s", name)
		}

		return responseText, nil
	})

	app.Start()
}
