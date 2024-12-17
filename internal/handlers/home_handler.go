package handlers

import (
	"html/template"
	"net/http"

	"github.com/philgear/philtree/internal/services"
)

type PageData struct {
	Websites       []services.Link
	SocialMedia    []services.Link
	AffiliateLinks []services.Link
}

func HomeHandler(ls *services.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Websites:       ls.GetWebsites(),
			SocialMedia:    ls.GetSocialMediaLinks(),
			AffiliateLinks: ls.GetAffiliateLinks(),
		}

		tmpl, err := template.ParseFiles("web/templates/home.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}