package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB // Exported database connection instance

// InitDB initializes the database connection
func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=Aritra@database123 dbname=product_management sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Database connected successfully!")
}
