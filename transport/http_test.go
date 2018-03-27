package transport

import (
	"bytes"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-common/auth"
	"github.com/seagullbird/headr-contentmgr/endpoint"
	svcmock "github.com/seagullbird/headr-contentmgr/service/mock"
	//"io/ioutil"
	"net/http"
	"net/http/httptest"
	//"strings"
	"github.com/seagullbird/headr-contentmgr/db"
	"io/ioutil"
	"strings"
	"testing"
)

func TestHTTPForbidden(t *testing.T) {
	// Mocking service.Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	logger := log.NewNopLogger()
	endpoints := endpoint.New(mockSvc, logger)
	server := httptest.NewServer(NewHTTPHandler(endpoints, logger))
	defer server.Close()
	client := &http.Client{}

	tests := []struct {
		name         string
		path         string
		method       string
		body         string
		expectedCode int
	}{
		{"DeletePostBadRouting", "/posts/a", "DELETE", "", http.StatusBadRequest},
		{"GetPostBadRouting", "/posts/a", "GET", "", http.StatusBadRequest},
		{"PatchPostBadRouting", "/posts/a", "PATCH", "", http.StatusBadRequest},
		{"NewPost", "/posts/", "POST", "{}", http.StatusForbidden},
		{"DeletePost", "/posts/1", "DELETE", "", http.StatusForbidden},
		{"GetPost", "/posts/1", "GET", "", http.StatusForbidden},
		{"PatchPost", "/posts/1", "PATCH", "{}", http.StatusForbidden},
		{"GetAllPosts", "/posts/", "GET", "", http.StatusForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := []byte(tt.body)
			req, err := http.NewRequest(tt.method, server.URL+tt.path, bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Error in creating %s to %s: %v", tt.method, tt.path, err)
			}
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error in %s to %s: %v", tt.method, tt.path, err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.expectedCode {
				t.Fatalf("Unexpected status code: %d\n Status code should be %d", resp.StatusCode, tt.expectedCode)
			}
		})
	}
}

func TestHTTP(t *testing.T) {
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
		times := 1
		mockSvc.EXPECT().NewPost(gomock.Any(), gomock.Any()).Return(rets["NewPost"]...).Times(times)
		mockSvc.EXPECT().DeletePost(gomock.Any(), gomock.Any()).Return(rets["DeletePost"]...).Times(times)
		mockSvc.EXPECT().GetPost(gomock.Any(), gomock.Any()).Return(rets["GetPost"]...).Times(times)
		mockSvc.EXPECT().PatchPost(gomock.Any(), gomock.Any()).Return(rets["PatchPost"]...).Times(times)
		mockSvc.EXPECT().GetAllPosts(gomock.Any()).Return(rets["GetAllPosts"]...).Times(times)
	}

	logger := log.NewNopLogger()
	endpoints := endpoint.New(mockSvc, logger)
	server := httptest.NewServer(NewHTTPHandler(endpoints, logger))
	defer server.Close()
	client := &http.Client{}

	// Login
	accessToken := auth.Login()
	// testcases
	tests := map[string][]struct {
		name             string
		path             string
		method           string
		body             string
		expectedCode     int
		expectedRespBody string
	}{
		"No Error": {
			{"NewPost", "/posts/", "POST", "{}", http.StatusOK, "{\"id\":1}"},
			{"DeletePost", "/posts/1", "DELETE", "", http.StatusOK, "{}"},
			{"GetPost", "/posts/1", "GET", "", http.StatusOK, "{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"site_id\":0,\"user_id\":\"\",\"filename\":\"\",\"filetype\":\"\",\"title\":\"\",\"date\":\"\",\"draft\":false,\"tags\":\"\",\"summary\":\"\",\"content\":\"\"}"},
			{"PatchPost", "/posts/1", "PATCH", "{}", http.StatusOK, "{}"},
			{"GetAllPosts", "/posts/", "GET", "", http.StatusOK, "{\"post_ids\":[]}"},
		},
		"Dummy Error": {
			{"NewPost", "/posts/", "POST", "{}", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"DeletePost", "/posts/1", "DELETE", "", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"GetPost", "/posts/1", "GET", "", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"PatchPost", "/posts/1", "PATCH", "{}", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
			{"GetAllPosts", "/posts/", "GET", "", http.StatusInternalServerError, "{\"error\":\"dummy error\"}"},
		},
	}

	// Start tests
	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			for _, tt := range v {
				t.Run(tt.name, func(t *testing.T) {
					body := []byte(tt.body)
					req, err := http.NewRequest(tt.method, server.URL+tt.path, bytes.NewBuffer(body))
					if err != nil {
						t.Fatalf("Error in creating %s to %s: %v", tt.method, tt.path, err)
					}
					req.Header.Add("Content-Type", "application/json")
					req.Header.Add("Authorization", "Bearer "+accessToken)
					resp, err := client.Do(req)
					if err != nil {
						t.Fatalf("Error in %s to %s: %v", tt.method, tt.path, err)
					}
					defer resp.Body.Close()
					if resp.StatusCode != tt.expectedCode {
						t.Fatalf("Unexpected status code: %d\n Status code should be %d", resp.StatusCode, http.StatusForbidden)
					}
					payload, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						t.Fatalf("Error in reading response body: %v", err)
					}
					respBody := strings.Trim(string(payload), "\n")
					if respBody != tt.expectedRespBody {
						t.Fatalf("Unexpected response body\nwant:\n%s\nget:\n%s\n", tt.expectedRespBody, respBody)
					}
				})
			}
		})
	}
}
