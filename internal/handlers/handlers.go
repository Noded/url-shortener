package handlers

import (
	"net/http"

	"url-shortener/internal/db"
	"url-shortener/internal/shorten"
)

// CreateShortURL Create new short url and save to db
func CreateShortURL() {
	http.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url") // getting url from request
		if url == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}

		shortURL := shorten.ShortingURL(url) // launch func to shorting url
		//	example: curl "http://localhost:8080/short?url=https://www.example.com/very/long/url/that/needs/shortening"
		db.AddUrl(r.URL.Query().Get("url"), shortURL)
		w.Write([]byte(shortURL + "\n"))
	})
}

// GetShortURLs Get list of urls
func GetShortURLs() {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		urls, err := db.GetShortURLs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		for _, url := range urls {
			w.Write([]byte(url + "\n"))
		}
	})
}
