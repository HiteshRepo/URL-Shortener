package repository

import "github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"

type UrlRepository interface {
	Add(shortUrl types.ShortUrl, longUrl types.LongUrl)
	Remove(shortUrl types.ShortUrl)
	GetLongUrl(shortUrl types.ShortUrl) types.LongUrl
	GetShortUrlIfExists(longUrl types.LongUrl) types.ShortUrl
	GetAll() (map[types.ShortUrl]types.LongUrl, map[types.LongUrl]types.ShortUrl)
}

type urlRepository struct {
	urlMap map[types.ShortUrl]types.LongUrl
	reverseUrlMap map[types.LongUrl]types.ShortUrl
}

func ProvideNewUrlRepository() UrlRepository {
	return &urlRepository{urlMap: make(map[types.ShortUrl]types.LongUrl), reverseUrlMap: make(map[types.LongUrl]types.ShortUrl)}
}

func (ur *urlRepository) Add(shortUrl types.ShortUrl, longUrl types.LongUrl) {
	ur.urlMap[shortUrl] = longUrl
	ur.reverseUrlMap[longUrl] = shortUrl
}

func (ur *urlRepository) Remove(shortUrl types.ShortUrl) {
	var lUrl types.LongUrl
	var ok bool
	if lUrl, ok = ur.urlMap[shortUrl]; ok {
		delete(ur.urlMap, shortUrl)
	}

	if _, ok = ur.reverseUrlMap[lUrl]; ok {
		delete(ur.reverseUrlMap, lUrl)
	}
}

func (ur *urlRepository) GetLongUrl(shortUrl types.ShortUrl) types.LongUrl {
	return ur.urlMap[shortUrl]
}

func (ur *urlRepository) GetShortUrlIfExists(longUrl types.LongUrl) types.ShortUrl {
	if sUrl,ok := ur.reverseUrlMap[longUrl]; ok {
		return sUrl
	}
	return ""
}

func (ur *urlRepository) GetAll() (map[types.ShortUrl]types.LongUrl, map[types.LongUrl]types.ShortUrl) {
	return ur.urlMap, ur.reverseUrlMap
}