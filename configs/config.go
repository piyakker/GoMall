package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBName        string
	DBUser        string
	DBPass        string
	RedisAddr     string
	RedisPassword string
	RabbitMQUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it, relying on environment variables")
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBName:        os.Getenv("DB_NAME"),
		DBUser:        os.Getenv("DB_USER"),
		DBPass:        os.Getenv("DB_PASS"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RabbitMQUrl:   os.Getenv("RABBITMQ_URL"),
	}
}
