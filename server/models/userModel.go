package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Email       string `gorm:"unique"`
	Password    string
	Preferences []Preference `gorm:"foreignKey:UserID"`
	Bookmarks   []Content    `gorm:"many2many:user_bookmarks;"`
	Likes       []Content    `gorm:"many2many:user_likes;"`
	Dislikes    []Content    `gorm:"many2many:user_dislikes;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
