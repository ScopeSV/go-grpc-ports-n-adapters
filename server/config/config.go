package config

import (
	"log"
	"os"
	"strconv"
)

func getEnvironmentValue(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("environment variable: %s is not set", key)
	}

	return value
}

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}
