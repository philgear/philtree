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

	r := SetupRouter(linkService)

	log.Printf("Starting server on :%s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func SetupRouter(ls *services.LinkService) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler(ls)).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/assets"))))

	return r
}