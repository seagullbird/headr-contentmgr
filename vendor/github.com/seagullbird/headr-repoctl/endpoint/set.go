package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-repoctl/service"
)

type Set struct {
	NewSiteEndpoint    endpoint.Endpoint
	DeleteSiteEndpoint endpoint.Endpoint
	NewPostEndpoint    endpoint.Endpoint
	RemovePostEndpoint endpoint.Endpoint
}

func New(svc service.Service, logger log.Logger) Set {
	var newsiteEndpoint endpoint.Endpoint
	{
		newsiteEndpoint = MakeNewSiteEndpoint(svc)
		newsiteEndpoint = LoggingMiddleware(logger)(newsiteEndpoint)
	}
	var deletesiteEndpoint endpoint.Endpoint
	{
		deletesiteEndpoint = MakeDeleteSiteEndpoint(svc)
		deletesiteEndpoint = LoggingMiddleware(logger)(deletesiteEndpoint)
	}
	var newpostEndpoint endpoint.Endpoint
	{
		newpostEndpoint = MakeNewPostEndpoint(svc)
		newpostEndpoint = LoggingMiddleware(logger)(newpostEndpoint)
	}
	var removepostEndpoint endpoint.Endpoint
	{
		removepostEndpoint = MakeRemovePostEndpoint(svc)
		removepostEndpoint = LoggingMiddleware(logger)(removepostEndpoint)
	}
	return Set{
		NewSiteEndpoint:    newsiteEndpoint,
		DeleteSiteEndpoint: deletesiteEndpoint,
		NewPostEndpoint:    newpostEndpoint,
		RemovePostEndpoint: removepostEndpoint,
	}
}

func (s Set) NewSite(ctx context.Context, email, sitename string) error {
	resp, err := s.NewSiteEndpoint(ctx, NewSiteRequest{Email: email, SiteName: sitename})
	if err != nil {
		return err
	}
	response := resp.(NewSiteResponse)
	return response.Err
}

func (s Set) DeleteSite(ctx context.Context, email, sitename string) error {
	resp, err := s.DeleteSiteEndpoint(ctx, DeleteSiteRequest{Email: email, SiteName: sitename})
	if err != nil {
		return err
	}
	response := resp.(DeleteSiteResponse)
	return response.Err
}

func (s Set) NewPost(ctx context.Context, author, sitename, filename, content string) error {
	resp, err := s.NewPostEndpoint(ctx, NewPostRequest{
		Author:   author,
		Sitename: sitename,
		Filename: filename,
		Content:  content,
	})
	if err != nil {
		return err
	}
	response := resp.(NewPostResponse)
	return response.Err
}

func (s Set) RemovePost(ctx context.Context, author, sitename, filename string) error {
	resp, err := s.RemovePostEndpoint(ctx, RemovePostRequest{
		Author:   author,
		Sitename: sitename,
		Filename: filename,
	})
	if err != nil {
		return err
	}
	response := resp.(RemovePostResponse)
	return response.Err
}

func MakeNewSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewSiteRequest)
		err = svc.NewSite(ctx, req.Email, req.SiteName)
		return NewSiteResponse{Err: err}, err
	}
}

func MakeDeleteSiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteSiteRequest)
		err = svc.DeleteSite(ctx, req.Email, req.SiteName)
		return DeleteSiteResponse{Err: err}, err
	}
}

func MakeNewPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewPostRequest)
		err = svc.NewPost(ctx, req.Author, req.Sitename, req.Filename, req.Content)
		return NewPostResponse{Err: err}, err
	}
}

func MakeRemovePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RemovePostRequest)
		err = svc.RemovePost(ctx, req.Author, req.Sitename, req.Filename)
		return RemovePostResponse{Err: err}, err
	}
}
