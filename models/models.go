package models

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type URL struct {
	UrlString string `json:"url"`
}

type EncodeUrls struct {
	EncodeUrl string `json:"encodedUrl"`
	ShortUrl  string `json:"shortUrl"`
}
