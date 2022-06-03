package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
	"net/http"
)

const (
	ErrorResponseFieldKey = "error"
)

type UrlShortenerHandler struct {
	urlSvc service.UrlService
}

func ProvideNewUrlShortenerHandler(urlSvc service.UrlService) *UrlShortenerHandler {
	return &UrlShortenerHandler{urlSvc: urlSvc}
}

func (ush *UrlShortenerHandler) UrlShortenerHandler(w http.ResponseWriter, r *http.Request) {
	var urlData map[string]string
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&urlData); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if r.Method == "POST" {
		shortUrl := ush.urlSvc.ShortenUrl(types.LongUrl(urlData["long_url"]))
		respondWithJSON(w, http.StatusOK, map[string]string{"shortened_url": string(shortUrl)})
	} else {
		respondWithError(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s method not supported", r.Method))
	}
}

func (ush *UrlShortenerHandler) FetchOriginalUrlHandler(w http.ResponseWriter, r *http.Request) {
	var urlData map[string]string
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&urlData); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if r.Method == "GET" {
		longUrl := ush.urlSvc.GetOriginalUrl(types.ShortUrl(urlData["short_url"]))
		respondWithJSON(w, http.StatusOK, map[string]string{"original_url": string(longUrl)})
	} else {
		respondWithError(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s method not supported", r.Method))
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{ErrorResponseFieldKey: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}