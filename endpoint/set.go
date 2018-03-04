package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
	"github.com/seagullbird/headr-contentmgr/service"
)

type Set struct {
	NewPostEndpoint    endpoint.Endpoint
	DeletePostEndpoint endpoint.Endpoint
	GetPostEndpoint    endpoint.Endpoint
}

func New(svc service.Service, logger log.Logger) Set {
	var newPostEndpoint endpoint.Endpoint
	{
		newPostEndpoint = MakeNewPostEndpoint(svc)
		newPostEndpoint = LoggingMiddleware(logger)(newPostEndpoint)
	}
	var deletePostEndpoint endpoint.Endpoint
	{
		deletePostEndpoint = MakeDeletePostEndpoint(svc)
		deletePostEndpoint = LoggingMiddleware(logger)(deletePostEndpoint)
	}
	var getPostEndpoint endpoint.Endpoint
	{
		getPostEndpoint = MakeGetPostEndpoint(svc)
		getPostEndpoint = LoggingMiddleware(logger)(getPostEndpoint)
	}
	return Set{
		NewPostEndpoint:    newPostEndpoint,
		DeletePostEndpoint: deletePostEndpoint,
		GetPostEndpoint:    getPostEndpoint,
	}
}

func (s Set) NewPost(ctx context.Context, post db.Post) (uint, error) {
	resp, err := s.NewPostEndpoint(ctx, NewPostRequest{Post: post})
	if err != nil {
		return 0, err
	}
	response := resp.(NewPostResponse)
	return response.Id, response.Err
}

func (s Set) DeletePost(ctx context.Context, id uint) error {
	resp, err := s.DeletePostEndpoint(ctx, DeletePostRequest{Id: id})
	if err != nil {
		return err
	}
	response := resp.(DeletePostResponse)
	return response.Err
}

func (s Set) GetPost(ctx context.Context, id uint) (*db.Post, error) {
	resp, err := s.GetPostEndpoint(ctx, GetPostRequest{Id: id})
	if err != nil {
		return nil, err
	}
	response := resp.(GetPostResponse)
	return &response.Post, err
}

func MakeNewPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewPostRequest)
		id, err := svc.NewPost(ctx, req.Post)
		return NewPostResponse{Id: id, Err: err}, err
	}
}

func MakeDeletePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeletePostRequest)
		err = svc.DeletePost(ctx, req.Id)
		return DeletePostResponse{Err: err}, err
	}
}

func MakeGetPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetPostRequest)
		postptr, err := svc.GetPost(ctx, req.Id)
		return GetPostResponse{Post: *postptr, Err: err}, err
	}
}

type Failer interface {
	Failed() error
}

func (r NewPostResponse) Failed() error    { return r.Err }
func (r DeletePostResponse) Failed() error { return r.Err }
func (r GetPostResponse) Failed() error    { return r.Err }
