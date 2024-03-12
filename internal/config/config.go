package config

import (
	"github.com/joho/godotenv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP               HTTPConfig
	GIN                GINConfig
	CyberleninkaClient ClientConfig
	ELibraryClient     ClientConfig
}

type (
	HTTPConfig struct {
		Host            string
		Port            string
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		ShutdownTimeout time.Duration
	}

	GINConfig struct {
		Mode string
	}

	ClientConfig struct {
		BaseURL string
		Timeout time.Duration
	}
)

const (
	envFile   = ".env"
	directory = "configs"
	ymlFile   = "main"
)

func New() (*Config, error) {
	if err := godotenv.Load(envFile); err != nil {
		return nil, err
	}

	viper.AddConfigPath(directory)
	viper.SetConfigName(ymlFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	if err := envconfig.Process("gin", &config.GIN); err != nil {
		return nil, err
	}

	return config, nil
}
