package endpoint

import (
	"context"
	"encoding/json"
	"errors"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/config"
	"net/http"
	"time"
)

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// AuthMiddleware returns an endpoint middleware that parse and verify
// a JWT token passed into the context, and that returns errors if the verification failed.
func AuthMiddleware() endpoint.Middleware {
	getPemCert := func(token *stdjwt.Token) (string, error) {
		type JSONWebKeys struct {
			Kty string   `json:"kty"`
			Kid string   `json:"kid"`
			Use string   `json:"use"`
			N   string   `json:"n"`
			E   string   `json:"e"`
			X5c []string `json:"x5c"`
		}

		type Jwks struct {
			Keys []JSONWebKeys `json:"keys"`
		}

		cert := ""
		resp, err := http.Get(config.Auth0Domain + "/.well-known/jwks.json")

		if err != nil {
			return cert, err
		}
		defer resp.Body.Close()

		var jwks = Jwks{}
		err = json.NewDecoder(resp.Body).Decode(&jwks)

		if err != nil {
			return cert, err
		}

		x5c := jwks.Keys[0].X5c
		for k, v := range x5c {
			if token.Header["kid"] == jwks.Keys[k].Kid {
				cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
			}
		}

		if cert == "" {
			err := errors.New("unable to find appropriate key")
			return cert, err
		}

		return cert, nil
	}

	keyFunc := func(token *stdjwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		aud := config.Auth0Audience
		checkAud := token.Claims.(stdjwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}
		// Verify 'iss' claim
		iss := config.Auth0Domain + "/"
		checkIss := token.Claims.(stdjwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := stdjwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}

	return jwt.NewParser(keyFunc, stdjwt.SigningMethodRS256, jwt.MapClaimsFactory)
}

// Middlewares chains all middlewares together, returning the final endpoint.
// This is just a convenient method that helps in clearing up codes in endpoints.New
func Middlewares(e endpoint.Endpoint, logger log.Logger) endpoint.Endpoint {
	chain := endpoint.Chain(AuthMiddleware(), LoggingMiddleware(logger))
	return chain(e)
}
