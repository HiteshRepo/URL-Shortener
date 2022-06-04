//go:build wireinject
// +build wireinject

//go:generate wire

package di

import (
	"context"
	"github.com/google/wire"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/handlers"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/router"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/pkg/configs"
)

func InitializeApp(ctx context.Context) (*app.App, error) {
	wire.Build(
		repository.ProvideNewUrlRepository,
		handlers.ProvideNewUrlShortenerHandler,
		service.ProvideUrlService,
		router.ProvideRouter,
		configs.ProvideAppConfig,

		wire.Struct(new(app.App), "*"),
	)
	return &app.App{}, nil
}
