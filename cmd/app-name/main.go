package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philgear/philtree/internal/config"
	"github.com/philgear/philtree/internal/handlers"
	"github.com/philgear/philtree/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	linkService := services.NewLinkService(cfg)

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler(linkService)).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/assets"))))

	log.Printf("Starting server on :%s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}