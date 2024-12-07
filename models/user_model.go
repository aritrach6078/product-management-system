package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	UserID             int     `json:"user_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductImages      string  `json:"product_images"`
	ProductPrice       float64 `json:"product_price"`
	CompressedProduct  string  `json:"compressed_product"`
}

func FetchAllUsers(conn *pgx.Conn) ([]User, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, email FROM users")
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

// AddProduct inserts a new product into the products table
func AddProduct(conn *pgx.Conn, product Product) error {
	query := `
        INSERT INTO products (user_id, product_name, product_description, product_images, product_price, compressed_product)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := conn.Exec(context.Background(), query,
		product.UserID, product.ProductName, product.ProductDescription,
		product.ProductImages, product.ProductPrice, product.CompressedProduct,
	)
	if err != nil {
		fmt.Println("Error adding product:", err)
		return err
	}
	return nil
}
