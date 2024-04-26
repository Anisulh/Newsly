package handlers

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
	JWTSecret string
}

func NewHandler(db *gorm.DB, secret string) *Handler {
	return &Handler{DB: db, JWTSecret: secret}
}
