package database

import (
	"log"

	"github.com/aritrach6078/product-management-system/models"
)

// InsertProduct inserts a new product into the database
func InsertProduct(product models.Product) error {
	query := `
		INSERT INTO products (user_id, product_name, product_description, product_images, product_price, compressed_product)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := DB.Exec(
		query,
		product.UserID,
		product.ProductName,
		product.ProductDescription,
		product.ProductImages,
		product.ProductPrice,
		product.CompressedProduct,
	)

	if err != nil {
		log.Printf("Failed to insert product: %v", err)
		return err
	}

	log.Println("Product inserted successfully")
	return nil
}

// GetProducts retrieves products by user ID
func GetProducts(userID int) ([]models.Product, error) {
	query := `
		SELECT id, user_id, product_name, product_description, product_images, product_price, compressed_product
		FROM products
		WHERE user_id = $1`

	rows, err := DB.Query(query, userID)
	if err != nil {
		log.Printf("Failed to retrieve products: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice, &product.CompressedProduct)
		if err != nil {
			log.Printf("Failed to scan product: %v", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
