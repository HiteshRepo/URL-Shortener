package repository

import "github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"

type UrlRepository struct {
	urlMap map[types.ShortUrl]types.LongUrl
	reverseUrlMap map[types.LongUrl]types.ShortUrl
}

func ProvideNewUrlRepository() *UrlRepository {
	return &UrlRepository{urlMap: make(map[types.ShortUrl]types.LongUrl), reverseUrlMap: make(map[types.LongUrl]types.ShortUrl)}
}

func (ur *UrlRepository) Add(shortUrl types.ShortUrl, longUrl types.LongUrl) {
	ur.urlMap[shortUrl] = longUrl
	ur.reverseUrlMap[longUrl] = shortUrl
}

func (ur *UrlRepository) Remove(shortUrl types.ShortUrl) {
	var lUrl types.LongUrl
	var ok bool
	if lUrl, ok = ur.urlMap[shortUrl]; ok {
		delete(ur.urlMap, shortUrl)
	}

	if _, ok = ur.reverseUrlMap[lUrl]; ok {
		delete(ur.reverseUrlMap, lUrl)
	}
}

func (ur *UrlRepository) GetLongUrl(shortUrl types.ShortUrl) types.LongUrl {
	return ur.urlMap[shortUrl]
}

func (ur *UrlRepository) Get() (map[types.ShortUrl]types.LongUrl, map[types.LongUrl]types.ShortUrl) {
	return ur.urlMap, ur.reverseUrlMap
}