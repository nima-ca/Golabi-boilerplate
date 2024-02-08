package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

var S3Session *session.Session
var S3Uploader *s3manager.Uploader

func CreateS3Client() {
	session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(viper.GetString("AWS_ACCESS_KEY_ID"),
			viper.GetString("AWS_SECRET_ACCESS_KEY"), ""),
		Region:   aws.String("default"),
		Endpoint: aws.String(viper.GetString("AWS_URL")),
	})

	if err != nil {
		panic(err)
	}

	// Put session in global variable
	S3Session = session

	// Put uploader in global variable
	S3Uploader = s3manager.NewUploader(session)
}

func GetS3Session() *session.Session {
	return S3Session
}

func GetS3Uploader() *s3manager.Uploader {
	return S3Uploader
}
