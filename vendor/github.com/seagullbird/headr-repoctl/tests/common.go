package tests

import (
	"context"
)

// Site is a convenient struct for sending parameters only used in package tests.
type Site struct {
	ctx      context.Context
	email    string
	sitename string
}

// Post is a convenient struct for sending parameters only used in package tests.
type Post struct {
	ctx      context.Context
	author   string
	sitename string
	filename string
	content  string
}
