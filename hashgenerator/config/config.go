package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load("redis.env")

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	return os.Getenv(key)
}
