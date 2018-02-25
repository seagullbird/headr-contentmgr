package service

import (
	"context"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// Logging Middleware
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

func (mw loggingMiddleware) NewPost(ctx context.Context, post Post) (err error) {
	err = mw.next.NewPost(ctx, post)
	mw.logger.Log("method", "NewPost", "author", post.Author, "sitename", post.Sitename, "title", post.FM.Title, "err", err)
	return
}
