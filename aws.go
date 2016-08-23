package main

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadToAws(file multipart.File, header *multipart.FileHeader) string {
	config := &aws.Config{
		Region:      aws.String(*region),
		Credentials: credentials.NewStaticCredentials(*accessID, *secretKey, ""),
	}
	s3client := s3.New(session.New(), config)

	params := &s3.PutObjectInput{
		Bucket:      aws.String(*bucket),
		Key:         aws.String(header.Filename),
		ACL:         aws.String("public-read"),
		Body:        file,
		ContentType: aws.String(header.Header.Get("Content-Type")),
	}

	_, err := s3client.PutObject(params)

	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("https://s3-%s.amazonaws.com/%s/%s", *region, *bucket, header.Filename)
	return url
}
