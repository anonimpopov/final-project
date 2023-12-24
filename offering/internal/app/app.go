package app

import (
	"context"
	"fmt"
	"net/http"
	"offering/internal/config"
	"offering/internal/handlers"
	"offering/internal/manager"
)

type App struct {
	cfg    *config.Config
	server *http.Server
}

func initServer(cfg *config.Config) http.Handler {
	serverMux := http.NewServeMux()

	handler := handlers.NewController(manager.NewManager(cfg))

	serverMux.HandleFunc("/parseOffer", handler.ParseOffer)
	serverMux.HandleFunc("/createOffer", handler.CreateOffer)

	return serverMux
}

func NewApp(cfg *config.Config) *App {
	newServer := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: initServer(cfg),
	}

	return &App{
		cfg:    cfg,
		server: newServer,
	}
}

func (a *App) Run() {
	go func() {
		err := a.server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()
}

func (a *App) Stop(ctx context.Context) {
	fmt.Println(a.server.Shutdown(ctx))
}
