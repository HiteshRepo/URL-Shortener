package service_test

import (
	"context"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository/mocks"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type urlServiceSuite struct {
	suite.Suite
	ctx     context.Context
	urlRepo *mocks.UrlRepository
	urlSvc  service.UrlService
}

func TestUrlService(t *testing.T) {
	suite.Run(t, new(urlServiceSuite))
}

func (suite *urlServiceSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.urlRepo = new(mocks.UrlRepository)
	suite.urlSvc = service.ProvideUrlService(suite.urlRepo)
}

func (suite *urlServiceSuite) TestUrlService_ShortenUrlShouldGiveShortenedUrlOfLength7ForAGivenLongUrl() {
	longUrl := types.LongUrl("https://a_very_log_url_to_be_shortened.com")
	shortUrl := suite.urlSvc.ShortenUrl(longUrl)

	suite.Assert().True(len(shortUrl) == service.ShortenLength + len(service.Domain) + 1)
}

func (suite *urlServiceSuite) TestUrlService_ShortenUrlShouldGiveDifferentShortenedUrlsOfLength7ForAGivenDifferentLongUrls() {
	longUrl1 := types.LongUrl("https://a_very_log_url_to_be_shortened1.com")
	longUrl2 := types.LongUrl("https://a_very_log_url_to_be_shortened2.com")
	shortUrl1 := suite.urlSvc.ShortenUrl(longUrl1)
	shortUrl2 := suite.urlSvc.ShortenUrl(longUrl2)

	suite.Assert().NotEqual(shortUrl1, shortUrl2)
}