package tests

import (
	"github.com/seagullbird/headr-repoctl/config"
	"github.com/seagullbird/headr-repoctl/service"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// DeleteSite executes RPC call DeleteSite and check if site is successfully deleted.
func DeleteSite(t *testing.T, client *service.Service, s Site) {
	c := *client
	err := c.DeleteSite(s.ctx, s.siteID)
	if err != nil {
		t.Fatal(err)
	}
	// Make sure path does not exist
	sitePath := filepath.Join(config.SITESDIR, strconv.Itoa(int(s.siteID)))
	if _, err := os.Stat(sitePath); !os.IsNotExist(err) {
		t.Fatal(err)
	}
}
