package s3

import (
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadToS3(file multipart.File, filename string, bucket string) error {
	// Get Uploader
	uploader := GetS3Uploader()

	// Upload file to S3
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}

func DeleteFromS3(filename string, bucket string) error {
	// Get Session
	session := GetS3Session()

	// Create an S3 client
	S3Client := s3.New(session)

	// Delete file from S3
	_, err := S3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		return err
	}

	return nil
}

func BatchDeleteFromS3(keys []string, bucket string) error {
	// Get Session
	session := GetS3Session()

	// Create an S3 client
	S3Client := s3.New(session)

	// Create Objects to delete
	var objects []*s3.ObjectIdentifier
	for _, key := range keys {
		objects = append(objects, &s3.ObjectIdentifier{
			Key: aws.String(key),
		})
	}

	// Delete files form S3
	_, err := S3Client.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &s3.Delete{
			Objects: objects,
			Quiet:   aws.Bool(false),
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func GeneratePresignedURL(bucket, filename string, expire time.Duration) (string, error) {
	// Get Session
	session := GetS3Session()

	// Create an S3 client
	S3Client := s3.New(session)

	req, _ := S3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})

	// Set an expiration time for the URL (e.g., 1 min)
	url, err := req.Presign(expire)
	if err != nil {
		return "", err
	}

	return url, nil
}
