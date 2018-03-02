package db

import "github.com/jinzhu/gorm"

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
}
