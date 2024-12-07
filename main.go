package main

import (
	"fmt"
	"net/http"

	"github.com/aritrach6078/product-management-system/api"
	"github.com/aritrach6078/product-management-system/database"
)

func main() {
	database.Connect()

	http.HandleFunc("/users", api.GetAllUsers)
	http.HandleFunc("/products", api.CreateProduct)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
