package endpoint

// NewSiteRequest collects the request parameters for the NewSite method.
type NewSiteRequest struct {
	Email    string
	SiteName string
}

// NewSiteResponse collects the response values for the NewSite method.
type NewSiteResponse struct {
	Err error `json:"-"`
}

// DeleteSiteRequest collects the request parameters for the DeleteSite method.
type DeleteSiteRequest struct {
	Email    string
	SiteName string
}

// DeleteSiteResponse collects the response values for the DeleteSite method.
type DeleteSiteResponse struct {
	Err error `json:"-"`
}

// WritePostRequest collects the request parameters for the WritePost method.
type WritePostRequest struct {
	Author   string
	Sitename string
	Filename string
	Content  string
}

// WritePostResponse collects the response values for the WritePost method.
type WritePostResponse struct {
	Err error `json:"-"`
}

// RemovePostRequest collects the request parameters for the RemovePost method.
type RemovePostRequest struct {
	Author   string
	Sitename string
	Filename string
}

// RemovePostResponse collects the response values for the RemovePost method.
type RemovePostResponse struct {
	Err error `json:"-"`
}

// ReadPostRequest collects the request parameters for the ReadPost method.
type ReadPostRequest struct {
	Author   string
	Sitename string
	Filename string
}

// ReadPostResponse collects the response values for the ReadPost method.
type ReadPostResponse struct {
	Content string `json:"content"`
	Err     error  `json:"-"`
}
