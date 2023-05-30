package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"time"

	"gorm.io/gorm"
)

// PublishedTheme will all informetion of Published Theme
type PublishedTheme struct {
	gorm.Model

	Title                string `gorm:"size:128,unique,uniqueIndex:unTheme"`
	ThemeID              uint
	BasePublishedThemeID uint
	Logo                 string
	Images               string
	Tags                 string
	Description          string
	Status               string // Fail to Upload, Submitted, Rejected, Accepted, Published
	AccountID            uint
	CreatedBy            uint
	UpdatedBy            uint
	PublishedOn          int64
	SubmitedOn           int64
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&PublishedTheme{})
}

// Update will update product by given post argumnets
func (theme *PublishedTheme) Update(ThemeMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := ThemeMap["Title"]; ok {
		Value := util.GetString(ThemeMap["Title"])
		if Value != theme.Title {
			theme.Title = Value
			flag = true
		}
	}

	if _, ok := ThemeMap["Logo"]; ok {
		Value := util.GetString(ThemeMap["Logo"])
		if Value != theme.Logo {
			theme.Logo = Value
			flag = true
		}
	}

	if _, ok := ThemeMap["Images"]; ok {
		Value := util.GetString(ThemeMap["Images"])
		if Value != theme.Images {
			theme.Images = Value
			flag = true
		}
	}

	if _, ok := ThemeMap["Tags"]; ok {
		Value := util.GetString(ThemeMap["Tags"])
		if Value != theme.Tags {
			theme.Tags = Value
			flag = true
		}
	}

	if _, ok := ThemeMap["Description"]; ok {
		Value := util.GetString(ThemeMap["Description"])
		if Value != theme.Description {
			theme.Description = Value
			flag = true
		}
	}

	var isStatusChanged = false
	if UpdatedBy.IsAdmin {
		if _, ok := ThemeMap["Status"]; ok {
			Value := util.GetString(ThemeMap["Status"])
			if Value != theme.Status {
				theme.Status = Value
				if theme.Status == "Published" {
					theme.PublishedOn = time.Now().Unix()
				}
				isStatusChanged = true
				flag = true
			}
		}
	}

	if flag {
		if theme.Title == "" {
			err = errors.New(" Title can not be empty")
		} else {
			theme.UpdatedBy = UpdatedBy.ID
			if theme.ID == 0 {
				theme.CreatedBy = UpdatedBy.ID
				var duplicates PublishedTheme
				err = db.Conn.First(&duplicates, " `title` = ? ", theme.Title, theme.AccountID).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&theme).Error
				}
			} else {
				err = db.Conn.Save(&theme).Error
			}
			if err == nil && isStatusChanged {
				var from_theme Theme
				from_theme_err := db.Conn.First(&from_theme, "id = ?", theme.ThemeID).Error
				if from_theme_err == nil {
					from_theme.PublishedOn = theme.PublishedOn
					from_theme.Status = theme.Status
					err = db.Conn.Save(&from_theme).Error
				}
			}
		}
	}
	return flag, err
}
