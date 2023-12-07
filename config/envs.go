// Package config project environment setup and bootstrap files
package config

import (
	"log"

	"github.com/spf13/viper"
)

// Env contains all environment variables as friendly props
type Env struct {
	ServerPort         string `mapstructure:"SERVER_PORT"`
	Environment        string `mapstructure:"APP_MODE"`
	DBUsername         string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenTTL     string `mapstructure:"ACCESS_TOKEN_TTL"`
	RefreshTokenSecret string `mapstructure:"REFRESH_SECRET"`
	RefreshTokenTTL    string `mapstructure:"REFRESH_TTL"`
	RedisHost          string `mapstructure:"REDIS_HOST"`
	RedisPort          string `mapstructure:"REDIS_PORT"`
}

// LoadEnvs loads environment variables from root .env file
func LoadEnvs() *Env {
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

	return &env
}
