package model

import (
	"gorm.io/gorm"
	"time"
)

type Text struct {
	gorm.Model
	TextHash   string    `gorm:"uniqueIndex;not null;size:30" json:"id"`
	Text       string    `json:"text"`
	TextLink   string    `gorm:"uniqueIndex;not null;size:255" json:"text_link"`
	TextExpire time.Time `gorm:"not null" json:"text_expire"`
}
