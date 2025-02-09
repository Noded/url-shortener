package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"url-shortener/internal/db"
	"url-shortener/internal/handlers"
)

func TestHandleShortenURL(t *testing.T) {
	db.InitDB()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/short?url=https://example.com", nil)

	handlers.HandleShortenURL()
	http.DefaultServeMux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHandleListURLs(t *testing.T) {
	db.InitDB()
	db.AddUrl("https://example.com", "short.ly/abc")
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/list", nil)

	handlers.HandleListURLs()
	http.DefaultServeMux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHandleRedirectURL(t *testing.T) {
	db.InitDB()
	db.AddUrl("https://example.com", "short.ly/abc")
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/original?url=short.ly/abc", nil)

	handlers.HandleRedirectURL()
	http.DefaultServeMux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

// TODO: Make test for delete handler
