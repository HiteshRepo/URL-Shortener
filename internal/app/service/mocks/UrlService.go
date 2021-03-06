// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
)

// UrlService is an autogenerated mock type for the UrlService type
type UrlService struct {
	mock.Mock
}

// GetOriginalUrl provides a mock function with given fields: shortUrl
func (_m *UrlService) GetOriginalUrl(shortUrl types.ShortUrl) types.LongUrl {
	ret := _m.Called(shortUrl)

	var r0 types.LongUrl
	if rf, ok := ret.Get(0).(func(types.ShortUrl) types.LongUrl); ok {
		r0 = rf(shortUrl)
	} else {
		r0 = ret.Get(0).(types.LongUrl)
	}

	return r0
}

// ShortenUrl provides a mock function with given fields: longUrl
func (_m *UrlService) ShortenUrl(longUrl types.LongUrl) types.ShortUrl {
	ret := _m.Called(longUrl)

	var r0 types.ShortUrl
	if rf, ok := ret.Get(0).(func(types.LongUrl) types.ShortUrl); ok {
		r0 = rf(longUrl)
	} else {
		r0 = ret.Get(0).(types.ShortUrl)
	}

	return r0
}

// NewUrlService creates a new instance of UrlService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewUrlService(t testing.TB) *UrlService {
	mock := &UrlService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
