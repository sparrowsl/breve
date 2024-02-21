package main

import (
	"fmt"
	"os"

	"github.com/sparrowsl/breve/config"
	"github.com/sparrowsl/breve/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := (&config.Config{}).Load()

	fmt.Println(config.DATABASE_URL)
	db, err := gorm.Open(mysql.Open(config.DATABASE_URL), &gorm.Config{})
	if err != nil {
		fmt.Fprintln(os.Stdin, err)
		os.Exit(1)
	}

	db.AutoMigrate(models.Breve{})

	fmt.Println(db)
}
