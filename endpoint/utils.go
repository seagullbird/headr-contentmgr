package endpoint

import (
	"github.com/seagullbird/headr-contentmgr/service"
	"time"
)

func (req NewPostRequest) toPost() service.Post {
	fm := service.FrontMatter{
		Title: req.Title,
		Date:  time.Now().Format("2018-02-14T14:55:42+08:00"),
		Draft: false,
		Tags:  req.Tags,
	}
	return service.Post{
		Author:   req.Author,
		Sitename: req.Sitename,
		Filename: postFileNameFormat(req.Title),
		Filetype: "md",
		FM:       fm,
		Summary:  req.Summary,
		Content:  req.Content,
	}
}

func postFileNameFormat(title string) string {
	return title
}
