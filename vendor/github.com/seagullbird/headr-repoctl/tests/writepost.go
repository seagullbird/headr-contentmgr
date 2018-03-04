package tests

import (
	"github.com/seagullbird/headr-repoctl/config"
	"github.com/seagullbird/headr-repoctl/service"
	"io/ioutil"
	"path/filepath"
	"testing"
)

// WritePost executes RPC call WritePost and check if file is correctly written.
func WritePost(t *testing.T, client *service.Service, p Post) {
	c := *client
	if err := c.WritePost(p.ctx, p.author, p.sitename, p.filename, p.content); err != nil {
		t.Fatal(err)
	}
	postPath := filepath.Join(config.SITESDIR, p.author, p.sitename, "source", "content", "posts", p.filename)
	contentRaw, err := ioutil.ReadFile(postPath)
	if err != nil {
		t.Fatal(err)
	}
	if p.content != string(contentRaw) {
		t.Fatal("Error: Writed content does not match content to write.")
	}
}
