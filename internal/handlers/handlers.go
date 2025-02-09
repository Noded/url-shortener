package handlers

import (
	"net/http"

	"url-shortener/internal/db"
	"url-shortener/internal/shorten"
)

// HandleShortenURL Create new short URL and save to db
// Example: curl "http://localhost:8080/short?url=https://www.example.com/very/long/url/that/needs/shortening"
func HandleShortenURL() {
	http.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		URL := r.URL.Query().Get("url") // Getting url from request
		if URL == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}

		shortURL := shorten.ShortingURL(URL) // launch func to shorting URL
		db.AddUrl(URL, shortURL)
		w.Write([]byte(shortURL + "\n"))
	})
}

// HandleListURLs Get list of URLs
//
//	Example: curl "http://localhost:8080/list"
func HandleListURLs() {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		urls, err := db.ListShortenedURLs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else if urls != nil {
			for _, url := range urls {
				w.Write([]byte(url + "\n"))
			}
		} else {
			w.Write([]byte("None\n"))
		}
	})
}

// HandleRedirectURL Returns original url from short URL
//
//	Example: curl "http://localhost:8080/original?url=someShortURL"
func HandleRedirectURL() {
	http.HandleFunc("/original", func(w http.ResponseWriter, r *http.Request) {
		origURL, err := db.RetrieveOriginalURL(r.URL.Query().Get("url")) // Getting real url from short version
		if err != nil {
			http.Error(w, "URL not find", http.StatusNotFound)
		}
		w.Write([]byte(origURL + "\n"))
	})
}

// HandleDeleteURL TODO: Make delete request
func HandleDeleteURL() {
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		URL := r.URL.Query().Get("url")
		if URL == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
		}
		w.Write([]byte("Deleted " + URL + "\n"))
		db.DeleteURL(URL)

	})
}
