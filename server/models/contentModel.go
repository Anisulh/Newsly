package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Title       string
	Description string
	URL         string
	ImageURL    string
	PublishedAt string
	Source      string
	CategoryID  uint
	Category    string //Category
	Keywords		[]string
	Likes       []User `gorm:"many2many:user_likes;"`
	Dislikes    []User `gorm:"many2many:user_dislikes;"`
}
