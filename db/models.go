package db

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strings"
)

// Post is the ORM model for database table Post
type Post struct {
	gorm.Model
	SiteID   uint   `json:"site_id"`
	UserID   string `json:"user_id"`
	Filename string `json:"filename"`
	Filetype string `json:"filetype"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Draft    bool   `json:"draft"`
	Tags     string `json:"tags"`
	Summary  string `json:"summary"`
	Content  string `json:"content" gorm:"-"`
}

// String convert a Post model to its literal form
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
