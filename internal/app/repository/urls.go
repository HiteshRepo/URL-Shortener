package repository

import "github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"

type UrlRepository struct {
	urlMap map[types.ShortUrl]types.LongUrl
}

func ProvideNewUrlRepository() *UrlRepository {
	return &UrlRepository{urlMap: make(map[types.ShortUrl]types.LongUrl)}
}

func (ur *UrlRepository) Add(shortUrl types.ShortUrl, longUrl types.LongUrl) {
	ur.urlMap[shortUrl] = longUrl
}

func (ur *UrlRepository) Get() map[types.ShortUrl]types.LongUrl {
	return ur.urlMap
}