package endpoint

type NewPostRequest struct {
}

type NewPostResponse struct {
	Err error `json:"-"`
}
