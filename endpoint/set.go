package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
	"github.com/seagullbird/headr-contentmgr/service"
)

// Set collects all of the endpoints that compose an contentmgr service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	NewPostEndpoint     endpoint.Endpoint
	DeletePostEndpoint  endpoint.Endpoint
	GetPostEndpoint     endpoint.Endpoint
	GetAllPostsEndpoint endpoint.Endpoint
}

// New returns a Set that wraps the provided server.
func New(svc service.Service, logger log.Logger) Set {
	return Set{
		NewPostEndpoint:     Middlewares(MakeNewPostEndpoint(svc), logger),
		DeletePostEndpoint:  Middlewares(MakeDeletePostEndpoint(svc), logger),
		GetPostEndpoint:     Middlewares(MakeGetPostEndpoint(svc), logger),
		GetAllPostsEndpoint: Middlewares(MakeGetAllPostsEndpoint(svc), logger),
	}
}

// NewPost implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) NewPost(ctx context.Context, post db.Post) (uint, error) {
	resp, err := s.NewPostEndpoint(ctx, NewPostRequest{Post: post})
	if err != nil {
		return 0, err
	}
	response := resp.(NewPostResponse)
	return response.ID, response.Err
}

// DeletePost implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) DeletePost(ctx context.Context, id uint) error {
	resp, err := s.DeletePostEndpoint(ctx, DeletePostRequest{ID: id})
	if err != nil {
		return err
	}
	response := resp.(DeletePostResponse)
	return response.Err
}

// GetPost implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) GetPost(ctx context.Context, id uint) (*db.Post, error) {
	resp, err := s.GetPostEndpoint(ctx, GetPostRequest{ID: id})
	if err != nil {
		return nil, err
	}
	response := resp.(GetPostResponse)
	return &response.Post, err
}

// GetAllPosts implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Set) GetAllPosts(ctx context.Context) ([]uint, error) {
	resp, err := s.GetAllPostsEndpoint(ctx, GetAllPostsRequest{})
	if err != nil {
		return nil, err
	}
	response := resp.(GetAllPostsResponse)
	return response.PostIDs, err
}

// MakeNewPostEndpoint constructs a NewPost endpoint wrapping the service.
func MakeNewPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewPostRequest)
		id, err := svc.NewPost(ctx, req.Post)
		return NewPostResponse{ID: id, Err: err}, err
	}
}

// MakeDeletePostEndpoint constructs a DeletePost endpoint wrapping the service.
func MakeDeletePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeletePostRequest)
		err = svc.DeletePost(ctx, req.ID)
		return DeletePostResponse{Err: err}, err
	}
}

// MakeGetPostEndpoint constructs a GetPost endpoint wrapping the service.
func MakeGetPostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetPostRequest)
		postptr, err := svc.GetPost(ctx, req.ID)
		return GetPostResponse{Post: *postptr, Err: err}, err
	}
}

// MakeGetAllPostsEndpoint constructs a GetAllPosts endpoint wrapping the service.
func MakeGetAllPostsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		postIDs, err := svc.GetAllPosts(ctx)
		return GetAllPostsResponse{PostIDs: postIDs, Err: err}, err
	}
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so if they've
// failed, and if so encode them using a separate write path based on the error.
type Failer interface {
	Failed() error
}

// Failed implements Failer.
func (r NewPostResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r DeletePostResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r GetPostResponse) Failed() error { return r.Err }

// Failed implements Failer.
func (r GetAllPostsResponse) Failed() error { return r.Err }
