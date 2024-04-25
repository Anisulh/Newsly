package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Title       string
	Description string
	Content     string
	URL         string `gorm:"uniqueIndex"`
	ImageURL    string
	PublishedAt string
	Source      string
	Category    Category
	Keywords    []Keyword `gorm:"many2many:content_keywords;"`
	Likes       []*User   `gorm:"many2many:user_likes;"`
	Dislikes    []*User   `gorm:"many2many:user_dislikes;"`
}

// Category is an enumerated type for content categories.
type Category string

const (
	World         Category = "World"
	Politics      Category = "Politics"
	Technology    Category = "Technology"
	Science       Category = "Science"
	Entertainment Category = "Entertainment"
	Business      Category = "Business"
	Sports        Category = "Sports"
)
