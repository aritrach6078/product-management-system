package models

import (
	"database/sql"
	"fmt"
)

// User represents the structure of the user table
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Product represents the structure of the product table
type Product struct {
	ID                 int     `json:"id"`
	UserID             int     `json:"user_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductImages      string  `json:"product_images"`
	ProductPrice       float64 `json:"product_price"`
	CompressedProduct  string  `json:"compressed_product"` // For storing compressed image URL
}

// FetchAllUsers retrieves all users from the database
func FetchAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FetchAllProducts retrieves all products from the database
func FetchAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, user_id, product_name, product_description, product_images, product_price, compressed_product FROM products")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice, &product.CompressedProduct)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// AddProduct inserts a new product into the database
func AddProduct(db *sql.DB, product Product) error {
	_, err := db.Exec(
		`INSERT INTO products (user_id, product_name, product_description, product_images, product_price, compressed_product)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice, product.CompressedProduct,
	)
	if err != nil {
		fmt.Println("Error adding product to the database:", err)
		return err
	}
	return nil
}
