package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Page will all informetion of Page
type Page struct {
	gorm.Model

	Title     string `gorm:"size:128,unique,uniqueIndex:unPage"`
	Content   string `gorm:"type:text"`
	Status    bool
	Weightage int
	AccountID uint
	CreatedBy uint
	UpdatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Page{})
}

// Update will update product by given post argumnets
func (page *Page) Update(PageMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := PageMap["Title"]; ok {
		Value := util.GetString(PageMap["Title"])
		if Value != page.Title {
			page.Title = Value
			flag = true
		}
	}

	if _, ok := PageMap["Content"]; ok {
		Value := util.GetString(PageMap["Content"])
		if Value != page.Content {
			page.Content = Value
			flag = true
		}
	}

	if _, ok := PageMap["Status"]; ok {
		Value := util.GetBool(PageMap["Status"])
		if Value != page.Status {
			page.Status = Value
			flag = true
		}
	}

	if _, ok := PageMap["Weightage"]; ok {
		Value := util.GetInt(PageMap["Weightage"])
		if Value != page.Weightage {
			page.Weightage = Value
			flag = true
		}
	}

	if flag {
		if page.Title == "" {
			err = errors.New(" Title can not be empty")
		} else if page.Content == "" {
			err = errors.New(" Content can not be empty")
		} else if page.AccountID == 0 {
			err = errors.New(" Account can not be empty")
		} else {
			page.UpdatedBy = UpdatedBy.ID
			if page.ID == 0 {
				page.CreatedBy = UpdatedBy.ID
				var duplicates Page
				err = db.Conn.First(&duplicates, " `title` = ? AND `account_id` = ?  ", page.Title, page.AccountID).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&page).Error
				}
			} else {
				err = db.Conn.Save(&page).Error
			}
		}
	}
	return flag, err
}
