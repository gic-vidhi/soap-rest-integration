package main

import (
	"log"
	"net/http"
	"os"
	"soap-rest-integration/config"
	"soap-rest-integration/router"
)

func main() {
	logFile, _ := os.OpenFile(config.APP_CONFIG.LOG_FILE_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close()
	log.SetOutput(logFile)

	r := router.Router()

	log.Printf("Server started at port %q ", config.APP_CONFIG.SOAP_REST_PORT)
	log.Printf("Connected to Mongo URI %q ", config.APP_CONFIG.MONGO_URI)

	if err := http.ListenAndServe(":"+config.APP_CONFIG.SOAP_REST_PORT, r); err != nil {
		log.Fatal(err)
	}
}
