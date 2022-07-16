package project_env

import (
	"github.com/joho/godotenv"
	"os"
)

func GoDotEnvVariable(key string) string {
	// load .env file
	projectFilePath := os.Getenv("ENV_FILE_PATH")

	if len(projectFilePath) == 0 {
		_ = godotenv.Load(".env")
	} else {
		_ = godotenv.Load(projectFilePath + "/.env")
	}

	return os.Getenv(key)
}
