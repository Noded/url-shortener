package handlers

import (
	"net/http"

	"url-shortener/internal/db"
	"url-shortener/internal/shorten"
)

// CreateShortURL Create new short url and save to db
func CreateShortURL() {
	http.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		URL := r.URL.Query().Get("url") // Getting url from request
		if URL == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}

		shortURL := shorten.ShortingURL(URL) // launch func to shorting url
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

func GetOrigURL() {
	http.HandleFunc("/original", func(w http.ResponseWriter, r *http.Request) {
		origURL, err := db.GetOriginalURL(r.URL.Query().Get("url")) // Getting real url from short version
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if origURL == "" {
			http.Error(w, "orig URL parameter is required", http.StatusBadRequest)
			return
		}

		w.Write([]byte(origURL + "\n"))
	})
}
