package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// Description will all informetion of Description
type Description struct {
	gorm.Model
	ProductID uint
	Content   string `gorm:"type:LONGBLOB"`
	AccountID uint
	CreatedBy uint
	UpdatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Description{})
}

// Update will update product by given post argumnets
func (description *Description) Update(DescriptionMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := DescriptionMap["Content"]; ok {
		Value := util.GetString(DescriptionMap["Content"])
		if Value != description.Content {
			description.Content = Value
			flag = true
		}
	}

	if flag {
		if description.ProductID == 0 {
			err = errors.New(" Product can not be empty")
		} else if description.AccountID == 0 {
			err = errors.New(" Account can not be empty")
		} else if strings.Contains(description.Content, "<script") {
			err = errors.New(" Can not add any script in this code, only html and css is allowed")
		} else {
			description.UpdatedBy = UpdatedBy.ID
			if description.ID == 0 {
				description.CreatedBy = UpdatedBy.ID
				err = db.Conn.Create(&description).Error
			} else {
				err = db.Conn.Save(&description).Error
			}
		}
	}
	return flag, err
}
