package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sparrowsl/breve/config"
	"github.com/sparrowsl/breve/database/sqlc"
)

func main() {
	ctx := context.Background()
	config := (&config.Config{}).Load()

	db, err := sql.Open("mysql", config.DATABASE_URL)
	if err != nil {
		fmt.Fprintln(os.Stdin, err)
		os.Exit(1)
	}
	defer db.Close()

	queries := sqlc.New(db)

	breves, err := queries.GetBreves(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(breves)
	fmt.Println(config.DATABASE_URL)
}
