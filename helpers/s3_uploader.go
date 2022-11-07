package helpers

import (
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "ae-main-bucket/prscreenshots/"
)

func UploadFileS3(reader io.Reader, name string) string {
	// The session the S3 Uploader will use
	sess, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		panic(err)
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(AWS_S3_BUCKET),
		Key:         aws.String(name),
		Body:        reader,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		log.Printf("failed to upload file, %v\n", err)
	}
	return aws.StringValue(&result.Location)
}
