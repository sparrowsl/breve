package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	router := chi.NewMux()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	router.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	}))

	router.Get("/links", app.getLinks)
	router.Get("/links/{id}", app.getLink)
	router.Delete("/links/{id}", app.deleteLink)

	return router
}
