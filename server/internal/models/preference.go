package models

import "gorm.io/gorm"

type Preference struct {
	gorm.Model
	UserID  uint
	Content string // Could be categories, specific tags, or other preference indicators
}
