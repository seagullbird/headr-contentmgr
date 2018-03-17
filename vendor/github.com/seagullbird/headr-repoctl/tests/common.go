package tests

import (
	"context"
)

// Site is a convenient struct for sending parameters only used in package tests.
type Site struct {
	ctx    context.Context
	siteID uint
}

// Post is a convenient struct for sending parameters only used in package tests.
type Post struct {
	ctx      context.Context
	siteID   uint
	filename string
	content  string
}
