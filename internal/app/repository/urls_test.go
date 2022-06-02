package repository_test

import (
	"context"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type urlRepositorySuite struct {
	suite.Suite
	ctx     context.Context
	urlRepo *repository.UrlRepository
}

func TestUrlRepository(t *testing.T) {
	suite.Run(t, new(urlRepositorySuite))
}

func (suite *urlRepositorySuite) SetupTest() {
	suite.ctx = context.Background()
	suite.urlRepo = repository.ProvideNewUrlRepository()
}

func (suite *urlRepositorySuite) TestUrlRepository_ShouldAddUrlToRepo() {
	shortUrl := types.ShortUrl("domain://test-url.com")
	longUrl := types.LongUrl("https://test-12345678-test.com")

	suite.urlRepo.Add(shortUrl, longUrl)

	expected := map[types.ShortUrl]types.LongUrl{shortUrl: longUrl}
	reverseExpected := map[types.LongUrl]types.ShortUrl{longUrl: shortUrl}
	actual, reverseActual := suite.urlRepo.Get()

	suite.Assert().Equal(expected, actual)
	suite.Assert().Equal(reverseExpected, reverseActual)
}

func (suite *urlRepositorySuite) TestUrlRepository_ShouldRemoveUrlFromRepo() {
	shortUrl1 := types.ShortUrl("domain://test-url.com")
	longUrl1 := types.LongUrl("https://test-12345678-test.com")

	shortUrl2 := types.ShortUrl("domain://test-url-2.com")
	longUrl2 := types.LongUrl("https://test-123456789-test.com")

	suite.urlRepo.Add(shortUrl1, longUrl1)
	suite.urlRepo.Add(shortUrl2, longUrl2)

	expected := map[types.ShortUrl]types.LongUrl{shortUrl2: longUrl2}
	reverseExpected := map[types.LongUrl]types.ShortUrl{longUrl2: shortUrl2}

	suite.urlRepo.Remove(shortUrl1)

	actual, reverseActual := suite.urlRepo.Get()

	suite.Assert().Equal(expected, actual)
	suite.Assert().Equal(reverseExpected, reverseActual)
}

func (suite *urlRepositorySuite) TestUrlRepository_ShouldGetLongUrlByShortUrl() {
	shortUrl := types.ShortUrl("domain://test-url.com")
	longUrl := types.LongUrl("https://test-12345678-test.com")

	suite.urlRepo.Add(shortUrl, longUrl)

	expected := longUrl
	actual := suite.urlRepo.GetLongUrl(shortUrl)

	suite.Assert().Equal(expected, actual)
}
