package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// Config stores configuration values for the application.
type Config struct {
	AppEnv               string
	ContractsProviderUrl string
	AccountCreatorUrl    string
	StorageUrl           string
}

var Conf *Config

// New creates a new Config object with values from environment variables or default values.
func New() {
	loadEnv()
	Conf = &Config{
		AppEnv:               getStr("APP_ENV", "dev"),
		ContractsProviderUrl: getStr("CONTRACTS_PROVIDER_URL", "http://bridge-contracts-provider:8000/v1/get-contracts"),
		AccountCreatorUrl:    getStr("ACCOUNT_CREATOR_URL", "http://bridge-accounts-creator:8000/v1/create-account"),
		StorageUrl:           getStr("STORAGE_URL", "http://bridge-storage:8000/v1/create-account"),
	}
}

// loadEnv loads environment variables from the .env file.
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Errorf("error: %v", err)
	}
}

// getStr retrieves a string value from the environment variables or returns a default value.
func getStr(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getNum retrieves an integer value from the environment variables or returns a default value.
func getNum(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		num, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return num
	}
	return defaultValue
}

func (c *Config) IsDevApp() bool {
	return c.AppEnv == "dev"
}
