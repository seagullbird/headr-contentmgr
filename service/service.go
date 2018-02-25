package service

import (
	"context"
	"github.com/go-kit/kit/log"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
)

type Service interface {
	NewPost(ctx context.Context) error
}

func New(repoctlsvc repoctlservice.Service, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(repoctlsvc)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
}

func NewBasicService(repoctlsvc repoctlservice.Service) basicService {
	return basicService{
		repoctlsvc: repoctlsvc,
	}
}

func (s basicService) NewPost(ctx context.Context) error {
	return nil
}
