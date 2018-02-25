package endpoint

import (
	"github.com/seagullbird/headr-contentmgr/service"
)

func (req NewPostRequest) toPost() service.Post {
	fm := service.FrontMatter{
		Title: req.Title,
		Date:  req.Date,
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
