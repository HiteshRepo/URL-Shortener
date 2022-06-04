package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/handlers"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/router"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/pkg/configs"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

type UrlShortenerSuite struct {
	suite.Suite
	ctx     context.Context
	rootDir string
	router  *mux.Router
	config  *configs.AppConfig
}

func TestUrlShortener(t *testing.T) {
	suite.Run(t, new(UrlShortenerSuite))
}

func (suite *UrlShortenerSuite) SetupTest() {
	suite.ctx = context.Background()

	workingDirectory, _ := os.Getwd()
	pathToConfig := path.Join(workingDirectory, "../../configs/local.yaml")

	configReader, err := os.Open(pathToConfig)
	suite.Require().NoError(err)

	appConfig, err := configs.LoadConfig(configReader)
	suite.Require().NoError(err)

	urlRepository := repository.ProvideNewUrlRepository()
	urlService := service.ProvideUrlService(urlRepository)
	urlShortenerHandler := handlers.ProvideNewUrlShortenerHandler(urlService)
	muxRouter := router.ProvideRouter(urlShortenerHandler)

	suite.router = muxRouter
	suite.config = appConfig
}

func (suite *UrlShortenerSuite) TestShortenEndpoint() {
	expectedBody, err := ioutil.ReadFile("../../test/response/shorten_url_response.json")
	suite.Require().NoError(err)

	var jsonStr = []byte(`{"long_url":"https://1very_log_url_to_be_shortened12"}`)
	response := suite.makeRequest("POST", "/shorten_url", bytes.NewBuffer(jsonStr))
	suite.Require().Equal(http.StatusOK, response.Code)

	body, err := io.ReadAll(response.Body)
	suite.Require().NoError(err)

	var actualResponse map[string]string
	err = json.Unmarshal(body, &actualResponse)
	suite.Require().NoError(err)

	var expectedResponse map[string]string
	err = json.Unmarshal(expectedBody, &expectedResponse)
	suite.Require().NoError(err)

	suite.Require().EqualValues(expectedResponse, actualResponse)
}

func (suite *UrlShortenerSuite) TestFetchUrlEndpoint() {
	expectedBody, err := ioutil.ReadFile("../../test/response/fetch_url_response.json")
	suite.Require().NoError(err)

	var jsonStr = []byte(`{"long_url":"https://1very_log_url_to_be_shortened12"}`)
	response := suite.makeRequest("POST", "/shorten_url", bytes.NewBuffer(jsonStr))
	suite.Require().Equal(http.StatusOK, response.Code)

	body, err := io.ReadAll(response.Body)
	suite.Require().NoError(err)

	var shortenUrlResponse map[string]string
	err = json.Unmarshal(body, &shortenUrlResponse)
	suite.Require().NoError(err)

	var expectedResponse map[string]string
	err = json.Unmarshal(expectedBody, &expectedResponse)
	suite.Require().NoError(err)

	jsonStr = []byte(fmt.Sprintf(`{"short_url":"%s"}`, shortenUrlResponse["shortened_url"]))
	response = suite.makeRequest("GET", "/fetch_url", bytes.NewBuffer(jsonStr))
	suite.Require().Equal(http.StatusOK, response.Code)

	body, err = io.ReadAll(response.Body)
	suite.Require().NoError(err)

	var actualResponse map[string]string
	err = json.Unmarshal(body, &actualResponse)
	suite.Require().NoError(err)

	suite.Require().EqualValues(expectedResponse, actualResponse)
}

func (suite *UrlShortenerSuite) makeRequest(method, url string, body *bytes.Buffer) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)

	suite.Require().NoError(err)
	recorder := httptest.NewRecorder()

	suite.router.ServeHTTP(recorder, req)
	return recorder
}

