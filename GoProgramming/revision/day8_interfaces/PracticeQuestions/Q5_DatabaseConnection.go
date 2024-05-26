package main

import "fmt"

type Database interface {
	Connect()
	Query()
	Close()
}

type MySQL struct{}

func (db MySQL) Connect() {
	fmt.Println("Connecting to MySQL...")
}

func (db MySQL) Query() {
	fmt.Println("Querying MySQL database...")
}

func (db MySQL) Close() {
	fmt.Println("Closing MySQL connection.")
}

type PostgreSQL struct{}

func (db PostgreSQL) Connect() {
	fmt.Println("Connecting to PostgreSQL...")
}

func (db PostgreSQL) Query() {
	fmt.Println("Querying PostgreSQL database...")
}

func (db PostgreSQL) Close() {
	fmt.Println("Closing PostgreSQL connection.")
}

func main() {
	var db Database

	db = MySQL{}
	db.Connect()
	db.Query()
	db.Close()

	fmt.Println()

	db = PostgreSQL{}
	db.Connect()
	db.Query()
	db.Close()
}
