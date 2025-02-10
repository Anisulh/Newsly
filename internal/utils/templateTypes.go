package utils

import (
	"Newsly/internal/models"
)

type BaseData struct {
	IsAuth      bool
	Account     *models.User
	Message     string
	MessageType string
}

type Category struct {
	Title       string
	Description string
}

type InterestTopicsData struct {
	BaseData
	Categories map[string]Category
}

type FeedData struct {
	BaseData
	Papers []ArxivEntry
}

type Notification struct {
	Message     string
	MessageType string
}
