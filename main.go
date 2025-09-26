package main

import (
	"fmt"
	"log"

	"github.com/aloneen/assign2-ex-1/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// Connect to DB
	db, err := initializers.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}
	defer db.Close()
	fmt.Println("Connected to DB")

	// Create table
	if err := CreateUserTable(db); err != nil {
		log.Fatal(err)
	}

	// Insert some users
	_ = InsertUser(db, "Dias", 22)
	_ = InsertUser(db, "Aigerim", 25)
	_ = InsertUser(db, "Rustem", 30)

	// Query all users
	if err := QueryUsers(db); err != nil {
		log.Println(err)
	}
}
