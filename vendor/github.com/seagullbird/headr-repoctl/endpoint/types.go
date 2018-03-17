package endpoint

// NewSiteRequest collects the request parameters for the NewSite method.
type NewSiteRequest struct {
	SiteID uint
}

// NewSiteResponse collects the response values for the NewSite method.
type NewSiteResponse struct {
	Err error `json:"-"`
}

// DeleteSiteRequest collects the request parameters for the DeleteSite method.
type DeleteSiteRequest struct {
	SiteID uint
}

// DeleteSiteResponse collects the response values for the DeleteSite method.
type DeleteSiteResponse struct {
	Err error `json:"-"`
}

// WritePostRequest collects the request parameters for the WritePost method.
type WritePostRequest struct {
	SiteID   uint
	Filename string
	Content  string
}

// WritePostResponse collects the response values for the WritePost method.
type WritePostResponse struct {
	Err error `json:"-"`
}

// RemovePostRequest collects the request parameters for the RemovePost method.
type RemovePostRequest struct {
	SiteID   uint
	Filename string
}

// RemovePostResponse collects the response values for the RemovePost method.
type RemovePostResponse struct {
	Err error `json:"-"`
}

// ReadPostRequest collects the request parameters for the ReadPost method.
type ReadPostRequest struct {
	SiteID   uint
	Filename string
}

// ReadPostResponse collects the response values for the ReadPost method.
type ReadPostResponse struct {
	Content string `json:"content"`
	Err     error  `json:"-"`
}
