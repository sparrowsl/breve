package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
}

func (c *Config) Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stdin, err)
		os.Exit(1)
	}

	// Initialize data from environment variables
	c.DATABASE_URL = os.Getenv("DATABASE_URL")

	return *c
}
