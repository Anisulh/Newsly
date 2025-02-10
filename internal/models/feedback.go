package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	UserID    uint
	ContentID uint
	Feedback  string // e.g., "positive", "negative", "neutral"
}
