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

func GetTodoServiceUri() string {
	return getEnvironmentValue("TODO_SERVICE_URI")
}

func GetHttpPort() int {
	portStr := getEnvironmentValue("HTTP_SRV_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}
