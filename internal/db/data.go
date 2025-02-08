package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var database *sql.DB

// InitDB establishes a connection to the database
// Notice: Exec this func once in program
func InitDB() error {
	var err error
	database, err = sql.Open("sqlite3", "urls.db")
	if err != nil {
		return err
	}

	// Checking connection
	if err := database.Ping(); err != nil {
		return err
	}

	// Create Table if not exists
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			shortUrl TEXT UNIQUE NOT NULL,
			url TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}

// CloseDB Close db
func CloseDB() {
	if database != nil {
		_ = database.Close()
		log.Println("Database connection closed")
	}
}

// AddUrl Adds a new URL to the database
func AddUrl(originalURL string, shortUrl string) error {
	_, err := database.Exec("INSERT INTO Urls (url, shortUrl) VALUES (?, ?)",
		originalURL, shortUrl)
	if err != nil {
		return err
	}
	return nil
}

// RetrieveOriginalURL Gets original URL from target shortURL
func RetrieveOriginalURL(shortURL string) (string, error) {
	var originalUrl string
	err := database.QueryRow("SELECT url FROM Urls WHERE shortUrl = ?", shortURL).Scan(&originalUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

// ListShortenedURLs Returns all short urls
func ListShortenedURLs() ([]string, error) {
	var urls []string
	rows, err := database.Query("SELECT shortUrl FROM Urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var shortUrl string
		if err := rows.Scan(&shortUrl); err != nil {
			return nil, err
		}
		urls = append(urls, shortUrl)
	}
	return urls, nil
}
