package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			userid INTEGER NOT NULL
		)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println("Could not create events table")
		panic(err)
	}
}
