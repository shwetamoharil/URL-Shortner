package main

import (
	"URL-Shortner/controllers"
	"URL-Shortner/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	utils.DatabaseClient = utils.ConnectDB()
	router.HandleFunc("/encode", utils.SetCorsHeaders(controllers.EncodeUrls)).Methods("POST")
	http.ListenAndServe(":8000", router)
}
