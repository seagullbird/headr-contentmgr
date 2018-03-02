package endpoint

import "github.com/seagullbird/headr-contentmgr/service"

type NewPostRequest struct {
	service.Post
}

type NewPostResponse struct {
	Id  uint  `json:"id"`
	Err error `json:"-"`
}
