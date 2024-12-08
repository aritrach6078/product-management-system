package main

import (
	"context"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
)

// Process an image from the message
func processImage(message string) {
	trimmedMessage := strings.TrimSpace(message) // Remove extra whitespace or newline
	log.Printf("Processing image URL: %s", trimmedMessage)

	// Step 1: Download the image from the URL
	response, err := http.Get(trimmedMessage)
	if err != nil {
		log.Printf("Failed to download image: %v", err)
		return
	}
	defer response.Body.Close()

	// Step 2: Decode the image
	var img image.Image
	var format string

	img, format, err = image.Decode(response.Body)
	if err != nil {
		log.Printf("Failed to decode image: %v", err)
		return
	}
	log.Printf("Image format: %s", format)

	// Step 3: Resize the image (100x100 pixels for demonstration)
	resizedImg := resize.Resize(100, 100, img, resize.Lanczos3)

	// Step 4: Save the processed image to disk
	outputFileName := "processed_image." + format
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Printf("Failed to create output file: %v", err)
		return
	}
	defer outputFile.Close()

	switch format {
	case "jpeg":
		err = jpeg.Encode(outputFile, resizedImg, nil)
	case "png":
		err = png.Encode(outputFile, resizedImg)
	default:
		log.Printf("Unsupported format: %s", format)
		return
	}

	if err != nil {
		log.Printf("Failed to encode processed image: %v", err)
		return
	}

	log.Printf("Processed image saved as '%s'", outputFileName)

	// Step 5: Upload the processed image to S3
	err = uploadToS3(outputFileName)
	if err != nil {
		log.Printf("Failed to upload image to S3: %v", err)
	} else {
		log.Printf("Image '%s' successfully uploaded to S3!", outputFileName)
	}
}

// Upload a file to S3
func uploadToS3(fileName string) error {
	// Load AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		return err
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Upload the file to the S3 bucket
	bucketName := "product-management-aritra"
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &fileName,
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

// Start RabbitMQ consumer
func startConsumer() {
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
		"image_processing_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			processImage(string(d.Body))
		}
	}()
	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Main function
func main() {
	log.Println("Starting RabbitMQ Consumer...")
	startConsumer()
}
