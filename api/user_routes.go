package api

import (
	"encoding/json"
	"net/http"

	"github.com/aritrach6078/product-management-system/database"
	"github.com/aritrach6078/product-management-system/models"
)

// Function to fetch all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.FetchAllUsers(database.Conn)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Function to create a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Decode the JSON body into the product struct
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the model function to insert the product into the database
	err = models.AddProduct(database.Conn, product)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product created successfully"})
}
