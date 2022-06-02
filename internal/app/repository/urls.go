package repository

type UrlRepository struct {
	urlMap map[string]string
}

func ProvideNewUrlRepository() *UrlRepository {
	return &UrlRepository{urlMap: make(map[string]string)}
}
