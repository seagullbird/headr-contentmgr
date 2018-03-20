package db

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"os"
)

func Example() {
	dbHost := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	dbPort := os.Getenv("POSTGRES_PORT_5432_TCP_PORT")
	dbPassword := os.Getenv("POSTGRES_ENV_POSTGRES_PASSWORD")
	args := fmt.Sprintf("host=%s port=%s user=postgres dbname=postgres password=%s sslmode=disable", dbHost, dbPort, dbPassword)
	dbConn, err := gorm.Open("postgres", args)
	checkError(err)
	store := New(dbConn)

	// prepare data
	postA := &Post{
		SiteID:   1,
		UserID:   "user_id_1",
		Filename: "postA",
		Filetype: "md",
		Title:    "Post A",
		Date:     "today",
		Draft:    false,
		Tags:     "tag1 tag2",
		Summary:  "summary A",
		Content:  "content A",
	}

	postB := &Post{
		SiteID:   1,
		UserID:   "user_id_1",
		Filename: "postB",
		Filetype: "md",
		Title:    "Post B",
		Date:     "today",
		Draft:    false,
		Tags:     "tag1 tag2",
		Summary:  "summary B",
		Content:  "content B",
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

	// summary A
	// <!--more-->

	// [{
	//   "title": "Post B",
	//   "date": "today",
	//   "draft": false,
	//   "tags": [
	//     "tag1",
	//     "tag2"
	//   ]
	// }

	// summary B
	// <!--more-->
	//  {
	//   "title": "Post A",
	//   "date": "today",
	//   "draft": false,
	//   "tags": [
	//     "tag1",
	//     "tag2"
	//   ]
	// }

	// summary A
	// <!--more-->
	// ]
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
