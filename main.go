package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint      `gorm:"primarykey, autoIncrement" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `gorm:"not null" json:"title" validate:"required"`
	Published bool      `gorm:"not null,default:false" json:"published"`
	Author    string    `gorm:"not null" json:"author" validate:"required"`
}

func main() {

	dsn := "host=localhost user=postgres password=root dbname=gormexplore port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&Post{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api": "v1",
		})
	})

	// get all the posts
	r.GET("/posts", func(c *gin.Context) {

		var posts []Post
		db.Find(&posts)

		c.JSON(200, gin.H{
			"posts": posts,
		})
	})

	log.Default().Println("Starting the server :8000")
	if err := r.Run("localhost:8000"); err != nil {
		log.Fatalf("Error running the server: %s", err.Error())
	}
	// db.Migrator().DropTable(&Post{})
	// db.Migrator().CreateTable(&Post{})

	// post := Post{
	// 	Title:  "A new post",
	// 	Author: "raajz",
	// }

	// db.Create(&post)

	// updating the post published
	// var firstPost Post
	// db.First(&firstPost).Update("Published", true)
	// validate := validator.New()

}
