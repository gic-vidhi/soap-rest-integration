package config

import (
	"os"
)

type config struct {
	SOAP_REST_PORT string
	MONGO_URI      string
	DB_NAME        string
	LOG_FILE_PATH  string
}

var APP_CONFIG config

func init() {
	port := os.Getenv("SOAP_REST_PORT")

	mongo_uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("DB_NAME")

	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath == "" {
		logFilePath = "./log.txt"
	}

	APP_CONFIG = config{
		SOAP_REST_PORT: port,
		MONGO_URI:      mongo_uri,
		DB_NAME:        db_name,
	}

}
