package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Insert data into a table
func insertData(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
	return err
}

// Retrieve data from a table
func retrieveData(db *sql.DB) error {
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var name, email string
		if err := rows.Scan(&name, &email); err != nil {
			return err
		}
		fmt.Println("Name:", name, "Email:", email)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

// Update data in a table
func updateData(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET email = $1 WHERE name = $2", "newemail@example.com", "John Doe")
	return err
}

// Delete data from a table
func deleteData(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE name = $1", "John Doe")
	return err
}

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Example usage of database operations
	if err := insertData(db); err != nil {
		log.Fatal(err)
	}

	if err := retrieveData(db); err != nil {
		log.Fatal(err)
	}

	if err := updateData(db); err != nil {
		log.Fatal(err)
	}

	if err := deleteData(db); err != nil {
		log.Fatal(err)
	}
}
