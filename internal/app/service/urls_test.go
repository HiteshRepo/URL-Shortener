package service_test

import (
	"context"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository/mocks"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service"
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