package service

import (
	"fmt"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
)

const (
	Domain = "https://bitly.com"
	ShortenLength = 7
)

type UrlService interface {
	ShortenUrl(longUrl types.LongUrl) types.ShortUrl
	GetOriginalUrl(shortUrl types.ShortUrl) types.LongUrl
}

type urlService struct {
	urlRepo repository.UrlRepository
}

func ProvideUrlService(urlRepo repository.UrlRepository) UrlService {
	return urlService{urlRepo: urlRepo}
}

func (us urlService) ShortenUrl(longUrl types.LongUrl) types.ShortUrl {
	shortUrl := fmt.Sprintf("%s/%s", Domain, string(longUrl)[0:ShortenLength])
	return types.ShortUrl(shortUrl)
}

func (us urlService) GetOriginalUrl(shortUrl types.ShortUrl) types.LongUrl {
	return ""
}