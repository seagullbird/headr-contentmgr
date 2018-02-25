package transport

import (
	"context"
	"github.com/go-errors/errors"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/pb"
)

// NewSite
func encodeGRPCNewSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewSiteRequest)
	return &pb.NewSiteRequest{Email: req.Email, Sitename: req.SiteName}, nil
}

func decodeGRPCNewSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.NewSiteRequest)
	return endpoint.NewSiteRequest{
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

func decodeGRPCNewSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.NewSiteReply)
	return endpoint.NewSiteResponse{Err: str2err(reply.Err)}, nil
}

// DeleteSite
func encodeGRPCDeleteSiteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteSiteRequest)
	return &pb.DeleteSiteRequest{Email: req.Email, Sitename: req.SiteName}, nil
}

func decodeGRPCDeleteSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteSiteRequest)
	return endpoint.DeleteSiteRequest{
		Email:    req.Email,
		SiteName: req.Sitename,
	}, nil
}

func encodeGRPCDeleteSiteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeleteSiteResponse)
	return &pb.DeleteSiteReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCDeleteSiteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.DeleteSiteReply)
	return endpoint.DeleteSiteResponse{Err: str2err(reply.Err)}, nil
}

// NewPost
func encodeGRPCNewPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewPostRequest)
	return &pb.NewPostRequest{
		Author:   req.Author,
		Sitename: req.Sitename,
		Filename: req.Filename,
		Content:  req.Content,
	}, nil
}

func decodeGRPCNewPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.NewPostRequest)
	return endpoint.NewPostRequest{
		Author:   req.Author,
		Sitename: req.Sitename,
		Filename: req.Filename,
		Content:  req.Content,
	}, nil
}

func encodeGRPCNewPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewPostResponse)
	return &pb.NewPostReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCNewPostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.NewPostReply)
	return endpoint.NewPostResponse{Err: str2err(reply.Err)}, nil
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
