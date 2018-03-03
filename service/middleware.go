package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
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

func (mw loggingMiddleware) NewPost(ctx context.Context, post db.Post) (id uint, err error) {
	id, err = mw.next.NewPost(ctx, post)
	mw.logger.Log("method", "NewPost", "id", id, "author", post.Author, "sitename", post.Sitename, "title", post.Title, "date", post.Date, "err", err)
	return
}
