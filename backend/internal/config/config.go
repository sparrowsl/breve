package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stdin, err)
		os.Exit(1)
	}

	// do extra check to make sure there's is value in the env file
}
