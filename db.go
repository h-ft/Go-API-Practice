package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test-devices"
	DB_HOST     = "localhost"
	DB_PORT     = "5433"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	fmt.Println("Connected to DB")
	return db
}

// A struct of Device which has 3 properties: Name, Brand, and Year
type Device struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Brand string `json:"Brand"`
	Year  string `json:"Year"`
}

// A global Device array to simulate a database
var Devices = []Device{
	{ID: "1", Name: "Macbook Pro", Brand: "Apple", Year: "2021"},
	{ID: "2", Name: "XPS", Brand: "Dell", Year: "2019"},
	{ID: "3", Name: "Thinkpad", Brand: "Lenovo", Year: "2020"},
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
