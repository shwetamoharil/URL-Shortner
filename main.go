package main

import (
	"URL-Shortner/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	utils.DatabaseClient = utils.ConnectDB()
	http.ListenAndServe(":8000", router)
}
