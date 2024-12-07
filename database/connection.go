package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// Conn is the global database connection instance
var Conn *pgx.Conn

// Connect initializes the database connection
func Connect() {
	var err error
	// PostgreSQL connection string
	url := "postgres://aritrach6078:Aritra@database123@localhost:5432/product_management?sslmode=disable"
	Conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to the database!")
}
