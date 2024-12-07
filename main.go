package main

import (
	"fmt"
	"net/http"

	"github.com/aritrach6078/product-management-system/api"
	"github.com/aritrach6078/product-management-system/database"
)

func main() {
	// Initialize the database connection
	database.Connect()

	// Register routes for API endpoints
	http.HandleFunc("/users", api.GetAllUsers)             // Get all users
	http.HandleFunc("/products", api.GetProducts)          // Get all products
	http.HandleFunc("/products/create", api.CreateProduct) // Create a new product

	// Start the HTTP server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
