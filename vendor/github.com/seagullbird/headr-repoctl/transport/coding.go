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
	return &pb.NewSiteRequest{SiteId: uint64(req.SiteID)}, nil
}

func decodeGRPCNewSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.NewSiteRequest)
	return endpoint.NewSiteRequest{
		SiteID: uint(req.SiteId),
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
	return &pb.DeleteSiteRequest{SiteId: uint64(req.SiteID)}, nil
}

func decodeGRPCDeleteSiteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteSiteRequest)
	return endpoint.DeleteSiteRequest{
		SiteID: uint(req.SiteId),
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

// WritePost
func encodeGRPCWritePostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.WritePostRequest)
	return &pb.WritePostRequest{
		SiteId:   uint64(req.SiteID),
		Filename: req.Filename,
		Content:  req.Content,
	}, nil
}

func decodeGRPCWritePostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WritePostRequest)
	return endpoint.WritePostRequest{
		SiteID:   uint(req.SiteId),
		Filename: req.Filename,
		Content:  req.Content,
	}, nil
}

func encodeGRPCWritePostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.WritePostResponse)
	return &pb.WritePostReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCWritePostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.WritePostReply)
	return endpoint.WritePostResponse{Err: str2err(reply.Err)}, nil
}

// RemovePost
func encodeGRPCRemovePostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.RemovePostRequest)
	return &pb.RemovePostRequest{
		SiteId:   uint64(req.SiteID),
		Filename: req.Filename,
	}, nil
}

func decodeGRPCRemovePostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemovePostRequest)
	return endpoint.RemovePostRequest{
		SiteID:   uint(req.SiteId),
		Filename: req.Filename,
	}, nil
}

func encodeGRPCRemovePostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.RemovePostResponse)
	return &pb.RemovePostReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCRemovePostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RemovePostReply)
	return endpoint.RemovePostResponse{Err: str2err(reply.Err)}, nil
}

// ReadPost
func encodeGRPCReadPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ReadPostRequest)
	return &pb.ReadPostRequest{
		SiteId:   uint64(req.SiteID),
		Filename: req.Filename,
	}, nil
}

func decodeGRPCReadPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ReadPostRequest)
	return endpoint.ReadPostRequest{
		SiteID:   uint(req.SiteId),
		Filename: req.Filename,
	}, nil
}

func encodeGRPCReadPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ReadPostResponse)
	return &pb.ReadPostReply{
		Content: resp.Content,
		Err:     err2str(resp.Err),
	}, nil
}

func decodeGRPCReadPostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ReadPostReply)
	return endpoint.ReadPostResponse{Content: reply.Content, Err: str2err(reply.Err)}, nil
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
