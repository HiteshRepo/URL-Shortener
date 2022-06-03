package service

import (
	"fmt"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
	"math/rand"
)

const (
	Domain = "https://bitly.com"
	ShortenLength = 7
	Base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
	if sUrl := us.urlRepo.GetShortUrlIfExists(longUrl); len(sUrl) > 0 {
		return sUrl
	}

	randomNum := us.rangeIn(100000000000, 999999999999)
	shortUrl := types.ShortUrl(fmt.Sprintf("%s/%s", Domain, us.base62Encode(randomNum)))

	done := false
	for !done {
		lUrl := us.urlRepo.GetLongUrl(shortUrl)
		if len(lUrl) == 0 {
			done = true
			continue
		}
		randomNum = us.rangeIn(100000000000, 999999999999)
		shortUrl = types.ShortUrl(fmt.Sprintf("%s/%s", Domain, us.base62Encode(randomNum)))
	}

	us.urlRepo.Add(shortUrl, longUrl)

	return shortUrl
}

func (us urlService) GetOriginalUrl(shortUrl types.ShortUrl) types.LongUrl {
	return ""
}

func (us urlService) base62Encode(randomNum int64) string {
	hashStr := ""

	for randomNum > 0 {
		hashStr = fmt.Sprintf("%s%c", hashStr, Base62[randomNum % 62])
		randomNum = randomNum / 62
	}

	return hashStr
}

func (us urlService) rangeIn(low, hi int64) int64 {
	return low + rand.Int63n(hi-low)
}