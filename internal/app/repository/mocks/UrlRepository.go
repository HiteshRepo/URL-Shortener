// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "github.com/hiteshpattanayak-tw/url_shortner/internal/app/types"
)

// UrlRepository is an autogenerated mock type for the UrlRepository type
type UrlRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: shortUrl, longUrl
func (_m *UrlRepository) Add(shortUrl types.ShortUrl, longUrl types.LongUrl) {
	_m.Called(shortUrl, longUrl)
}

// GetAll provides a mock function with given fields:
func (_m *UrlRepository) GetAll() (map[types.ShortUrl]types.LongUrl, map[types.LongUrl]types.ShortUrl) {
	ret := _m.Called()

	var r0 map[types.ShortUrl]types.LongUrl
	if rf, ok := ret.Get(0).(func() map[types.ShortUrl]types.LongUrl); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.ShortUrl]types.LongUrl)
		}
	}

	var r1 map[types.LongUrl]types.ShortUrl
	if rf, ok := ret.Get(1).(func() map[types.LongUrl]types.ShortUrl); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[types.LongUrl]types.ShortUrl)
		}
	}

	return r0, r1
}

// GetLongUrl provides a mock function with given fields: shortUrl
func (_m *UrlRepository) GetLongUrl(shortUrl types.ShortUrl) types.LongUrl {
	ret := _m.Called(shortUrl)

	var r0 types.LongUrl
	if rf, ok := ret.Get(0).(func(types.ShortUrl) types.LongUrl); ok {
		r0 = rf(shortUrl)
	} else {
		r0 = ret.Get(0).(types.LongUrl)
	}

	return r0
}

// GetShortUrlIfExists provides a mock function with given fields: longUrl
func (_m *UrlRepository) GetShortUrlIfExists(longUrl types.LongUrl) types.ShortUrl {
	ret := _m.Called(longUrl)

	var r0 types.ShortUrl
	if rf, ok := ret.Get(0).(func(types.LongUrl) types.ShortUrl); ok {
		r0 = rf(longUrl)
	} else {
		r0 = ret.Get(0).(types.ShortUrl)
	}

	return r0
}

// Remove provides a mock function with given fields: shortUrl
func (_m *UrlRepository) Remove(shortUrl types.ShortUrl) {
	_m.Called(shortUrl)
}

// NewUrlRepository creates a new instance of UrlRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewUrlRepository(t testing.TB) *UrlRepository {
	mock := &UrlRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
