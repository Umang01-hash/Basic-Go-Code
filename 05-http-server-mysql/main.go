package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type data struct {
	EnrollmentNumber int    `json:"enrollmentNumber"`
	Name             string `json:"name"`
}

// conn creates connection to the SQL database
func conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/Students")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := conn()
	if err != nil {
		fmt.Printf("Couldn't initialize database")
	}

	// Define a handler function to create a student entry
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		var d data

		// Reading the request body
		reqBody, _ := io.ReadAll(r.Body)

		// Converting json body to go struct
		err = json.Unmarshal(reqBody, &d)
		if err != nil {
			fmt.Println("Unmarshal error", err)
		}

		// Query inserts the student details in db
		_, err := db.Exec("INSERT INTO students (enrollment_number,name) values(?,?)", d.EnrollmentNumber, d.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Student added successfully!"))
	})

	// Define a handler function to fetch entries from student table and return as HTTP response
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		// Query returns all matching rows as a Rows struct your code can loop over
		rows, err := db.Query("SELECT name from students;")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Generate a dynamic response based on the retrieved data
		responseText := "Hello from Go Web Server!\n"

		// Next prepares the next result row for reading with the Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it
		for rows.Next() {
			var name string

			// Scan copies the columns from the matched row into the values pointed at by dest
			err := rows.Scan(&name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			responseText += fmt.Sprintf("User: %s\n", name)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseText))
	})

	port := 8080
	serverAddress := fmt.Sprintf(":%d", port)

	// Inform the user that the server is running
	fmt.Printf("Server is running on http://localhost%s\n", serverAddress)

	// Start the HTTP server
	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		fmt.Printf("Server encountered an error: %s\n", err)
	}
}
