package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type HTTPConfig struct {
	Port string `mapstructure:"HTTP_PORT"`
}

type S3Config struct {
	BucketName string `mapstructure:"S3_BUCKET"`
	Region     string `mapstructure:"REGION"`
}

type AppConfig struct {
	Database DatabaseConfig
	HTTP     HTTPConfig
	S3       S3Config
}

func NewAppConfig() *AppConfig {
	viper.AutomaticEnv()

	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &appConfig
}
