package controller

import (
	"net/http"
	"soap-rest-integration/constants"
	"soap-rest-integration/helper"
	"soap-rest-integration/service"
)

func GetFromSoap(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, constants.METHOD_NOT_ALLOWED)
		return
	}

	country := r.URL.Query().Get("country")

	if country == "" {
		helper.RespondWithError(w, http.StatusBadRequest,
			"Country parameter should be provided in the URL")
		return
	}

	if filePath, err := service.CallGetSoap(country); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, constants.SOAP_CALLED_SUCCESSFULLY, helper.StatusCodeOK, filePath)
	}

}

func PostToSoap(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, constants.METHOD_NOT_ALLOWED)
		return
	}

	country := r.URL.Query().Get("country")

	if country == "" {
		helper.RespondWithError(w, http.StatusBadRequest,
			"Country parameter should be provided in the URL")
		return
	}

	if filePath, err := service.CallPostSoap(country); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, constants.SOAP_CALLED_SUCCESSFULLY, helper.StatusCodeOK, filePath)
	}

}
