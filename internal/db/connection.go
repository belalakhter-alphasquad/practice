package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Conn *sql.DB
	Err  error
}

func NewClient() *Connection {
	
	connStr := "root:demo@tcp(localhost:3306)/demo"

	conn, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Error preparing database connection: %v", err)
	}
	err = conn.Ping()
	if err != nil {
		log.Printf("Error establishing a connection with the database: %v", err)
		return &Connection{nil, err}
	}
	return &Connection{conn, nil}
}
