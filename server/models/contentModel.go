package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Title       string
	Description string
	URL         string
	CategoryID  uint
	Category    Category
	Likes       []User `gorm:"many2many:user_likes;"`
	Dislikes    []User `gorm:"many2many:user_dislikes;"`
}
