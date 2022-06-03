package router

import (
	"github.com/gorilla/mux"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/handlers"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service"
)

func ProvideRouter(urlSvc service.UrlService) *mux.Router {
	r := mux.NewRouter()

	urlShortenerHandler := handlers.ProvideNewUrlShortenerHandler(urlSvc)

	r.HandleFunc("/shorten_url", urlShortenerHandler.UrlShortenerHandler).Methods("POST")

	return r
}