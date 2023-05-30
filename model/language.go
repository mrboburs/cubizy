package model

import (
	"gorm.io/gorm"
)

// Language hold Language key values in database
type Language struct {
	gorm.Model
	Key       string `gorm:"type:varchar(100);UNIQUE"`
	EnValue   string `gorm:"type:TEXT;"`
	Domains   string `gorm:"type:TEXT;"`
	UpdatedBy uint
}
