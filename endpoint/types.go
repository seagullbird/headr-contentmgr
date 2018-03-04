package endpoint

import "github.com/seagullbird/headr-contentmgr/db"

type NewPostRequest struct {
	db.Post
}

type NewPostResponse struct {
	Id  uint  `json:"id"`
	Err error `json:"-"`
}

type DeletePostRequest struct {
	Id uint `json:"id"`
}

type DeletePostResponse struct {
	Err error `json:"-"`
}

type GetPostRequest struct {
	Id uint `json:"id"`
}

type GetPostResponse struct {
	db.Post
	Err error `json:"-"`
}
