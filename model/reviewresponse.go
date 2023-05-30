package model

import (
	"cubizy/plugins/db"

	"gorm.io/gorm"
)

// ReviewResponse will all Response of Review new_cubizy.review_responses
type ReviewResponse struct {
	gorm.Model
	ReviewID  uint
	CreatedBy uint
	Response  string
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&ReviewResponse{})
}
