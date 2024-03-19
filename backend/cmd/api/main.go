package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "breve/internal/config"
	"breve/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	db    *sql.DB
	links *database.Link
}

func main() {
	db, err := openDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	app := &application{
		db: db,
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
