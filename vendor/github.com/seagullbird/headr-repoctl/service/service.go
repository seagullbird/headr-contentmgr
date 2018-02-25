package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/dispatch"
	"github.com/seagullbird/headr-repoctl/config"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Service interface {
	NewSite(ctx context.Context, email, sitename string) error
	DeleteSite(ctx context.Context, email, sitename string) error
}

func New(dispatcher dispatch.Dispatcher, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(dispatcher)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	dispatcher dispatch.Dispatcher
}

func NewBasicService(dispatcher dispatch.Dispatcher) basicService {
	return basicService{
		dispatcher: dispatcher,
	}
}

func (s basicService) NewSite(ctx context.Context, email, sitename string) error {
	evt := mq.NewSiteEvent{
		email,
		sitename,
		time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("new_site", evt)
}

func (s basicService) DeleteSite(ctx context.Context, email, sitename string) error {
	sitepath := filepath.Join(config.SITESDIR, email, sitename)
	if _, err := os.Stat(sitepath); err != nil {
		if os.IsNotExist(err) {
			return MakeErrPathNotExist(sitepath)
		}
		return MakeErrUnexpected(err)
	}
	cmd := exec.Command("rm", "-rf", sitepath)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
