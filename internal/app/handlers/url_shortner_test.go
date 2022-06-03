package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/handlers"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service/mocks"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUrlShortenerHandler_UrlShortenerHandler(t *testing.T) {
	var jsonStr = []byte(`{"long_url":"https://a_very_log_url_to_be_shortened1.com"}`)
	req, _ := http.NewRequest("POST", "/shorten_url", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	longUrl := types.LongUrl("https://a_very_log_url_to_be_shortened1.com")
	urlSvc := new(mocks.UrlService)
	urlSvc.On("ShortenUrl", longUrl).Return(types.ShortUrl("https://bitly.com/aBcDeFg"))

	ush := handlers.ProvideNewUrlShortenerHandler(urlSvc)
	ush.UrlShortenerHandler(w, req)

	expectedResp := map[string]string{"shortened_url": "https://bitly.com/aBcDeFg"}
	b, err := json.Marshal(expectedResp)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(b), string(w.Body.Bytes()))
}