package models

import "gorm.io/gorm"

type Mensagem struct {
	gorm.Model
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Texto string `json:"mensagem" gorm:"type:text"`
}
