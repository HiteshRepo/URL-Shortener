package main

import (
	"fmt"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/repository"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/router"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/service"
	"log"
	"net/http"
	"time"
)

const (
	Host = "localhost"
	Port = 9091
)

func main() {
	urlRepo := repository.ProvideNewUrlRepository()
	urlSvc := service.ProvideUrlService(urlRepo)
	r := router.ProvideRouter(urlSvc)

	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", Host, Port),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server")
	log.Fatal(srv.ListenAndServe())

}
