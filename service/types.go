package service

import (
	"encoding/json"
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
