package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

var Env *Config

func LoadEnv() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Env = &Config{
		PORT: getEnvOrPanic("PORT"),
	}
}

func getEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
