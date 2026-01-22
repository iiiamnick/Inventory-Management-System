package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "user=manager database=postgres sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Cannot open the database", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect to the database", err)
		panic(err)

	}
	fmt.Println("Successfully connected to the database")
	DB = db
}
