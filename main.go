package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Model represents a sample data structure for demonstration
type Model struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	// Dependency management tools: Godep, Sltr, etc. (not used in this example)

	// Go templating language (not used in this example)

	// Go code generation tools, such as Stringer (not used in this example)

	// Popular Go web frameworks/libraries
	// Router packages: Gorilla Mux
	r := mux.NewRouter()

	// SQL and NoSQL database setup
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table for demonstration
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS models (
			id INTEGER PRIMARY KEY,
			name TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Common Goroutine and channel patterns (not used in this example)

	// MVC, RESTful, ORM
	r.HandleFunc("/models", func(w http.ResponseWriter, r *http.Request) {
		// Sample endpoint demonstrating CRUD operations with an ORM (sqlx)
		var models []Model
		err := db.Select(&models, "SELECT * FROM models")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Display models
		fmt.Fprintf(w, "Models: %+v", models)
	}).Methods("GET")

	// Basic experience with Reactjs (not used in this example)

	// Code versioning tools: Git (not used in this example)

	// Basic knowledge of macOS and terminal commands (not used in this example)

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
