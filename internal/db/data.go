package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SQLStorage struct {
	data *sql.DB
}

func NewSQLStorage() *SQLStorage {
	return &SQLStorage{}
}

// InitDB establishes a connection to the database
// Notice: Exec this func once in program
func (s *SQLStorage) InitDB() error {
	var err error
	s.data, err = sql.Open("sqlite3", "urls.db")
	if err != nil {
		return err
	}

	// Checking connection
	if err := s.data.Ping(); err != nil {
		return err
	}

	// Create Table if not exists
	_, err = s.data.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			shortUrl TEXT UNIQUE NOT NULL,
			url TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("error create table: %v", err)
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}

// CloseDB Close db
func (s *SQLStorage) CloseDB() {
	if s.data != nil {
		_ = s.data.Close()
		log.Println("Database connection closed")
	}
}

// AddUrl Adds a new URL to the database
func (s *SQLStorage) AddUrl(originalURL string, shortUrl string) error {
	_, err := s.data.Exec("INSERT INTO Urls (url, shortUrl) VALUES (?, ?)",
		originalURL, shortUrl)
	if err != nil {
		return err
	}
	return nil
}

// RetrieveOriginalURL Gets original URL from target shortURL
func (s *SQLStorage) RetrieveOriginalURL(shortURL string) (string, error) {
	var originalUrl string
	err := s.data.QueryRow("SELECT url FROM urls WHERE shortUrl = ?", shortURL).Scan(&originalUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

// ListShortenedURLs Returns all short urls
func (s *SQLStorage) ListShortenedURLs() ([]string, error) {
	var urls []string
	rows, err := s.data.Query("SELECT shortUrl FROM Urls")
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

// DeleteURL Delete target URL
func (s *SQLStorage) DeleteURL(shortURL string) error {
	_, err := s.data.Exec("DELETE FROM urls WHERE shortUrl = ?", shortURL)
	if err != nil {
		return err
	}

	return nil
}
