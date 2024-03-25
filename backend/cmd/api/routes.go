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
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	}))

	router.Get("/links/r/{redirect}", app.redirectTo)
	router.Get("/links", app.getLinks)
	router.Get("/links/{id}", app.getLink)
	router.Post("/links", app.createLink)
	router.Patch("/links/{id}", app.updateLink)
	router.Delete("/links/{id}", app.deleteLink)

	return router
}
