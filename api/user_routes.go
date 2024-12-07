package api

import (
	"encoding/json"
	"net/http"

	"github.com/aritrach6078/product-management-system/database"
	"github.com/aritrach6078/product-management-system/models"
)

// GetAllUsers fetches all users from the database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.FetchAllUsers(database.Conn)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetProducts fetches all products from the database
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.FetchAllProducts(database.Conn)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CreateProduct adds a new product to the database
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = models.AddProduct(database.Conn, product)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
