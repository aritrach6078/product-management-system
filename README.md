# product-management-system
This repository contains the backend system for a Product Management Application. The project includes features like asynchronous image processing, caching, logging, and scalable architecture. Built using Golang, PostgreSQL, Redis, RabbitMQ, and S3, the application emphasizes high performance and modular design.
## 📜 Features

- **Asynchronous Image Processing**: Efficient handling of image uploads and transformations.
- **Logging**: Real-time logging to monitor and debug application behavior.
- **Caching**: Faster response times using caching mechanisms with Redis.
- **Scalable Architecture**: Easily expandable to accommodate future features or integrations.
- **Database Management**: Secure and optimized interaction with PostgreSQL.
- **API Endpoints**:
  - `GET /users`: Fetch all users.
  - `POST /products`: Add a new product.
  - `GET /products`: Fetch all products.

---

## 🚀 Tech Stack

- **Backend**: Golang
- **Database**: PostgreSQL
- **Caching**: Redis
- **Message Broker**: RabbitMQ
- **Cloud Storage**: S3
- **API Testing**: Postman

---

## 📂 Project Structure

product-management-system/ ├── api/ # API route handlers ├── config/ # Application configuration files ├── database/ # Database connection and queries ├── image_processing/ # Image processing logic ├── middleware/ # Middleware for request validation, logging, etc. ├── models/ # Data models for Users and Products ├── services/ # Service layer for business logic ├── tests/ # Unit and integration tests ├── go.mod # Go module dependencies ├── go.sum # Module checksums └── main.go # Entry point of the application
---

## 🛠️ Installation

### Prerequisites

1. Install **Golang** (v1.18 or later).
2. Install and set up **PostgreSQL**.
3. Install **Redis** and **RabbitMQ**.
4. Configure **AWS S3** for cloud storage.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/aritrach6078/product-management-system.git
   cd product-management-system
Install dependencies:
go mod tidy


Set up the database:
 Create a PostgreSQL database named product_management.
 Update the connection string in database/connection.go with your credentials.



Start the application:
 go run main.go



📋 API Documentation
Users
Fetch All Users
 Endpoint: GET /users
 Description: Retrieves all users from the database.


[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com"
  }
]




Products
Add a Product
 Endpoint: POST /products
 Description: Adds a new product to the database.
 Request Body:


{
  "user_id": 1,
  "product_name": "Tablet",
  "product_description": "A sleek and portable tablet",
  "product_images": "tablet.jpg",
  "product_price": 599.99,
  "compressed_product": "compressed_tablet.jpg"
}



Fetch All Products
 Endpoint: GET /products
 Description: Retrieves all products from the database.


🔍 Testing
Use Postman to test the endpoints
 GET /users: Fetch all users.
 POST /products: Add a product.
 GET /products: Fetch all products.
Verify database changes using pgAdmin or the PostgreSQL command line.

🧑‍💻 Author
Aritra Choudhary
Email: ad7342@srmist.edu.in

