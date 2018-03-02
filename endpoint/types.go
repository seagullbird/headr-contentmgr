package endpoint

import "github.com/seagullbird/headr-contentmgr/service"

type NewPostRequest struct {
	service.Post
}

type NewPostResponse struct {
	Id  string `json:"id"`
	Err error  `json:"-"`
}
