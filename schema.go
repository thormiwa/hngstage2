// database.go
package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreatePerson(db *sql.DB, name string) (int64, error) {
	result, err := db.Exec("INSERT INTO persons (name) VALUES (?)", name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetPerson(db *sql.DB, id int) (*Person, error) {
	var person Person
	err := db.QueryRow("SELECT id, name FROM persons WHERE id = ?", id).Scan(&person.ID, &person.Name)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func UpdatePerson(db *sql.DB, id int, name string) error {
	_, err := db.Exec("UPDATE persons SET name = ? WHERE id = ?", name, id)
	return err
}

func DeletePerson(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM persons WHERE id = ?", id)
	return err
}
