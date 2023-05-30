package model

import (
	"cubizy/plugins/db"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Title         string
	Image         string
	Content       string
	Type          int
	RecivedAt     int64
	SeenAt        int64
	ReplayTo      uint
	UpdatedBy     uint
	CreatedBy     uint
	CreatedFor    uint
	MessageID     uint
	TotalMessages uint
	TotalUsers    uint
	Ispublic      bool
}

type ChatUsers struct {
	ChatID uint
	UserID uint
}

type ChatListView struct {
	gorm.Model
	ChatID         string
	Name           string
	Photo          string
	Online         bool
	LastActiveOn   int64
	Content        string
	Type           int
	RecivedAt      int64
	SeenAt         int64
	ReplayTo       uint
	UpdatedBy      uint
	CreatedBy      uint
	CreatedFor     uint
	MessageID      uint
	TotalMessages  uint
	TotalUsers     uint
	Ispublic       bool
	IsSupportagent bool
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Chat{})
	db.Conn.AutoMigrate(&ChatUsers{})
}
