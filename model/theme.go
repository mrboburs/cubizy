package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Theme will all informetion of Theme
type Theme struct {
	gorm.Model

	Title                string `gorm:"size:128,unique,uniqueIndex:unTheme"`
	ThemeID              uint
	BasePublishedThemeID uint
	PublishedThemeID     uint
	Status               string
	SubmitedOn           int64
	PublishedOn          int64
	Logo                 string
	Tags                 string
	Description          string
	Images               string
	IsUpdated            bool
	AccountID            uint
	CreatedBy            uint
	UpdatedBy            uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Theme{})
}

// Update will update product by given post argumnets
func (theme *Theme) Update(ThemeMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 || (theme.AccountID != UpdatedBy.SellerAccountID && !UpdatedBy.IsAdmin) {
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

	if flag {
		if theme.Title == "" {
			err = errors.New(" Title can not be empty")
		} else {
			theme.UpdatedBy = UpdatedBy.ID
			if theme.ID == 0 {
				theme.CreatedBy = UpdatedBy.ID
				var duplicates Theme
				err = db.Conn.First(&duplicates, " `title` = ? AND `account_id` = ? ", theme.Title, theme.AccountID).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&theme).Error
				}
			} else {
				err = db.Conn.Save(&theme).Error
			}
		}
	}
	return flag, err
}
