package db_test

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/seagullbird/headr-contentmgr/db"
	"os"
)

func Example() {
	dbHost := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	dbPort := os.Getenv("POSTGRES_PORT_5432_TCP_PORT")
	dbPassword := os.Getenv("POSTGRES_ENV_POSTGRES_PASSWORD")
	args := fmt.Sprintf("host=%s port=%s user=postgres dbname=postgres password=%s sslmode=disable", dbHost, dbPort, dbPassword)
	dbConn, err := gorm.Open("postgres", args)
	checkError(err)
	store := db.New(dbConn)

	// prepare data
	postA := &db.Post{
		SiteID:   1,
		UserID:   "user_id_1",
		Filename: "postA",
		Filetype: "md",
		Title:    "Post A",
		Date:     "today",
		Draft:    false,
		Tags:     "tag1 tag2",
		Summary:  "summary A",
	}

	postB := &db.Post{
		SiteID:   1,
		UserID:   "user_id_1",
		Filename: "postB",
		Filetype: "md",
		Title:    "Post B",
		Date:     "today",
		Draft:    false,
		Tags:     "tag1 tag2",
		Summary:  "summary B",
	}

	modifiedPostB := &db.Post{
		Title:   "Modified Post B",
		Date:    "yesterday",
		Draft:   true,
		Tags:    "tag3 tag4",
		Summary: "modified summary B",
	}

	// insert two posts
	postAID, err := store.InsertPost(postA)
	checkError(err)
	postBID, err := store.InsertPost(postB)
	checkError(err)
	fmt.Printf("postA id=%d\n", postAID)
	fmt.Printf("postB id=%d\n", postBID)

	// get postA
	post, err := store.GetPost(postAID)
	checkError(err)
	fmt.Println(post)

	// get two posts
	posts, err := store.GetAllPosts("user_id_1")
	checkError(err)
	fmt.Println(posts)

	// delete postA
	err = store.DeletePost(postA)
	checkError(err)
	// get postA again to check delete
	post, err = store.GetPost(postAID)
	if err != gorm.ErrRecordNotFound {
		panic(errors.New("delete failed"))
	}
	// patch postB
	modifiedPostB.ID = postB.ID
	_, err = store.PatchPost(postB, modifiedPostB)
	checkError(err)
	// Get new postB
	newPostB, err := store.GetPost(postB.ID)
	checkError(err)
	// check if modification succeeded
	if newPostB.Title == modifiedPostB.Title &&
		newPostB.Tags == modifiedPostB.Tags &&
		newPostB.Summary == modifiedPostB.Summary &&
		newPostB.Date == modifiedPostB.Date &&
		newPostB.Draft == modifiedPostB.Draft {
		fmt.Println("Modification succeeded")
	} else {
		checkError(errors.New("modification failed"))
	}

	// Output:
	// postA id=1
	// postB id=2
	// {
	//   "title": "Post A",
	//   "date": "today",
	//   "draft": false,
	//   "tags": [
	//     "tag1",
	//     "tag2"
	//   ]
	// }
	//
	// summary A
	// <!--more-->
	//
	// [2 1]
	// Modification succeeded
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
