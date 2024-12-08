package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aritrach6078/product-management-system/database"
	"github.com/aritrach6078/product-management-system/models"
	"github.com/gorilla/mux" // Using Gorilla Mux for routing
	"github.com/streadway/amqp"
)

// publishToQueue publishes a message to RabbitMQ
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
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strings.TrimSpace(imageURL)),
		},
	)
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
	} else {
		log.Printf("Published message to queue: %s", imageURL)
	}
}

// GetAllUsers fetches all users from the database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.FetchAllUsers(database.DB)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetProducts fetches all products from the database
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.FetchAllProducts(database.DB)
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

	err = models.AddProduct(database.DB, product)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	// Publish product_images URLs to RabbitMQ queue
	if len(product.ProductImages) > 0 {
		for _, imageURL := range strings.Split(product.ProductImages, ",") {
			publishToQueue(strings.TrimSpace(imageURL))
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// Router sets up the API routes and returns the router
func Router() *mux.Router {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")

	// Add a health check route
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return router
}

// HealthCheck verifies the server and database connectivity
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request: GET /health")

	// Check database connection
	err := database.DB.Ping()
	if err != nil {
		log.Printf("Database health check failed: %v", err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	log.Println("Health check passed")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
