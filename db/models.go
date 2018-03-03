package db

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strings"
)

type Post struct {
	gorm.Model
	Author   string `json:"author"`
	Sitename string `json:"sitename"`
	Filename string `json:"filename"`
	Filetype string `json:"filetype"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Draft    bool   `json:"draft"`
	Tags     string `json:"tags"`
	Summary  string `json:"summary"`
	Content  string `json:"content" gorm:"-"`
}

func (p Post) String() string {
	type FrontMatter struct {
		Title string   `json:"title"`
		Date  string   `json:"date"`
		Draft bool     `json:"draft"`
		Tags  []string `json:"tags"`
	}

	fm := FrontMatter{
		Title: p.Title,
		Date:  p.Date,
		Draft: p.Draft,
		Tags:  strings.Split(p.Tags, " "),
	}
	fmsRaw, _ := json.MarshalIndent(&fm, "", "  ")
	fms := string(fmsRaw) + "\n"
	lines := []string{
		fms,
		p.Summary,
		"<!--more-->",
		p.Content,
	}
	return strings.Join(lines, "\n")
}
