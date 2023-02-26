package models

import "gorm.io/gorm"

// 博客
type Post struct {
	gorm.Model
	Title string
	Body  string
}
