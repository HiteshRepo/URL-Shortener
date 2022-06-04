package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/pkg/configs"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router    *mux.Router
	AppConfig *configs.AppConfig
}

func (a *App) Start() {
	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", a.AppConfig.ServerConfig.Host, a.AppConfig.ServerConfig.Port),
		Handler:      a.Router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server")
	log.Fatal(srv.ListenAndServe())
}

func (a *App) Shutdown() {}
