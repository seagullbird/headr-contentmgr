package endpoint_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-common/auth"
	"github.com/seagullbird/headr-contentmgr/db"
	"github.com/seagullbird/headr-contentmgr/endpoint"
	svcmock "github.com/seagullbird/headr-contentmgr/service/mock"
	"testing"
)

func TestSet(t *testing.T) {
	// Mocking Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	endpoints := endpoint.New(mockSvc, logger)

	// Login
	ctx := context.Background()
	accessToken := auth.Login()
	ctx = context.WithValue(ctx, jwt.JWTTokenContextKey, accessToken)

	dummyError := errors.New("dummy error")
	tests := []struct {
		name string
		rets map[string][]interface{}
	}{
		{"No Error", map[string][]interface{}{
			"NewPost":     {uint(1), nil},
			"DeletePost":  {nil},
			"GetPost":     {db.Post{}, nil},
			"PatchPost":   {nil},
			"GetAllPosts": {[]uint{}, nil},
		}},
		{"Dummy Error", map[string][]interface{}{
			"NewPost":     {uint(0), dummyError},
			"DeletePost":  {dummyError},
			"GetPost":     {db.Post{}, dummyError},
			"PatchPost":   {dummyError},
			"GetAllPosts": {[]uint{}, dummyError},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set EXPECTS
			times := 2
			mockSvc.EXPECT().NewPost(gomock.Any(), gomock.Any()).Return(tt.rets["NewPost"]...).Times(times)
			mockSvc.EXPECT().DeletePost(gomock.Any(), gomock.Any()).Return(tt.rets["DeletePost"]...).Times(times)
			mockSvc.EXPECT().GetPost(gomock.Any(), gomock.Any()).Return(tt.rets["GetPost"]...).Times(times)
			mockSvc.EXPECT().PatchPost(gomock.Any(), gomock.Any()).Return(tt.rets["PatchPost"]...).Times(times)
			mockSvc.EXPECT().GetAllPosts(gomock.Any()).Return(tt.rets["GetAllPosts"]...).Times(times)

			t.Run("NewPost", func(t *testing.T) {
				var post db.Post
				setPostID, setErr := endpoints.NewPost(ctx, post)
				svcPostID, svcErr := mockSvc.NewPost(ctx, post)
				if setPostID != svcPostID || setErr != svcErr {
					t.Fatal("\nsetPostID: ", setPostID, "\nsetErr: ", setErr, "\nsvcPostID: ", svcPostID, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("DeletePost", func(t *testing.T) {
				postID := uint(1)
				setErr := endpoints.DeletePost(ctx, postID)
				svcErr := mockSvc.DeletePost(ctx, postID)
				if setErr != svcErr {
					t.Fatal("\nsetErr: ", setErr, "\nsvcErr: ", svcErr)

				}
			})
			t.Run("GetPost", func(t *testing.T) {
				postID := uint(1)
				setOutput, setErr := endpoints.GetPost(ctx, postID)
				svcOutput, svcErr := mockSvc.GetPost(ctx, postID)
				if setErr != svcErr {
					t.Fatal("\nsetOutput: ", setOutput, "\nsetErr: ", setErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("PatchPost", func(t *testing.T) {
				var post db.Post
				setErr := endpoints.PatchPost(ctx, post)
				svcErr := mockSvc.PatchPost(ctx, post)
				if setErr != svcErr {
					t.Fatal("\nsetErr: ", setErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("GetAllPosts", func(t *testing.T) {
				setOutput, setErr := endpoints.GetAllPosts(ctx)
				svcOutput, svcErr := mockSvc.GetAllPosts(ctx)
				if setErr != svcErr {
					t.Fatal("\nsetOutput: ", setOutput, "\nsetErr: ", setErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
		})
	}
}

// In fact this part is tested in grpc_test.TestGRPCTransport, dual here for good coverage report
func TestSetBadEndpoint(t *testing.T) {
	makeBadEndpoint := func(resp interface{}) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := resp.(endpoint.Failer)
			return nil, r.Failed()
		}
	}

	endpoints := endpoint.Set{
		NewPostEndpoint:     makeBadEndpoint(endpoint.NewPostResponse{ID: 1, Err: errors.New("dummy error")}),
		DeletePostEndpoint:  makeBadEndpoint(endpoint.DeletePostResponse{Err: errors.New("dummy error")}),
		GetPostEndpoint:     makeBadEndpoint(endpoint.GetPostResponse{Err: errors.New("dummy error")}),
		GetAllPostsEndpoint: makeBadEndpoint(endpoint.GetAllPostsResponse{PostIDs: []uint{}, Err: errors.New("dummy error")}),
		PatchPostEndpoint:   makeBadEndpoint(endpoint.PatchPostResponse{Err: errors.New("dummy error")}),
	}

	expectedMsg := "dummy error"
	if _, err := endpoints.NewPost(context.Background(), db.Post{}); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.DeletePost(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.GetPost(context.Background(), 1); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if _, err := endpoints.GetAllPosts(context.Background()); err.Error() != expectedMsg {
		t.Fatal(err)
	}
	if err := endpoints.PatchPost(context.Background(), db.Post{}); err.Error() != expectedMsg {
		t.Fatal(err)
	}
}
