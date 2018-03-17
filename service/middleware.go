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
	mw.logger.Log("method", "NewPost", "id", id, "siteID", post.SiteID, "title", post.Title, "date", post.Date, "err", err)
	return
}

func (mw loggingMiddleware) DeletePost(ctx context.Context, id uint) error {
	err := mw.next.DeletePost(ctx, id)
	mw.logger.Log("method", "DeletePost", "id", id, "err", err)
	return err
}

func (mw loggingMiddleware) GetPost(ctx context.Context, id uint) (*db.Post, error) {
	postptr, err := mw.next.GetPost(ctx, id)
	mw.logger.Log("method", "GetPost", "id", id, "err", err)
	return postptr, err
}
