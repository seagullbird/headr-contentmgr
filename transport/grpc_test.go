package transport_test

import (
	"context"
	"errors"
	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-common/auth"
	"github.com/seagullbird/headr-contentmgr/db"
	"github.com/seagullbird/headr-contentmgr/endpoint"
	"github.com/seagullbird/headr-contentmgr/pb"
	svcmock "github.com/seagullbird/headr-contentmgr/service/mock"
	"github.com/seagullbird/headr-contentmgr/transport"
	"google.golang.org/grpc"
	"net"
	"testing"
)

const (
	port = ":1234"
)

func startServer(t *testing.T, baseServer *grpc.Server, endpoints endpoint.Set, logger log.Logger) {
	grpcServer := transport.NewGRPCServer(endpoints, logger)
	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatal(err)
	}
	pb.RegisterContentmgrServer(baseServer, grpcServer)
	baseServer.Serve(grpcListener)
}

func TestGRPCApplication(t *testing.T) {
	// Mocking service.Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	// Set mock service expectations
	dummyError := errors.New("dummy error")
	for _, rets := range []map[string][]interface{}{
		{
			"NewPost":     {uint(1), nil},
			"DeletePost":  {nil},
			"GetPost":     {db.Post{}, nil},
			"PatchPost":   {nil},
			"GetAllPosts": {[]uint{}, nil},
		},
		{
			"NewPost":     {uint(0), dummyError},
			"DeletePost":  {dummyError},
			"GetPost":     {db.Post{}, dummyError},
			"PatchPost":   {dummyError},
			"GetAllPosts": {[]uint{}, dummyError},
		},
	} {
		times := 2
		mockSvc.EXPECT().NewPost(gomock.Any(), gomock.Any()).Return(rets["NewPost"]...).Times(times)
		mockSvc.EXPECT().DeletePost(gomock.Any(), gomock.Any()).Return(rets["DeletePost"]...).Times(times)
		mockSvc.EXPECT().GetPost(gomock.Any(), gomock.Any()).Return(rets["GetPost"]...).Times(times)
		mockSvc.EXPECT().PatchPost(gomock.Any(), gomock.Any()).Return(rets["PatchPost"]...).Times(times)
		mockSvc.EXPECT().GetAllPosts(gomock.Any()).Return(rets["GetAllPosts"]...).Times(times)
	}

	// Start GRPC server with the mock service
	logger := log.NewNopLogger()
	endpoints := endpoint.New(mockSvc, logger)
	baseServer := grpc.NewServer()
	go startServer(t, baseServer, endpoints, logger)

	// Start GRPC client
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := transport.NewGRPCClient(conn, nil)
	// Login
	ctx := context.Background()
	accessToken := auth.Login()
	ctx = context.WithValue(ctx, jwt.JWTTokenContextKey, accessToken)

	// testcases
	tests := []struct {
		name   string
		judger func(err1, err2 error) bool
	}{
		{
			"No Error",
			func(err1, err2 error) bool {
				if err1 != nil || err2 != nil {
					return false
				}
				return true
			},
		},
		{
			"Dummy Error",
			func(err1, err2 error) bool {
				if err1.Error() != "dummy error" || err2.Error() != "dummy error" {
					return false
				}
				return true
			},
		},
	}

	// Start tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("NewPost", func(t *testing.T) {
				var post db.Post
				clientPostID, clientErr := client.NewPost(ctx, post)
				svcPostID, svcErr := mockSvc.NewPost(ctx, post)
				if clientPostID != svcPostID || !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientPostID: ", clientPostID, "\nclientErr: ", clientErr, "\nsvcPostID: ", svcPostID, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("DeletePost", func(t *testing.T) {
				siteID := uint(1)
				clientErr := client.DeletePost(ctx, siteID)
				svcErr := mockSvc.DeletePost(ctx, siteID)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)

				}
			})
			t.Run("GetPost", func(t *testing.T) {
				postID := uint(1)
				clientOutput, clientErr := client.GetPost(ctx, postID)
				svcOutput, svcErr := mockSvc.GetPost(ctx, postID)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientOutput: ", clientOutput, "\nclientErr: ", clientErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("PatchPost", func(t *testing.T) {
				var post db.Post
				clientErr := client.PatchPost(ctx, post)
				svcErr := mockSvc.PatchPost(ctx, post)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("GetAllPosts", func(t *testing.T) {
				clientpostIDs, clientErr := client.GetAllPosts(ctx)
				svcPostIDs, svcErr := mockSvc.GetAllPosts(ctx)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientpostIDs: ", clientpostIDs, "\nclientErr: ", clientErr, "\nsvcPostIDs: ", svcPostIDs, "\nsvcErr: ", svcErr)
				}
			})
		})
	}

	baseServer.Stop()
}

func TestGRPCTransport(t *testing.T) {
	makeBadEndpoint := func() kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return nil, errors.New("dummy error")
		}
	}

	endpoints := endpoint.Set{
		NewPostEndpoint:     makeBadEndpoint(),
		DeletePostEndpoint:  makeBadEndpoint(),
		GetPostEndpoint:     makeBadEndpoint(),
		PatchPostEndpoint:   makeBadEndpoint(),
		GetAllPostsEndpoint: makeBadEndpoint(),
	}
	baseServer := grpc.NewServer()
	go startServer(t, baseServer, endpoints, log.NewNopLogger())

	// Start GRPC client
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := transport.NewGRPCClient(conn, nil)
	var post db.Post
	expectedMsg := "rpc error: code = Unknown desc = dummy error"
	if _, err := client.NewPost(context.Background(), post); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.DeletePost(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := client.GetPost(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := client.PatchPost(context.Background(), post); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := client.GetAllPosts(context.Background()); err.Error() != expectedMsg {
		t.Fatal(err)
	}

	baseServer.Stop()
}
