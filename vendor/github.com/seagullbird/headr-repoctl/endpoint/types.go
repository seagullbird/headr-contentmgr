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

type NewPostRequest struct {
	Author   string
	Sitename string
	Filename string
	Content  string
}

type NewPostResponse struct {
	Err error `json:"-"`
}
