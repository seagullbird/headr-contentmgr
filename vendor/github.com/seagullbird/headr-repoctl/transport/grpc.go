package transport

import (
	"context"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/pb"
	"github.com/seagullbird/headr-repoctl/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	newsite    grpctransport.Handler
	deletesite grpctransport.Handler
	newpost    grpctransport.Handler
	rmpost     grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.RepoctlServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		newsite: grpctransport.NewServer(
			endpoints.NewSiteEndpoint,
			decodeGRPCNewSiteRequest,
			encodeGRPCNewSiteResponse,
			options...,
		),
		deletesite: grpctransport.NewServer(
			endpoints.DeleteSiteEndpoint,
			decodeGRPCDeleteSiteRequest,
			encodeGRPCDeleteSiteResponse,
			options...,
		),
		newpost: grpctransport.NewServer(
			endpoints.NewPostEndpoint,
			decodeGRPCNewPostRequest,
			encodeGRPCNewPostResponse,
			options...,
		),
		rmpost: grpctransport.NewServer(
			endpoints.RemovePostEndpoint,
			decodeGRPCRemovePostRequest,
			encodeGRPCRemovePostResponse,
			options...,
		),
	}
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	var newsiteEndpoint kitendpoint.Endpoint
	{
		newsiteEndpoint = grpctransport.NewClient(
			conn,
			"pb.Repoctl",
			"NewSite",
			encodeGRPCNewSiteRequest,
			decodeGRPCNewSiteResponse,
			pb.NewSiteReply{},
		).Endpoint()
	}
	var deletesiteEndpoint kitendpoint.Endpoint
	{
		deletesiteEndpoint = grpctransport.NewClient(
			conn,
			"pb.Repoctl",
			"DeleteSite",
			encodeGRPCDeleteSiteRequest,
			decodeGRPCDeleteSiteResponse,
			pb.DeleteSiteReply{},
		).Endpoint()
	}
	var newpostEndpoint kitendpoint.Endpoint
	{
		newpostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Repoctl",
			"NewPost",
			encodeGRPCNewPostRequest,
			decodeGRPCNewPostResponse,
			pb.NewPostReply{},
		).Endpoint()
	}
	var deletepostEndpoint kitendpoint.Endpoint
	{
		deletepostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Repoctl",
			"RemovePost",
			encodeGRPCRemovePostRequest,
			decodeGRPCRemovePostResponse,
			pb.RemovePostReply{},
		).Endpoint()
	}
	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewSiteEndpoint:    newsiteEndpoint,
		DeleteSiteEndpoint: deletesiteEndpoint,
		NewPostEndpoint:    newpostEndpoint,
		RemovePostEndpoint: deletepostEndpoint,
	}
}

func (s *grpcServer) NewSite(ctx context.Context, req *pb.NewSiteRequest) (*pb.NewSiteReply, error) {
	_, rep, err := s.newsite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.NewSiteReply), nil
}

func (s *grpcServer) DeleteSite(ctx context.Context, req *pb.DeleteSiteRequest) (*pb.DeleteSiteReply, error) {
	_, rep, err := s.deletesite.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteSiteReply), nil
}

func (s *grpcServer) NewPost(ctx context.Context, req *pb.NewPostRequest) (*pb.NewPostReply, error) {
	_, rep, err := s.newpost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.NewPostReply), nil
}

func (s *grpcServer) RemovePost(ctx context.Context, req *pb.RemovePostRequest) (*pb.RemovePostReply, error) {
	_, rep, err := s.rmpost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RemovePostReply), nil
}
