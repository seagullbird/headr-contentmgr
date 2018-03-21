package endpoint

import "github.com/seagullbird/headr-contentmgr/db"

// NewPostRequest collects the request parameters for the NewPost method.
type NewPostRequest struct {
	db.Post
}

// NewPostResponse collects the response values for the NewPost method.
type NewPostResponse struct {
	ID  uint  `json:"id"`
	Err error `json:"-"`
}

// DeletePostRequest collects the request parameters for the DeletePost method.
type DeletePostRequest struct {
	ID uint `json:"id"`
}

// DeletePostResponse collects the response values for the DeletePost method.
type DeletePostResponse struct {
	Err error `json:"-"`
}

// GetPostRequest collects the request parameters for the GetPost method.
type GetPostRequest struct {
	ID uint `json:"id"`
}

// GetPostResponse collects the response values for the GetPost method.
type GetPostResponse struct {
	db.Post
	Err error `json:"-"`
}

// GetAllPostsRequest collects the request parameters for the GetAllPosts method.
type GetAllPostsRequest struct {
}

// GetAllPostsResponse collects the response values for the GetAllPosts method.
type GetAllPostsResponse struct {
	PostIDs []uint `json:"post_ids"`
	Err     error  `json:"-"`
}

// PatchPostRequest collects the request parameters for the PatchPost method.
type PatchPostRequest struct {
	db.Post
}

// PatchPostResponse collects the response values for the PatchPost method.
type PatchPostResponse struct {
	Err error `json:"-"`
}
