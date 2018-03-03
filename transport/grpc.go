package transport

import (
	"context"
	"github.com/go-errors/errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/seagullbird/headr-contentmgr/db"
	"github.com/seagullbird/headr-contentmgr/endpoint"
	"github.com/seagullbird/headr-contentmgr/pb"
	"github.com/seagullbird/headr-contentmgr/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	newpost grpctransport.Handler
	delpost grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.ContentmgrServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		newpost: grpctransport.NewServer(
			endpoints.NewPostEndpoint,
			decodeGRPCNewPostRequest,
			encodeGRPCNewPostResponse,
			options...,
		),
		delpost: grpctransport.NewServer(
			endpoints.DeletePostEndpoint,
			decodeGRPCDeletePostRequest,
			encodeGRPCDeletePostResponse,
			options...,
		),
	}
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	var newpostEndpoint kitendpoint.Endpoint
	{
		newpostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Contentmgr",
			"NewPost",
			encodeGRPCNewPostRequest,
			decodeGRPCNewPostResponse,
			pb.CreateNewPostReply{},
		).Endpoint()
	}

	var delpostEndpoint kitendpoint.Endpoint
	{
		delpostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Contentmgr",
			"DeletePost",
			encodeGRPCDeletePostRequest,
			decodeGRPCDeletePostResponse,
			pb.DeletePostReply{},
		).Endpoint()
	}
	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewPostEndpoint:    newpostEndpoint,
		DeletePostEndpoint: delpostEndpoint,
	}
}

func (s *grpcServer) NewPost(ctx context.Context, req *pb.CreateNewPostRequest) (*pb.CreateNewPostReply, error) {
	_, rep, err := s.newpost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateNewPostReply), nil
}

func (s *grpcServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	_, rep, err := s.delpost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeletePostReply), nil
}

// NewPost
func encodeGRPCNewPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewPostRequest)
	return &pb.CreateNewPostRequest{
		Title:    req.Post.Title,
		Summary:  req.Post.Summary,
		Content:  req.Post.Content,
		Tags:     req.Post.Tags,
		Author:   req.Post.Author,
		Sitename: req.Sitename,
		Date:     req.Post.Date,
	}, nil
}

func decodeGRPCNewPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateNewPostRequest)
	return endpoint.NewPostRequest{
		Post: db.Post{
			Author:   req.Author,
			Sitename: req.Sitename,
			Filename: req.Title,
			Filetype: "md",
			Title:    req.Title,
			Date:     req.Date,
			Draft:    false,
			Tags:     req.Tags,
			Summary:  req.Summary,
			Content:  req.Content,
		},
	}, nil
}

func encodeGRPCNewPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewPostResponse)
	return &pb.CreateNewPostReply{
		Id:  uint64(resp.Id),
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCNewPostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateNewPostReply)
	return endpoint.NewPostResponse{Id: uint(reply.Id), Err: str2err(reply.Err)}, nil
}

// DeletePost
func encodeGRPCDeletePostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeletePostRequest)
	return &pb.DeletePostRequest{
		Id: uint64(req.Id),
	}, nil
}

func decodeGRPCDeletePostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeletePostRequest)
	return endpoint.DeletePostRequest{
		Id: uint(req.Id),
	}, nil
}

func encodeGRPCDeletePostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeletePostResponse)
	return &pb.DeletePostReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCDeletePostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.DeletePostReply)
	return endpoint.DeletePostResponse{Err: str2err(reply.Err)}, nil
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}
