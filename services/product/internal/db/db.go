package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    log.Println("Connected to PostgreSQL database")
}
