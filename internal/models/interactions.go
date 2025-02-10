package models

import "gorm.io/gorm"

// Like represents a user's like on a research paper.
type Like struct {
	gorm.Model
	UserID          uint `gorm:"index"` // Foreign key to User
	ResearchPaperID uint `gorm:"index"` // Foreign key to ResearchPaper
}

// Comment represents a comment made by a user on a research paper.
type Comment struct {
	gorm.Model
	UserID          uint   `gorm:"index"`              // Foreign key to User
	ResearchPaperID uint   `gorm:"index"`              // Foreign key to ResearchPaper
	Content         string `gorm:"type:text;not null"` // The text of the comment
}

// SavedPaper represents a research paper saved/bookmarked by a user.
type SavedPaper struct {
	gorm.Model
	UserID          uint `gorm:"index"` // Foreign key to User
	ResearchPaperID uint `gorm:"index"` // Foreign key to ResearchPaper
}
