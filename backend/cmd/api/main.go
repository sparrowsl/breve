package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "breve/internal/config"
	"breve/internal/database"

	_ "github.com/lib/pq"
)

type application struct {
	db    *database.Queries
	links *database.Link
	ctx   context.Context
}

func main() {
	conn, err := openDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}
	defer conn.Close()

	queries := database.New(conn)
	app := &application{
		db: queries,
		ctx: context.Background(),
	}

	server := &http.Server{
		Addr:    ":5000",
		Handler: app.routes(),
	}

	fmt.Printf("Starting server on :%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
