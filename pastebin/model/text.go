package model

import (
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	TextHash string `gorm:"uniqueIndex;not null;size:30"`
	FileName string `gorm:"not null;size:30"`
	ExpireAt string `gorm:"not null"`
}
