package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() {
	var err error
	url := "postgres://aritrach6078:Aritra@database123@localhost:5432/product_management?sslmode=disable"
	Conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to the database!")
}
