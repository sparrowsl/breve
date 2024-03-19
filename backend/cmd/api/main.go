package main

import (
	"fmt"
	"net/http"

	_ "breve/internal/config"
	"breve/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	db *database.Link
}

func main() {
	app := &application{
		db: &database.Link{},
	}

	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println(app)
	_ = server.ListenAndServe()
}
