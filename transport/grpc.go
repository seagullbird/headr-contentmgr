package transport

import (
	"context"
	"github.com/go-errors/errors"
	"github.com/go-kit/kit/auth/jwt"
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
	newpost     grpctransport.Handler
	delpost     grpctransport.Handler
	getpost     grpctransport.Handler
	getallposts grpctransport.Handler
	patchpost   grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC ContentmgrServer.
func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.ContentmgrServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(jwt.GRPCToContext()),
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
		getpost: grpctransport.NewServer(
			endpoints.GetPostEndpoint,
			decodeGRPCGetPostRequest,
			encodeGRPCGetPostResponse,
			options...,
		),
		getallposts: grpctransport.NewServer(
			endpoints.GetAllPostsEndpoint,
			decodeGRPCGetAllPostsRequest,
			encodeGRPCGetAllPostsResponse,
			options...,
		),
		patchpost: grpctransport.NewServer(
			endpoints.PatchPostEndpoint,
			decodeGRPCPatchPostRequest,
			encodeGRPCPatchPostResponse,
			options...,
		),
	}
}

// NewGRPCClient returns an ContentmgrService backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport.
func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(jwt.ContextToGRPC()),
	}
	var newpostEndpoint kitendpoint.Endpoint
	{
		newpostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Contentmgr",
			"NewPost",
			encodeGRPCNewPostRequest,
			decodeGRPCNewPostResponse,
			pb.CreateNewPostReply{},
			options...,
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
			options...,
		).Endpoint()
	}
	var getpostEndpoint kitendpoint.Endpoint
	{
		getpostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Contentmgr",
			"GetPost",
			encodeGRPCGetPostRequest,
			decodeGRPCGetPostResponse,
			pb.GetPostReply{},
			options...,
		).Endpoint()
	}
	var getallpostsEndpoint kitendpoint.Endpoint
	{
		getallpostsEndpoint = grpctransport.NewClient(
			conn,
			"pb.Contentmgr",
			"GetAllPosts",
			encodeGRPCGetAllPostsRequest,
			decodeGRPCGetAllPostsResponse,
			pb.GetAllPostsReply{},
			options...,
		).Endpoint()
	}
	var patchpostEndpoint kitendpoint.Endpoint
	{
		patchpostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Contentmgr",
			"PatchPost",
			encodeGRPCPatchPostRequest,
			decodeGRPCPatchPostResponse,
			pb.PatchPostReply{},
			options...,
		).Endpoint()
	}
	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewPostEndpoint:     newpostEndpoint,
		DeletePostEndpoint:  delpostEndpoint,
		GetPostEndpoint:     getpostEndpoint,
		GetAllPostsEndpoint: getallpostsEndpoint,
		PatchPostEndpoint:   patchpostEndpoint,
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

func (s *grpcServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostReply, error) {
	_, rep, err := s.getpost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetPostReply), nil
}

func (s *grpcServer) GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsReply, error) {
	_, rep, err := s.getallposts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllPostsReply), nil
}

func (s *grpcServer) PatchPost(ctx context.Context, req *pb.PatchPostRequest) (*pb.PatchPostReply, error) {
	_, rep, err := s.patchpost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PatchPostReply), nil
}

// NewPost
func encodeGRPCNewPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewPostRequest)
	return &pb.CreateNewPostRequest{
		Title:   req.Post.Title,
		Summary: req.Post.Summary,
		Content: req.Post.Content,
		Tags:    req.Post.Tags,
		SiteId:  uint64(req.Post.SiteID),
		Date:    req.Post.Date,
		Draft:   req.Post.Draft,
	}, nil
}

func decodeGRPCNewPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateNewPostRequest)
	return endpoint.NewPostRequest{
		Post: db.Post{
			SiteID:   uint(req.SiteId),
			Filename: req.Title,
			Filetype: "md",
			Title:    req.Title,
			Date:     req.Date,
			Draft:    req.Draft,
			Tags:     req.Tags,
			Summary:  req.Summary,
			Content:  req.Content,
		},
	}, nil
}

func encodeGRPCNewPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewPostResponse)
	return &pb.CreateNewPostReply{
		Id:  uint64(resp.ID),
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCNewPostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateNewPostReply)
	return endpoint.NewPostResponse{ID: uint(reply.Id), Err: str2err(reply.Err)}, nil
}

// DeletePost
func encodeGRPCDeletePostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeletePostRequest)
	return &pb.DeletePostRequest{
		Id: uint64(req.ID),
	}, nil
}

func decodeGRPCDeletePostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeletePostRequest)
	return endpoint.DeletePostRequest{
		ID: uint(req.Id),
	}, nil
}

func encodeGRPCDeletePostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeletePostResponse)
	return &pb.DeletePostReply{
		Id:  uint64(resp.ID),
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCDeletePostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.DeletePostReply)
	return endpoint.DeletePostResponse{ID: uint(reply.Id), Err: str2err(reply.Err)}, nil
}

// GetPost
func encodeGRPCGetPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetPostRequest)
	return &pb.GetPostRequest{
		Id: uint64(req.ID),
	}, nil
}

func decodeGRPCGetPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetPostRequest)
	return endpoint.GetPostRequest{
		ID: uint(req.Id),
	}, nil
}

func encodeGRPCGetPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetPostResponse)
	return &pb.GetPostReply{
		Title:   resp.Post.Title,
		Summary: resp.Post.Summary,
		Content: resp.Post.Content,
		Tags:    resp.Post.Tags,
		SiteId:  uint64(resp.Post.SiteID),
		Date:    resp.Post.Date,
		Err:     err2str(resp.Err),
		Draft:   resp.Post.Draft,
	}, nil
}

func decodeGRPCGetPostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetPostReply)
	return endpoint.GetPostResponse{
		Post: db.Post{
			Title:   reply.Title,
			Summary: reply.Summary,
			Content: reply.Content,
			Tags:    reply.Tags,
			SiteID:  uint(reply.SiteId),
			Date:    reply.Date,
			Draft:   reply.Draft,
		},
		Err: str2err(reply.Err),
	}, nil
}

// GetAllPosts
func encodeGRPCGetAllPostsRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.GetAllPostsRequest{}, nil
}

func decodeGRPCGetAllPostsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoint.GetAllPostsRequest{}, nil
}

func encodeGRPCGetAllPostsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetAllPostsResponse)
	var temp = make([]uint64, len(resp.PostIDs))
	for i, v := range resp.PostIDs {
		temp[i] = uint64(v)
	}
	return &pb.GetAllPostsReply{
		PostIds: temp,
		Err:     err2str(resp.Err),
	}, nil
}

func decodeGRPCGetAllPostsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetAllPostsReply)
	var temp = make([]uint, len(reply.PostIds))
	for i, v := range reply.PostIds {
		temp[i] = uint(v)
	}
	return endpoint.GetAllPostsResponse{
		PostIDs: temp,
		Err:     str2err(reply.Err),
	}, nil
}

// PatchPost
func encodeGRPCPatchPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.PatchPostRequest)
	return &pb.PatchPostRequest{
		PostId:  uint64(req.Post.ID),
		Title:   req.Post.Title,
		Summary: req.Post.Summary,
		Content: req.Post.Content,
		Tags:    req.Post.Tags,
		Date:    req.Post.Date,
		Draft:   req.Post.Draft,
	}, nil
}

func decodeGRPCPatchPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.PatchPostRequest)
	post := db.Post{
		Title:   req.Title,
		Date:    req.Date,
		Draft:   req.Draft,
		Tags:    req.Tags,
		Summary: req.Summary,
		Content: req.Content,
	}
	post.ID = uint(req.PostId)
	return endpoint.PatchPostRequest{Post: post}, nil
}

func encodeGRPCPatchPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.PatchPostResponse)
	return &pb.PatchPostReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCPatchPostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.PatchPostReply)
	return endpoint.PatchPostResponse{Err: str2err(reply.Err)}, nil
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
