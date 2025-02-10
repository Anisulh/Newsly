package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UserID uint
	Key    string `gorm:"uniqueIndex;not null"`
}
