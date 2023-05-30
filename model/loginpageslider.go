package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Loginpageslider will all informetion of Loginpageslider
type Loginpageslider struct {
	gorm.Model

	Image   string
	Title   string `gorm:"size:128,unique,uniqueIndex:unLoginpageslider"`
	Content string `gorm:"type:text"`
	Footer  string `gorm:"size:128,unique,uniqueIndex:unLoginpageslider"`
	Status  bool

	AccountID uint
	CreatedBy uint
	UpdatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Loginpageslider{})
}

// Update will update product by given post argumnets
func (loginpageslider *Loginpageslider) Update(LoginpagesliderMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := LoginpagesliderMap["Title"]; ok {
		Value := util.GetString(LoginpagesliderMap["Title"])
		if Value != loginpageslider.Title {
			loginpageslider.Title = Value
			flag = true
		}
	}

	if _, ok := LoginpagesliderMap["Image"]; ok {
		Value := util.GetString(LoginpagesliderMap["Image"])
		if Value != loginpageslider.Image {
			loginpageslider.Image = Value
			flag = true
		}
	}

	if _, ok := LoginpagesliderMap["Content"]; ok {
		Value := util.GetString(LoginpagesliderMap["Content"])
		if Value != loginpageslider.Content {
			loginpageslider.Content = Value
			flag = true
		}
	}

	if _, ok := LoginpagesliderMap["Footer"]; ok {
		Value := util.GetString(LoginpagesliderMap["Footer"])
		if Value != loginpageslider.Footer {
			loginpageslider.Footer = Value
			flag = true
		}
	}

	if _, ok := LoginpagesliderMap["Status"]; ok {
		Value := util.GetBool(LoginpagesliderMap["Status"])
		if Value != loginpageslider.Status {
			loginpageslider.Status = Value
			flag = true
		}
	}

	if flag {
		if loginpageslider.Title == "" {
			err = errors.New(" Title can not be empty")
		} else if loginpageslider.Image == "" {
			err = errors.New(" Image can not be empty")
		} else {
			loginpageslider.UpdatedBy = UpdatedBy.ID
			if loginpageslider.ID == 0 {
				loginpageslider.CreatedBy = UpdatedBy.ID
				var duplicates Loginpageslider
				err = db.Conn.First(&duplicates, " `title` = ? ", loginpageslider.Title).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&loginpageslider).Error
				}
			} else {
				err = db.Conn.Save(&loginpageslider).Error
			}
		}
	}
	return flag, err
}
