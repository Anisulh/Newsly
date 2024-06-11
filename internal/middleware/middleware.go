package middleware

import "gorm.io/gorm"

type Middleware struct {
	DB *gorm.DB
	JWTSecret string
	Environment string
}

func NewMiddleware(db *gorm.DB, secret string, env string) *Middleware {
	return &Middleware{DB: db, JWTSecret: secret, Environment: env}
}