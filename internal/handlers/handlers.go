package handlers

import "gorm.io/gorm"

type Handler struct {
	DB          *gorm.DB
	JWTSecret   string
	Environment string
	NewsAPIKey  string
}

func NewHandler(db *gorm.DB, secret string, env string, NewsAPIKey string) *Handler {
	return &Handler{DB: db, JWTSecret: secret, Environment: env, NewsAPIKey: NewsAPIKey}
}
