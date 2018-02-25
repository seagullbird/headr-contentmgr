package transport

import (
	"context"
	"github.com/go-errors/errors"
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
	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewSiteEndpoint:    newsiteEndpoint,
		DeleteSiteEndpoint: deletesiteEndpoint,
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

func encodeGRPCNewSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewSiteRequest)
	return &pb.NewSiteRequest{Email: req.Email, Sitename: req.SiteName}, nil
}

func encodeGRPCDeleteSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteSiteRequest)
	return &pb.DeleteSiteRequest{Email: req.Email, Sitename: req.SiteName}, nil
}

func decodeGRPCNewSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.NewSiteRequest)
	return endpoint.NewSiteRequest{
		Email:    req.Email,
		SiteName: req.Sitename,
	}, nil
}

func decodeGRPCDeleteSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteSiteRequest)
	return endpoint.DeleteSiteRequest{
		Email:    req.Email,
		SiteName: req.Sitename,
	}, nil
}

func encodeGRPCNewSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewSiteResponse)
	return &pb.NewSiteReply{
		Err: err2str(resp.Err),
	}, nil
}

func encodeGRPCDeleteSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeleteSiteResponse)
	return &pb.DeleteSiteReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCNewSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.NewSiteReply)
	return endpoint.NewSiteResponse{Err: str2err(reply.Err)}, nil
}

func decodeGRPCDeleteSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.DeleteSiteReply)
	return endpoint.DeleteSiteResponse{Err: str2err(reply.Err)}, nil
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
