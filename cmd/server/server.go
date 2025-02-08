package main

import (
	"log"
	"net/http"

	"url-shortener/internal/db"
	"url-shortener/internal/handlers"
)

func main() {
	// Init DB
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Register handlers
	handlers.HandleShortenURL()
	handlers.HandleListURLs()
	handlers.HandleRedirectURL()

	http.ListenAndServe("localhost:8080", nil)
}
