package tests

import (
	"github.com/seagullbird/headr-repoctl/service"
	"testing"
)

// ReadPost executes RPC call ReadPost to test if previously written file can be corretly read and returned.
func ReadPost(t *testing.T, client *service.Service, p Post) {
	c := *client
	content, err := c.ReadPost(p.ctx, p.siteID, p.filename)
	if err != nil {
		t.Fatal(err)
	}
	if content != p.content {
		t.Fatal("Error: Read content does not match content written.")
	}
}
