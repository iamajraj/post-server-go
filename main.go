package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint      `gorm:"primarykey, autoIncrement" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `gorm:"not null" json:"title"`
	Published bool      `gorm:"not null,default:false" json:"published"`
	Author    string    `gorm:"not null" json:"author"`
}

func main() {

	dsn := "host=localhost user=postgres password=root dbname=gormexplore port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&Post{})
	// db.Migrator().DropTable(&Post{})
	// db.Migrator().CreateTable(&Post{})

	// post := Post{
	// 	Title:  "A new post",
	// 	Author: "raajz",
	// }

	// db.Create(&post)

	var posts []Post
	db.Find(&posts)

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Published: %t\n", post.Published)
		fmt.Printf("Author: %s\n", post.Author)
	}
}
