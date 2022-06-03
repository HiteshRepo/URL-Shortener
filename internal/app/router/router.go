package router

import (
	"github.com/gorilla/mux"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/handlers"
)

func ProvideRouter(urlShortenerHandler *handlers.UrlShortenerHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/shorten_url", urlShortenerHandler.UrlShortenerHandler).Methods("POST")
	r.HandleFunc("/fetch_url", urlShortenerHandler.FetchOriginalUrlHandler).Methods("GET")

	return r
}