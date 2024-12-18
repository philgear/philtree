package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philgear/philtree/internal/handlers"
	"github.com/philgear/philtree/internal/services"
)

// SetupRouter configures and returns the application router
func SetupRouter(ls *services.LinkService) *mux.Router {
	r := mux.NewRouter()

	// Set up routes
	r.HandleFunc("/", handlers.HomeHandler(ls)).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/assets"))))

	return r
}
