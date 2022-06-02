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

	expected := map[types.ShortUrl]types.LongUrl{"domain://test-url.com": "https://test-12345678-test.com"}
	actual := suite.urlRepo.Get()

	suite.Assert().Equal(expected, actual)
}
