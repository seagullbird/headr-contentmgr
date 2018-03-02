package service

import (
	"encoding/json"
	"github.com/seagullbird/headr-contentmgr/db"
	"strings"
)

type FrontMatter struct {
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Draft bool     `json:"draft"`
	Tags  []string `json:"tags"`
}

type Post struct {
	Id       int         `json:"id"`
	Author   string      `json:"author"`
	Sitename string      `json:"sitename"`
	Filename string      `json:"filename"`
	Filetype string      `json:"filetype"`
	FM       FrontMatter `json:"fm`
	Summary  string      `json:"summary"`
	Content  string      `json:"content" gorm:"-"`
}

func (fm FrontMatter) String() string {
	fmRaw, _ := json.MarshalIndent(&fm, "", "  ")
	return string(fmRaw) + "\n"
}

func (p Post) String() string {
	lines := []string{
		p.FM.String(),
		p.Summary,
		"<!--more-->",
		p.Content,
	}
	return strings.Join(lines, "\n")
}

func (p Post) Model() *db.Post {
	return &db.Post{
		Id:       p.Id,
		Author:   p.Author,
		Sitename: p.Sitename,
		Filename: p.Filename,
		Filetype: p.Filetype,
		Title:    p.FM.Title,
		Date:     p.FM.Date,
		Draft:    p.FM.Draft,
		Tags:     strings.Join(p.FM.Tags, " "),
		Summary:  p.Summary,
	}
}
