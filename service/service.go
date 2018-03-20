package service

import (
	"context"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
)

// Service describes a service that deals with content management operations (contentmgr).
type Service interface {
	NewPost(ctx context.Context, post db.Post) (uint, error)
	DeletePost(ctx context.Context, id uint) error
	GetPost(ctx context.Context, id uint) (*db.Post, error)
	GetAllPosts(ctx context.Context) ([]uint, error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(repoctlsvc repoctlservice.Service, store db.Store, logger log.Logger) Service {
	var svc Service
	{
		svc = newBasicService(repoctlsvc, store)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
	store      db.Store
}

func newBasicService(repoctlsvc repoctlservice.Service, store db.Store) basicService {
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
	return id, s.repoctlsvc.WritePost(ctx, post.SiteID, filename, filecontent)
}

func (s basicService) DeletePost(ctx context.Context, id uint) error {
	postptr, _ := s.store.GetPost(id)
	err := s.repoctlsvc.RemovePost(ctx, postptr.SiteID, postptr.Filename+"."+postptr.Filetype)
	if err != nil {
		return err
	}
	s.store.DeletePost(postptr)
	return nil
}

func (s basicService) GetPost(ctx context.Context, id uint) (*db.Post, error) {
	postptr, _ := s.store.GetPost(id)
	content, err := s.repoctlsvc.ReadPost(ctx, postptr.SiteID, postptr.Filename+"."+postptr.Filetype)
	if err != nil {
		return nil, err
	}
	postptr.Content = content
	return postptr, nil
}

func (s basicService) GetAllPosts(ctx context.Context) ([]uint, error) {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	return s.store.GetAllPosts(userID)
}

// EmptyService is only used for transport tests
type EmptyService struct{}

// NewPost implements Service.NewPost
func (e EmptyService) NewPost(ctx context.Context, post db.Post) (uint, error) {
	return 0, nil
}

// DeletePost implements Service.DeletePost
func (e EmptyService) DeletePost(ctx context.Context, id uint) error {
	return nil
}

// GetPost implements Service.GetPost
func (e EmptyService) GetPost(ctx context.Context, id uint) (*db.Post, error) {
	return nil, nil
}

// GetAllPosts implements Service.GetAllPosts
func (e EmptyService) GetAllPosts(ctx context.Context) ([]uint, error) {
	return nil, nil
}
