package endpoint

import "github.com/seagullbird/headr-contentmgr/db"

type NewPostRequest struct {
	db.Post
}

type NewPostResponse struct {
	Id  uint  `json:"id"`
	Err error `json:"-"`
}
