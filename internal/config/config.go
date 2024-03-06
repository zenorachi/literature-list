package config

import (
	"time"

	"github.com/joho/godotenv"

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

func init() {
	if err := godotenv.Load(envFile); err != nil {
		//logger.Fatal("config", ".env initialization failed")
	}

	viper.AddConfigPath(directory)
	viper.SetConfigName(ymlFile)
	if err := viper.ReadInConfig(); err != nil {
		//logger.Fatal("config", "viper initialization failed")
	}
}

func New() *Config {
	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		//logger.Fatal("viper config", err.Error())
	}

	if err := envconfig.Process("gin", &config.GIN); err != nil {
		//logger.Fatal("gin config", err.Error())
	}

	return config
}
