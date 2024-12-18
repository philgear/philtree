// Package handlers provides HTTP request handlers for the application
package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/philgear/philtree/internal/services"
)

// HomeHandler returns an HTTP handler for the home page
func HomeHandler(ls *services.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Websites       []services.Website
			SocialMedia    []services.SocialMedia
			AffiliateLinks []services.AffiliateLink
			LastUpdated    string
		}{
			Websites:       ls.Websites,
			SocialMedia:    ls.SocialMedia,
			AffiliateLinks: ls.AffiliateLinks,
			LastUpdated:    time.Now().Format("2006-01-02"),
		}

		tmpl, err := template.ParseFiles("web/templates/home.html")
		if err != nil {
			http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
