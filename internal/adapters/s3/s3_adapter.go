package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Adapter struct {
	client     *s3.Client
	bucketName string
}

func NewS3Adapter(accessKey, secretAccessKey, bucketName, region string) (*S3Adapter, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretAccessKey, "")),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return &S3Adapter{client: client, bucketName: bucketName}, nil
}

func (sa *S3Adapter) UploadFile(objectKey string, file io.Reader) error {
	uploader := manager.NewUploader(sa.client)
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(sa.bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	return err
}
