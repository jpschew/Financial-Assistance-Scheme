package config

import (
	"github.com/spf13/viper"
	"log"
)

// AppConfig holds the structure for the application configuration
type AppConfig struct {
	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`

	JWT struct {
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"jwt"`
}

var appConfig AppConfig

// InitConfig initializes the application configuration using Viper
func InitConfig() AppConfig {
	// Set the file name of the configuration file
	viper.SetConfigName("config")
	// Set the path to look for the configuration file
	viper.AddConfigPath("./config")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	// Set the type of the configuration file
	viper.SetConfigType("yaml")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Unmarshal the configuration into AppConfig struct
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return appConfig
}

func GetConfig() *AppConfig {
	return &appConfig
}
