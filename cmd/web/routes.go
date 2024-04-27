package main

import (
	"net/http"

	"github.com/alg3n/bookings/pkg/config"
	"github.com/alg3n/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// routes sets up the routes for the web application.
// It creates a new chi router, adds middleware, and defines the routes.
// The routes are mapped to the corresponding handler functions from the handlers package.
// The configured router is returned as an http.Handler.
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
