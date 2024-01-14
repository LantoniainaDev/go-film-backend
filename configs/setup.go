package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Setup() {
	godotenv.Load(".env")
}

func Getenv(varName string) string {
	return os.Getenv(varName)
}
