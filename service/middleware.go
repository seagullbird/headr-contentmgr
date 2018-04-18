package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
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

func (mw loggingMiddleware) NewPost(ctx context.Context, post db.Post) (id uint, err error) {
	id, err = mw.next.NewPost(ctx, post)
	mw.logger.Log("method", "NewPost", "id", id, "siteID", post.SiteID, "title", post.Title, "date", post.Date, "err", err)
	return
}

func (mw loggingMiddleware) DeletePost(ctx context.Context, id uint) (uint, error) {
	id, err := mw.next.DeletePost(ctx, id)
	mw.logger.Log("method", "DeletePost", "id", id, "err", err)
	return id, err
}

func (mw loggingMiddleware) GetPost(ctx context.Context, id uint) (db.Post, error) {
	post, err := mw.next.GetPost(ctx, id)
	mw.logger.Log("method", "GetPost", "id", id, "err", err)
	return post, err
}

func (mw loggingMiddleware) GetAllPosts(ctx context.Context) ([]uint, error) {
	postIDs, err := mw.next.GetAllPosts(ctx)
	mw.logger.Log("method", "GetAllPosts")
	return postIDs, err
}

func (mw loggingMiddleware) PatchPost(ctx context.Context, post db.Post) error {
	err := mw.next.PatchPost(ctx, post)
	mw.logger.Log("method", "PatchPost", "post_id", post.ID)
	return err
}
