package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load environment variables
	err1 := godotenv.Load(".env")
	if err1 != nil {
		fmt.Println("Could not load env file")
		panic(err1.Error())
	}

	// Get environment variables
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_DATABASE := os.Getenv("DB_DATABASE")

	// Create a connection to the database
	database_uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)
	db, err := sql.Open("mysql", database_uri)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to make sure the connection is still active
	connect := db.Ping()
	if connect != nil {
		fmt.Println("Could not connect to database")
		panic(connect.Error())
	}

	// Select data from the database
	users, err := db.Query(`SELECT * from Users`)
	if err != nil {
		fmt.Println(("Could not get users"))
		panic(err.Error())
	}

	fmt.Println(users)
	defer users.Close()

	fmt.Println("Connection successful")
}
