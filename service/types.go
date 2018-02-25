package service

import (
	"strconv"
	"strings"
)

type FrontMatter struct {
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Draft bool     `json:"draft"`
	Tags  []string `json:"tags"`
}

type Post struct {
	Author   string      `json:"author"`
	Sitename string      `json:"sitename"`
	Filename string      `json:"filename"`
	Filetype string      `json:"filetype"`
	FM       FrontMatter `json:"fm"`
	Summary  string      `json:"summary"`
	Content  string      `json:"content"`
}

func (fm FrontMatter) String() string {
	// TODO: Replace this with yaml encoding lib
	lines := []string{
		"---",
		"title: " + fm.Title,
		"date: " + fm.Date,
		"draft: " + strconv.FormatBool(fm.Draft),
		"tags: [" + strings.Join(fm.Tags, ", ") + "]",
		"---",
	}
	return strings.Join(lines, "\n")
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
