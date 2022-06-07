package utils

import (
	"URL-Shortner/models"
	"encoding/json"
	"log"
	"net/http"
)

func handlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func HandleHttpErrors(w http.ResponseWriter, customErrMsg string, statusCode int, err error) (isError bool) {
	if err != nil {
		errorMsg := err.Error()
		if customErrMsg != "" {
			errorMsg = customErrMsg
		}
		response := models.ErrorResponse{
			Message: errorMsg,
			Status:  statusCode,
		}

		message, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return true
		}

		w.WriteHeader(response.Status)
		_, err = w.Write(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return true
		}
		return true
	}
	return
}
