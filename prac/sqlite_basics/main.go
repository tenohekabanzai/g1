package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Slug  string `gorm:"uniqueIndex:idx_slug"`
	Likes uint
}

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func main() {
	fmt.Println("SQLITE with GORM")
	db.AutoMigrate(&Post{})

	// np := createPost("New Post Title", "new-slug")
	fp := getPost("new-slug")
	fmt.Println(fp)
}

func createPost(title string, slug string) Post {
	newPost := Post{Title: title, Slug: slug}
	if res := db.Create(&newPost); res.Error != nil {
		panic(res.Error)
	}
	fmt.Println("Post added to DB")
	return newPost
}

func getPost(slug string) Post {
	targetPost := Post{Slug: slug}
	if res := db.First(&targetPost); res.Error != nil {
		panic(res.Error)
	}

	fmt.Println("Post fetched from DB")
	return targetPost
}
