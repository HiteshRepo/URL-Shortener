package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	Host = "localhost"
	Port = 9091
)

type App struct {
	Router  *mux.Router
}

func (a *App) Start() {
	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", Host, Port),
		Handler:      a.Router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server")
	log.Fatal(srv.ListenAndServe())
}

func (a *App) Shutdown() {}
