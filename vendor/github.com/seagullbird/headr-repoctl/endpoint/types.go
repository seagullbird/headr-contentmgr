package endpoint

type NewSiteRequest struct {
	Email    string
	SiteName string
}

type NewSiteResponse struct {
	Err error `json:"-"`
}

type DeleteSiteRequest struct {
	Email    string
	SiteName string
}

type DeleteSiteResponse struct {
	Err error `json:"-"`
}

type WritePostRequest struct {
	Author   string
	Sitename string
	Filename string
	Content  string
}

type WritePostResponse struct {
	Err error `json:"-"`
}

type RemovePostRequest struct {
	Author   string
	Sitename string
	Filename string
}

type RemovePostResponse struct {
	Err error `json:"-"`
}

type ReadPostRequest struct {
	Author   string
	Sitename string
	Filename string
}

type ReadPostResponse struct {
	Content string `json:"content"`
	Err     error  `json:"-"`
}
