package service

import (
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
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
	return ""
}

func (us urlService) GetOriginalUrl(shortUrl types.ShortUrl) types.LongUrl {
	return ""
}