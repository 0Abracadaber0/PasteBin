package model

import (
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	TextHash   string `gorm:"uniqueIndex;not null;size:30"`
	TextExpire string `gorm:"not null"`
}
