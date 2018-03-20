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
	"strconv"
	"time"
)

// Service describes a service that deals with local files in the persistent volume (repoctl).
type Service interface {
	NewSite(ctx context.Context, siteID uint) error
	DeleteSite(ctx context.Context, siteID uint) error
	WritePost(ctx context.Context, siteID uint, filename, content string) error
	RemovePost(ctx context.Context, siteID uint, filename string) error
	ReadPost(ctx context.Context, siteID uint, filename string) (content string, err error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(dispatcher dispatch.Dispatcher, logger log.Logger) Service {
	var svc Service
	{
		svc = newBasicService(dispatcher)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	dispatcher dispatch.Dispatcher
}

func newBasicService(dispatcher dispatch.Dispatcher) basicService {
	return basicService{
		dispatcher: dispatcher,
	}
}

func (s basicService) NewSite(ctx context.Context, siteID uint) error {
	evt := mq.SiteUpdatedEvent{
		SiteId:     siteID,
		Theme:      config.InitialTheme,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("new_site", evt)
}

func (s basicService) DeleteSite(ctx context.Context, siteID uint) error {
	sitepath := filepath.Join(config.SITESDIR, strconv.Itoa(int(siteID)))
	if _, err := os.Stat(sitepath); err != nil {
		if os.IsNotExist(err) {
			return MakeErrPathNotExist(sitepath)
		}
		return MakeErrUnexpected(err)
	}
	cmd := exec.Command("rm", "-rf", sitepath)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s basicService) WritePost(ctx context.Context, siteID uint, filename, content string) error {
	postsPath := filepath.Join(config.SITESDIR, strconv.Itoa(int(siteID)), "source", "content", "posts")
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
		SiteId:     siteID,
		Theme:      config.InitialTheme,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("re_generate", evt)
}

func (s basicService) RemovePost(ctx context.Context, siteID uint, filename string) error {
	postPath := filepath.Join(config.SITESDIR, strconv.Itoa(int(siteID)), "source", "content", "posts", filename)
	if _, err := os.Stat(postPath); err != nil {
		if os.IsNotExist(err) {
			return MakeErrPathNotExist(postPath)
		}
		return MakeErrUnexpected(err)
	}
	cmd := exec.Command("rm", postPath)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	// Generate site
	evt := mq.SiteUpdatedEvent{
		SiteId:     siteID,
		Theme:      config.InitialTheme,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("re_generate", evt)
}

func (s basicService) ReadPost(ctx context.Context, siteID uint, filename string) (content string, err error) {
	postPath := filepath.Join(config.SITESDIR, strconv.Itoa(int(siteID)), "source", "content", "posts", filename)
	if _, err := os.Stat(postPath); err != nil {
		if os.IsNotExist(err) {
			return "", MakeErrPathNotExist(postPath)
		}
		return "", MakeErrUnexpected(err)
	}
	contentRaw, err := ioutil.ReadFile(postPath)
	if err != nil {
		return "", MakeErrUnexpected(err)
	}
	return string(contentRaw), nil
}
