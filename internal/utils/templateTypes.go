package utils

import (
	"Newsly/internal/models"
)

type BaseData struct {
	IsAuth      bool
	Account     *models.User
	Message     string
	MessageType string
	Articles    []Article
}

type Notification struct {
	Message     string
	MessageType string
}
