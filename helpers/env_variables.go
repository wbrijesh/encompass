package helpers

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVariable(key string) (value string, err error) {
	err = godotenv.Load()
	_ = HandleError(err, Panic)

	value = os.Getenv(key)

	return value, err
}
