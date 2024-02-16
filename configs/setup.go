package config

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

func Setup() {
	wd, _ := os.Getwd()
	binDir := path.Join(wd, "bin")

	godotenv.Write(map[string]string{
		"PATH": binDir,
	}, "/ffmpeg.env")

	godotenv.Load(".env", "ffmpeg.env")
}

func Getenv(varName string) string {
	return os.Getenv(varName)
}
