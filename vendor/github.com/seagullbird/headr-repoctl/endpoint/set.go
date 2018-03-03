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
	WritePostEndpoint  endpoint.Endpoint
	RemovePostEndpoint endpoint.Endpoint
	ReadPostEndpoint   endpoint.Endpoint
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
	var writepostEndpoint endpoint.Endpoint
	{
		writepostEndpoint = MakeWritePostEndpoint(svc)
		writepostEndpoint = LoggingMiddleware(logger)(writepostEndpoint)
	}
	var removepostEndpoint endpoint.Endpoint
	{
		removepostEndpoint = MakeRemovePostEndpoint(svc)
		removepostEndpoint = LoggingMiddleware(logger)(removepostEndpoint)
	}
	var readpostEndpoint endpoint.Endpoint
	{
		readpostEndpoint = MakeReadPostEndpoint(svc)
		readpostEndpoint = LoggingMiddleware(logger)(readpostEndpoint)
	}
	return Set{
		NewSiteEndpoint:    newsiteEndpoint,
		DeleteSiteEndpoint: deletesiteEndpoint,
		WritePostEndpoint:  writepostEndpoint,
		RemovePostEndpoint: readpostEndpoint,
		ReadPostEndpoint:   readpostEndpoint,
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

func (s Set) WritePost(ctx context.Context, author, sitename, filename, content string) error {
	resp, err := s.WritePostEndpoint(ctx, WritePostRequest{
		Author:   author,
		Sitename: sitename,
		Filename: filename,
		Content:  content,
	})
	if err != nil {
		return err
	}
	response := resp.(WritePostResponse)
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

func (s Set) ReadPost(ctx context.Context, author, sitename, filename string) (string, error) {
	resp, err := s.ReadPostEndpoint(ctx, ReadPostRequest{
		Author:   author,
		Sitename: sitename,
		Filename: filename,
	})
	if err != nil {
		return "", err
	}
	response := resp.(ReadPostResponse)
	return response.Content, response.Err
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

func MakeWritePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WritePostRequest)
		err = svc.WritePost(ctx, req.Author, req.Sitename, req.Filename, req.Content)
		return WritePostResponse{Err: err}, err
	}
}

func MakeRemovePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RemovePostRequest)
		err = svc.RemovePost(ctx, req.Author, req.Sitename, req.Filename)
		return RemovePostResponse{Err: err}, err
	}
}

func MakeReadPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ReadPostRequest)
		content, err := svc.ReadPost(ctx, req.Author, req.Sitename, req.Filename)
		return ReadPostResponse{Content: content, Err: err}, err
	}
}
