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

var counter int64 = 100000000000

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
	fmt.Println(longUrl)
	fmt.Println(counter)
	if sUrl := us.urlRepo.GetShortUrlIfExists(longUrl); len(sUrl) > 0 {
		fmt.Println(longUrl)
		fmt.Println(sUrl)
		return sUrl
	}

	fmt.Println("generating")
	randomNum := us.rangeIn(counter, 999999999999)
	shortUrl := types.ShortUrl(fmt.Sprintf("%s/%s", Domain, us.base62Encode(randomNum)))

	fmt.Println(shortUrl)
	counter += 1

	us.urlRepo.Add(shortUrl, longUrl)

	return shortUrl
}

func (us urlService) GetOriginalUrl(shortUrl types.ShortUrl) types.LongUrl {
	return us.urlRepo.GetLongUrl(shortUrl)
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