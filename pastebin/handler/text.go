package handler

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"log"
	"main/database"
	"main/model"
	"strings"
)

type BucketBasics struct {
	S3Client *s3.Client
}

func (basics BucketBasics) UploadFile(bucketName string, objectKey string, fileContent string) error {
	_, err := basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   strings.NewReader(fileContent),
	})
	if err != nil {
		log.Printf("Couldn't upload file to %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return err
}

func CreateText(c *fiber.Ctx) error {
	db := database.DB
	text := new(model.Text)

	if err := c.BodyParser(text); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create product",
			"data":    err,
		})
	}

	db.Create(&text)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created product",
		"data":    text,
	})
}
