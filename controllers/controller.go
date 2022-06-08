package controllers

import (
	"URL-Shortner/models"
	"URL-Shortner/utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

func EncodeUrls(w http.ResponseWriter, r *http.Request) {
	var url models.URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if isError := utils.HandleHttpErrors(w, "", http.StatusInternalServerError, err); isError {
		return
	}

	err = utils.ValidateUrl(url.UrlString)
	if isError := utils.HandleHttpErrors(w, "invalid url", http.StatusOK, err); isError {
		return
	}

	collection, err := utils.GetCollection(utils.DatabaseClient, utils.ENCODED_URLS)
	if isError := utils.HandleHttpErrors(w, "could not fetch records", http.StatusInternalServerError, err); isError {
		return
	}

	var base64EncodedUrl string
	host := "http://localhost:8000/"
	base64EncodedUrl = base64.StdEncoding.EncodeToString([]byte(url.UrlString))

	encodedUrl := models.EncodeUrls{
		EncodeUrl: base64EncodedUrl,
		ShortUrl:  host + base64EncodedUrl[0:4],
	}

	_, err = collection.InsertOne(context.TODO(), encodedUrl)
	if isError := utils.HandleHttpErrors(w, "unable to insert", http.StatusInternalServerError, err); isError {
		return
	}

	var response struct {
		ShortUrl string `json:"shortUrl"`
		Message  string `json:"message"`
	}

	response.ShortUrl = encodedUrl.ShortUrl
	response.Message = "short url created successfully"
	err = json.NewEncoder(w).Encode(response)
	if isError := utils.HandleHttpErrors(w, "Unable to fetch records", http.StatusInternalServerError, err); isError {
		return
	}
}
