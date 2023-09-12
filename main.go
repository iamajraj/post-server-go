package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primarykey, autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Published bool
	Author    string
}

func main() {

	dsn := "host=localhost user=postgres password=root dbname=gormexplore port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&Post{})

}
