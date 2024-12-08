package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aritrach6078/product-management-system/database"
	"github.com/aritrach6078/product-management-system/models"
	"github.com/streadway/amqp" // Import RabbitMQ package
)

// Helper function to publish a message to RabbitMQ
func publishToQueue(imageURL string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"image_processing_queue", // Queue name
		true,                     // Durable
		false,                    // Delete when unused
		false,                    // Exclusive
		false,                    // No-wait
		nil,                      // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.Publish(
		"",     // Exchange
		q.Name, // Routing key (queue name)
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strings.TrimSpace(imageURL)), // Clean up any extra spaces
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf("Published message to queue: %s", imageURL)
}

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

	// Publish product_images URL to RabbitMQ queue
	if len(product.ProductImages) > 0 {
		publishToQueue(product.ProductImages)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
