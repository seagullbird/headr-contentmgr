package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
)

type Service interface {
	NewPost(ctx context.Context, post Post) error
}

func New(repoctlsvc repoctlservice.Service, store db.Store, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(repoctlsvc, store)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
	store      db.Store
}

func NewBasicService(repoctlsvc repoctlservice.Service, store db.Store) basicService {
	return basicService{
		repoctlsvc: repoctlsvc,
		store:      store,
	}
}

func (s basicService) NewPost(ctx context.Context, post Post) error {
	s.store.InsertPost(post.Model())
	filename := post.Filename + "." + post.Filetype
	filecontent := post.String()
	return s.repoctlsvc.NewPost(ctx, post.Author, post.Sitename, filename, filecontent)
}
