package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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

func (s *S3Adapter) GetObject(key string) ([]byte, error) {
	return nil, nil
}

func (s *S3Adapter) PutObject(key string, data []byte) error {
	return nil
}
