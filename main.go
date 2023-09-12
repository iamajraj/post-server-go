package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID        int `gorm `
	Title     string
	Published bool
	Author    string
}

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=gormexplore port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&Post{})

}
