package tests

import (
	"github.com/seagullbird/headr-repoctl/config"
	"github.com/seagullbird/headr-repoctl/service"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
)

// NewSite executes RPC call NewSite and manually create an empty source site directory (for later tests in main_test.go).
func NewSite(t *testing.T, client *service.Service, s Site) {
	c := *client
	err := c.NewSite(s.ctx, s.siteID)
	if err != nil {
		t.Fatal(err)
	}
	// Mock s source creation locally
	if err := exec.Command("mkdir", "-p", filepath.Join(config.SITESDIR, strconv.Itoa(int(s.siteID)), "source")).Run(); err != nil {
		t.Fatal(err)
	}
}
