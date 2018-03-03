package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/dispatch"
	"github.com/seagullbird/headr-repoctl/config"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Service interface {
	NewSite(ctx context.Context, email, sitename string) error
	DeleteSite(ctx context.Context, email, sitename string) error
	NewPost(ctx context.Context, author, sitename, filename, content string) error
	RemovePost(ctx context.Context, author, sitename, filename string) error
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
	evt := mq.SiteUpdatedEvent{
		email,
		sitename,
		config.InitialTheme,
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

func (s basicService) NewPost(ctx context.Context, author, sitename, filename, content string) error {
	postsPath := filepath.Join(config.SITESDIR, author, sitename, "source", "content", "posts")
	if _, err := os.Stat(postsPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(postsPath, 0644)
		} else {
			return err
		}
	}
	postPath := filepath.Join(postsPath, filename)
	if err := ioutil.WriteFile(postPath, []byte(content), 0644); err != nil {
		return err
	}
	// Generate site
	evt := mq.SiteUpdatedEvent{
		Email:      author,
		SiteName:   sitename,
		Theme:      config.InitialTheme,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("re_generate", evt)
}

func (s basicService) RemovePost(ctx context.Context, author, sitename, filename string) error {
	postPath := filepath.Join(config.SITESDIR, author, sitename, "source", "content", "posts", filename)
	if _, err := os.Stat(postPath); err != nil {
		if os.IsNotExist(err) {
			return MakeErrPathNotExist(postPath)
		}
		return MakeErrUnexpected(err)
	}
	cmd := exec.Command("rm", postPath)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
