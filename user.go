package main

import (
	"database/sql"
	"fmt"
)

func CreateUserTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	fmt.Println("Table created (if not exists)")
	return nil
}

func InsertUser(db *sql.DB, name string, age int) error {
	query := `INSERT INTO users (name, age) VALUES ($1, $2)`
	_, err := db.Exec(query, name, age)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	fmt.Printf("Inserted user: %s (%d)\n", name, age)
	return nil
}

func QueryUsers(db *sql.DB) error {
	rows, err := db.Query(`SELECT id, name, age FROM users`)
	if err != nil {
		return fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	fmt.Println("Users in DB:")
	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			return err
		}
		fmt.Printf("ID: %d | Name: %s | Age: %d\n", id, name, age)
	}
	return nil
}
