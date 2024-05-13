package utils

import "Newsly/internal/models"



type BaseData struct {
  IsAuth  bool        
  Account *models.User
}