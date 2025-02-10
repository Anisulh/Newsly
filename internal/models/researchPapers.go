package models

import (
	"time"

	"gorm.io/gorm"
)

type ResearchPaper struct {
	gorm.Model
	Title       string    `gorm:"not null"`           // Paper title
	Abstract    string    `gorm:"type:text;not null"` // Paper abstract
	Authors     string    // Could be a comma-separated list; for more complex needs, consider a separate Author model.
	URL         string    // URL pointing to the full paper or external source.
	PublishedAt time.Time // Publication date
	// Many-to-many: A paper can be tagged with multiple categories.
	Categories []Category `gorm:"many2many:research_paper_categories;"`
	// One-to-many relations for interactions.
	Likes       []Like
	Comments    []Comment
	SavedPapers []SavedPaper
}
