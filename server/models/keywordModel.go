package models

import "gorm.io/gorm"

type Keyword struct {
	gorm.Model
	Name     string
	Contents []Content `gorm:"many2many:content_keywords;"`
}
