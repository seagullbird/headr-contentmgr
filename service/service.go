package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
)

type Service interface {
	NewPost(ctx context.Context, post db.Post) (uint, error)
	DeletePost(ctx context.Context, id uint) error
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

func (s basicService) NewPost(ctx context.Context, post db.Post) (uint, error) {
	id, err := s.store.InsertPost(&post)
	if err != nil {
		return 0, err
	}
	filename := post.Filename + "." + post.Filetype
	filecontent := post.String()
	return id, s.repoctlsvc.NewPost(ctx, post.Author, post.Sitename, filename, filecontent)
}

func (s basicService) DeletePost(ctx context.Context, id uint) error {
	postptr, _ := s.store.GetPost(id)
	err := s.repoctlsvc.DeletePost(ctx, postptr.Author, postptr.Sitename, postptr.Filename+"."+postptr.Filetype)
	if err != nil {
		return err
	}
	s.store.DeletePost(postptr)
	return nil
}
