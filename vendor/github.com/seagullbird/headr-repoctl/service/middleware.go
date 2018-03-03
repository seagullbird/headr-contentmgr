package service

import (
	"context"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			logger,
			next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) NewSite(ctx context.Context, email, sitename string) (err error) {
	err = mw.next.NewSite(ctx, email, sitename)
	mw.logger.Log("method", "NewSite", "email", email, "sitename", sitename, "err", err)
	return
}

func (mw loggingMiddleware) DeleteSite(ctx context.Context, email, sitename string) (err error) {
	err = mw.next.DeleteSite(ctx, email, sitename)
	mw.logger.Log("method", "DeleteSite", "email", email, "sitename", sitename, "err", err)
	return
}

func (mw loggingMiddleware) NewPost(ctx context.Context, author, sitename, filename, content string) (err error) {
	err = mw.next.NewPost(ctx, author, sitename, filename, content)
	mw.logger.Log("method", "NewPost", "author", author, "sitename", sitename, "filename", filename, "err", err)
	return
}

func (mw loggingMiddleware) RemovePost(ctx context.Context, author, sitename, filename string) (err error) {
	err = mw.next.RemovePost(ctx, author, sitename, filename)
	mw.logger.Log("method", "DeletePost", "author", author, "sitename", sitename, "filename", filename, "err", err)
	return
}
