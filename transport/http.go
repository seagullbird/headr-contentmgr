package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/seagullbird/headr-contentmgr/endpoint"
	"net/http"
	"strconv"
)

type errorWrapper struct {
	Error string `json:"error"`
}

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(jwt.HTTPToContext()),
	}
	r := mux.NewRouter()

	// POST 	/posts/			add a post
	// DELETE	/posts/:id		remove the given post
	// GET    	/posts/:id	 	retrieve the given post by id
	// GET	    /posts/         retrieve all posts of the authenticated user
	// PATCH    /posts/:id		partial update a post

	r.Methods("POST").Path("/posts/").Handler(httptransport.NewServer(
		endpoints.NewPostEndpoint,
		decodeHTTPNewPostRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("DELETE").Path("/posts/{id}").Handler(httptransport.NewServer(
		endpoints.DeletePostEndpoint,
		decodeHTTPDeletePostRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("GET").Path("/posts/{id}").Handler(httptransport.NewServer(
		endpoints.GetPostEndpoint,
		decodeHTTPGetPostRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("GET").Path("/posts/").Handler(httptransport.NewServer(
		endpoints.GetAllPostsEndpoint,
		decodeHTTPGetAllPostsRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("PATCH").Path("/posts/{id}").Handler(httptransport.NewServer(
		endpoints.PatchPostEndpoint,
		decodeHTTPPatchPostRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	return r
}

func err2code(err error) int {
	switch err {
	case jwt.ErrTokenContextMissing, jwt.ErrTokenExpired, jwt.ErrTokenInvalid, jwt.ErrTokenMalformed, jwt.ErrTokenNotActive, jwt.ErrUnexpectedSigningMethod:
		return http.StatusForbidden
	case ErrBadRouting:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func decodeHTTPNewPostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.NewPostRequest
	err := json.NewDecoder(r.Body).Decode(&req.Post)
	return req, err
}

func decodeHTTPDeletePostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRouting
	}
	return endpoint.DeletePostRequest{ID: uint(i)}, nil
}

func decodeHTTPGetPostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRouting
	}
	return endpoint.GetPostRequest{ID: uint(i)}, nil
}

func decodeHTTPGetAllPostsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoint.GetAllPostsRequest{}, nil
}

func decodeHTTPPatchPostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRouting
	}
	var req endpoint.PatchPostRequest
	err = json.NewDecoder(r.Body).Decode(&req.Post)
	req.Post.ID = uint(i)
	return req, err
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}
