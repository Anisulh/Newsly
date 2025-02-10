package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email             string     `gorm:"uniqueIndex;not null"` // User email
	Username          string     `gorm:"uniqueIndex;not null"` // Unique username
	Password          string     `gorm:"not null"`
	CategoryInterests []Category `gorm:"many2many:user_categories;"`
	Likes             []Like
	Comments          []Comment
	SavedPapers       []SavedPaper
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
