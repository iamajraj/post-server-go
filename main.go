package main

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	// create post
	r.POST("/posts", func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)

		if err != nil {
			SendError(500, "Please provide the correct data", ctx)
			return
		}

		var post Post
		if err := json.Unmarshal(body, &post); err != nil {
			SendError(500, "Failed to parse the body", ctx)
			return
		}

		validate := validator.New()
		if err := validate.Struct(post); err != nil {
			SendError(400, "Please provide the required field", ctx)
			return
		}

		db.Create(&post)
		ctx.JSON(201, gin.H{
			"message": "Successfully post has been created",
		})
	})

	log.Default().Println("Starting the server :8000")
	if err := r.Run("localhost:8000"); err != nil {
		log.Fatalf("Error running the server: %s", err.Error())
	}

}

func SendError(code int, message string, ctx *gin.Context) {
	ctx.JSON(code, gin.H{
		"message": message,
	})
}
