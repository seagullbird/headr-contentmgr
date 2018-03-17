package service

import (
	"context"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
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

func (mw loggingMiddleware) NewSite(ctx context.Context, siteID uint) (err error) {
	err = mw.next.NewSite(ctx, siteID)
	mw.logger.Log("method", "NewSite", "siteID", siteID, "err", err)
	return
}

func (mw loggingMiddleware) DeleteSite(ctx context.Context, siteID uint) (err error) {
	err = mw.next.DeleteSite(ctx, siteID)
	mw.logger.Log("method", "DeleteSite", "siteID", siteID, "err", err)
	return
}

func (mw loggingMiddleware) WritePost(ctx context.Context, siteID uint, filename, content string) (err error) {
	err = mw.next.WritePost(ctx, siteID, filename, content)
	mw.logger.Log("method", "WritePost", "siteID", siteID, "filename", filename, "err", err)
	return
}

func (mw loggingMiddleware) RemovePost(ctx context.Context, siteID uint, filename string) (err error) {
	err = mw.next.RemovePost(ctx, siteID, filename)
	mw.logger.Log("method", "DeletePost", "siteID", siteID, "filename", filename, "err", err)
	return
}

func (mw loggingMiddleware) ReadPost(ctx context.Context, siteID uint, filename string) (content string, err error) {
	content, err = mw.next.ReadPost(ctx, siteID, filename)
	mw.logger.Log("method", "ReadPost", "siteID", siteID, "filename", filename, "err", err)
	return
}
