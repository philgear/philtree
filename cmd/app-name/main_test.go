package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/philgear/philtree/internal/config"
	"github.com/philgear/philtree/internal/services"
)

func TestHomeHandler(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	linkService := services.NewLinkService(cfg)
	router := SetupRouter(linkService)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.Len() == 0 {
		t.Errorf("handler returned empty body")
	}
}

func TestStaticFileServing(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	linkService := services.NewLinkService(cfg)
	router := SetupRouter(linkService)

	req, err := http.NewRequest("GET", "/static/styles.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "text/css; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}
}
