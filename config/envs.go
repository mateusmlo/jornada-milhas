// Package config project environment setup and bootstrap files
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"APP_MODE"`
	LogOutput   string `mapstructure:"LOG_OUTPUT"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

// LoadEnvs loads environment variables from root .env file
func LoadEnvs() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ failed to read envs")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ failed to load envs")
	}

	return env
}
