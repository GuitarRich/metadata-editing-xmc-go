package sitecore

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return os.Getenv(key)
}

func ToString(field interface{}) string {
	if field == nil {
		return ""
	}
	return fmt.Sprintf("%s", field)
}

func AddIfNotEmpty(name string, value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%s=\"%s\"", name, value)
}
