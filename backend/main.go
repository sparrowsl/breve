package main

import (
	"fmt"

	"github.com/sparrowsl/breve/config"
)

func main() {
	config := (&config.Config{}).Load()

	fmt.Println(config.DATABASE_URL)
}
