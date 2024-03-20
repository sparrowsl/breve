package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	_ "breve/internal/config"
	"breve/internal/database"

	"github.com/jackc/pgx/v5"
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
	defer conn.Close(context.Background())

	// queries := database.New(conn)
	app := &application{
		// db:  queries,
		ctx: context.Background(),
	}

	server := &http.Server{
		Addr:    ":5000",
		Handler: app.routes(),
	}

	fmt.Println("Starting server on :5000")
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
