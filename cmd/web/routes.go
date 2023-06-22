package main

import (
	"net/http"

	"github.com/Denis-Andrei/goapp/pkg/config"
	"github.com/Denis-Andrei/goapp/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.Appconfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}