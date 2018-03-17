package tests

import (
	"github.com/seagullbird/headr-repoctl/config"
	"github.com/seagullbird/headr-repoctl/service"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// RemovePost executes RPC call RemovePost and check if post is actually removed.
func RemovePost(t *testing.T, client *service.Service, p Post) {
	c := *client
	err := c.RemovePost(p.ctx, p.siteID, p.filename)
	if err != nil {
		t.Fatal(err)
	}
	// Make sure file is deleted
	postPath := filepath.Join(config.SITESDIR, strconv.Itoa(int(p.siteID)), "source", "content", "posts", p.filename)
	if _, err := os.Stat(postPath); !os.IsNotExist(err) {
		t.Fatal(err)
	}
}
