package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListBuckets() {
	// Load the AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// List S3 buckets
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("Name: %s, Creation Date: %s\n", *bucket.Name, bucket.CreationDate)
	}
}

func UploadFileToS3(filePath, bucketName, objectKey string) error {
	// Load the AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %v", err)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Open the file for uploading
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file %v: %v", filePath, err)
	}
	defer file.Close()

	// Upload the file
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("unable to upload file to S3, %v", err)
	}

	log.Printf("File uploaded successfully to bucket %s with key %s", bucketName, objectKey)
	return nil
}
