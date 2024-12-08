This repository contains the backend system for a Product Management Application. The project includes features like asynchronous image processing, caching, logging, scalable architecture, and cloud-based storage. Built using Golang, PostgreSQL, Redis, RabbitMQ, and AWS S3, the application emphasizes high performance, modular design, and cloud-based integration.

📜 Features
Asynchronous Image Processing: Efficient handling of image uploads, resizing, and transformations using RabbitMQ for messaging.
Cloud Storage: Uploaded and processed images are stored securely in AWS S3 buckets.
Health Check API: Added /healthcheck endpoint to monitor system health.
Logging: Real-time logging to monitor and debug application behavior.
Caching: Faster response times using Redis.
Scalable Architecture: Easily expandable to accommodate future features or integrations.
Database Management: Secure and optimized interaction with PostgreSQL.
API Endpoints:
GET /users: Fetch all users.
POST /products: Add a new product.
GET /products: Fetch all products.
🚀 Tech Stack
Backend: Golang
Database: PostgreSQL
Caching: Redis
Message Broker: RabbitMQ
Cloud Storage: AWS S3
API Testing: Postman
📂 Project Structure
perl
Copy code
product-management-system/
├── api/                 # API route handlers
├── config/              # Application configuration files
├── database/            # Database connection and queries
├── image_processing/    # Image processing logic
├── middleware/          # Middleware for request validation, logging, etc.
├── models/              # Data models for Users and Products
├── services/            # Service layer for business logic
├── tests/               # Unit and integration tests
├── go.mod               # Go module dependencies
├── go.sum               # Module checksums
└── main.go              # Entry point of the application
🛠️ Installation
Prerequisites
Install Golang (v1.18 or later).
Install and set up PostgreSQL.
Install Redis and RabbitMQ.
Configure AWS S3:
Create an AWS S3 bucket (e.g., product-management-aritra).
Configure the bucket region (e.g., ap-south-1).
Set up AWS IAM user with S3 access permissions and generate access/secret keys.
Steps
Clone the Repository:

bash
Copy code
git clone https://github.com/aritrach6078/product-management-system.git
cd product-management-system
Install Dependencies:

bash
Copy code
go mod tidy
Set Up the Database:

Create a PostgreSQL database named product_management.
Update the connection string in database/connection.go with your credentials:
go
Copy code
"user=your_user password=your_password dbname=product_management sslmode=disable"
Configure AWS S3:

Update the S3 bucket name in main.go under the uploadToS3 function:
go
Copy code
bucketName := "product-management-aritra"
Ensure your AWS credentials are configured locally using AWS CLI or environment variables.
Run the Application:

bash
Copy code
go run main.go
📋 API Documentation
Users
Fetch All Users
Endpoint: GET /users
Description: Retrieves all users from the database.
Response Example:
json
Copy code
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
Request Body Example:
json
Copy code
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
Response Example:
json
Copy code
[
  {
    "id": 1,
    "user_id": 1,
    "product_name": "Tablet",
    "product_description": "A sleek and portable tablet",
    "product_images": "tablet.jpg",
    "product_price": 599.99,
    "compressed_product": "compressed_tablet.jpg"
  }
]
🔍 Testing
Steps to Test
Use Postman:
GET /users: Fetch all users.
POST /products: Add a product.
GET /products: Fetch all products.
Verify Database:
Use pgAdmin or PostgreSQL CLI to verify database entries.
Health Check:
Check GET /healthcheck for system status.
🌩️ AWS S3 Integration
Functionality:
Uploaded images are stored securely in AWS S3.
Processed images (resized/compressed) are uploaded back to S3 for efficient storage and retrieval.
Configuration:
The uploadToS3 function handles uploading images to the configured S3 bucket.
Ensure your S3 bucket has appropriate permissions for public/private access as per your requirements.
🧑‍💻 Author
Aritra Choudhary
Email: ad7342@srmist.edu.in
