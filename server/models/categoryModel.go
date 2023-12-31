package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name    string
	Contents []Content `gorm:"foreignKey:CategoryID"`
}
