package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// use this struct instead of gorm.Model
type model struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
}

type Todo struct {
	Default model  `gorm:"embedded"`
	Title   string `gorm:"title"`
	Status  string `gorm:"status"`
}
