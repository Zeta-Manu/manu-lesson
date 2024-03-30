package config

import "os"

type AWSConfig struct {
	AccessKey       string
	SecretAccessKey string
}

type CloudFrontConfig struct {
	Domain string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type HTTPConfig struct {
	Port string
}

type S3Config struct {
	BucketName string
	Region     string
}

type AppConfig struct {
	AWS        AWSConfig
	CloudFront CloudFrontConfig
	Database   DatabaseConfig
	HTTP       HTTPConfig
	S3         S3Config
}

func NewAppConfig() *AppConfig {
	awsConfig := AWSConfig{
		AccessKey:       os.Getenv("AWS_ACCESS_KEY"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	cloudFrontConfig := CloudFrontConfig{
		Domain: os.Getenv("CLOUDFRONT_DOMAIN"),
	}

	dbConfig := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	httpConfig := HTTPConfig{
		Port: os.Getenv("HTTP_PORT"),
	}

	S3Config := S3Config{
		BucketName: os.Getenv("S3_BUCKET"),
		Region:     os.Getenv("REGION"),
	}

	return &AppConfig{
		AWS:        awsConfig,
		CloudFront: cloudFrontConfig,
		Database:   dbConfig,
		HTTP:       httpConfig,
		S3:         S3Config,
	}
}
