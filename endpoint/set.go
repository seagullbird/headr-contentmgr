package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/service"
)

type Set struct {
	NewPostEndpoint endpoint.Endpoint
}

func New(svc service.Service, logger log.Logger) Set {
	var NewPostEndpoint endpoint.Endpoint
	{
		NewPostEndpoint = MakeNewPostEndpoint(svc)
		NewPostEndpoint = LoggingMiddleware(logger)(NewPostEndpoint)
	}
	return Set{
		NewPostEndpoint: NewPostEndpoint,
	}
}

func (s Set) NewPost(ctx context.Context, post service.Post) (uint, error) {
	resp, err := s.NewPostEndpoint(ctx, NewPostRequest{Post: post})
	if err != nil {
		return 0, err
	}
	response := resp.(NewPostResponse)
	return response.Id, response.Err
}

func MakeNewPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewPostRequest)
		id, err := svc.NewPost(ctx, req.Post)
		return NewPostResponse{Id: id, Err: err}, err
	}
}

type Failer interface {
	Failed() error
}

func (r NewPostResponse) Failed() error { return r.Err }
