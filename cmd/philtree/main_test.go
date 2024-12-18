package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/philgear/philtree/internal"
	"github.com/philgear/philtree/internal/config"
	"github.com/philgear/philtree/internal/services"
)

func TestHomeHandler(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	linkService := services.NewLinkService(cfg)
	router := internal.SetupRouter(linkService)

	tests := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{"Home Page", "/", http.StatusOK},
		{"Non-existent Page", "/notfound", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.path, nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}

func TestStaticFileServing(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	linkService := services.NewLinkService(cfg)
	router := internal.SetupRouter(linkService)

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
