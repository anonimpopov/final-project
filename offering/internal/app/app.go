package app

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	server *http.Server
}

func initServer() http.Handler {
	serverMux := http.NewServeMux()

	//TODO handlers

	return serverMux
}

func NewApp(address string) *App {
	newServer := &http.Server{
		Addr:    address,
		Handler: initServer(),
	}

	return &App{
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
