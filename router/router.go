package router

import (
	"net/http"
	"soap-rest-integration/controller"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/soap-rest-integration-service/get-from-soap-api/", http.HandlerFunc(controller.GetFromSoap))
	router.Handle("/soap-rest-integration-service/post-to-soap-api/", http.HandlerFunc(controller.PostToSoap))

	return router
}
