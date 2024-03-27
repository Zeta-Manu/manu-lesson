package s3

import (
	"bytes"
	"context"
	"io"

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
	input := &s3.GetObjectInput{
		Bucket: &s.bucketName,
		Key:    &key,
	}

	resp, err := s.client.GetObject(context.Background(), input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *S3Adapter) PutObject(key string, data []byte) error {
	input := &s3.PutObjectInput{
		Bucket: &s.bucketName,
		Key:    &key,
		Body:   bytes.NewReader(data),
	}

	_, err := s.client.PutObject(context.Background(), input)
	return err
}
