package model

import (
	"cubizy/plugins/db"

	"gorm.io/gorm"
)

// Wishlist will all informetion of Wishlist
type Wishlist struct {
	gorm.Model
	ProductID uint
	CreatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Wishlist{})
}
