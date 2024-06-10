package infra

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// use this struct instead of gorm.Model
type Base struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
}
