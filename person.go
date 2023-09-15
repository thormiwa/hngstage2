package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS person (
		id INTEGER PRIMARY KEY,
		name TEXT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating 'person' table: %v", err)
		return err
	}

	log.Println("Database and table 'person' initialized successfully.")
	return nil
}