package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Category will all informetion of Category
type Category struct {
	gorm.Model
	Name           string
	Description    string
	Logo           string
	Icon           string
	Active         bool
	Subcategories  int64
	Products       int64
	Top            int64
	Featured       int64
	Revenue        int64
	CreatedBy      uint
	UpdatedBy      uint
	AccountID      uint
	BaseCategoryID uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Category{})
	db.Conn.Exec("UPDATE categories SET categories.subcategories = (SELECT count(subcategories.id) FROM subcategories WHERE subcategories.category_id = categories.id && subcategories.deleted_at IS NULL)")
	db.Conn.Exec("UPDATE categories SET categories.products = (SELECT count(products.id) FROM products WHERE products.base_category_id = categories.id)")
}

// Update will update product by given post argumnets
func (category *Category) Update(CategoryMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := CategoryMap["Name"]; ok {
		Value := util.GetString(CategoryMap["Name"])
		if Value != category.Name {
			category.Name = Value
			flag = true
		}
	}

	if _, ok := CategoryMap["Description"]; ok {
		Value := util.GetString(CategoryMap["Description"])
		if Value != category.Description {
			category.Description = Value
			flag = true
		}
	}

	if _, ok := CategoryMap["Logo"]; ok {
		Value := util.GetString(CategoryMap["Logo"])
		if Value != category.Logo {
			category.Logo = Value
			flag = true
		}
	}

	if _, ok := CategoryMap["Icon"]; ok {
		Value := util.GetString(CategoryMap["Icon"])
		if Value != category.Icon {
			category.Icon = Value
			flag = true
		}
	}

	if _, ok := CategoryMap["Active"]; ok {
		Value := util.GetBool(CategoryMap["Active"])
		if Value != category.Active {
			category.Active = Value
			flag = true
		}
	}

	if _, ok := CategoryMap["Top"]; ok {
		Value := util.GetInt64(CategoryMap["Top"])
		if Value != category.Top {
			category.Top = Value
			flag = true
		}
	}

	if _, ok := CategoryMap["Featured"]; ok {
		Value := util.GetInt64(CategoryMap["Featured"])
		if Value != category.Featured {
			category.Featured = Value
			flag = true
		}
	}

	if _, ok := CategoryMap["BaseCategoryID"]; ok {
		Value := util.GetUint(CategoryMap["BaseCategoryID"])
		if Value != category.BaseCategoryID {
			category.BaseCategoryID = Value
			flag = true
		}
	}

	if flag {
		if category.Name == "" {
			err = errors.New(" Name can not be empty")
		} else if category.AccountID == 0 {
			err = errors.New(" Account can not be empty")
		} else if !UpdatedBy.IsAdmin && category.BaseCategoryID == 0 {
			err = errors.New(" Base Category can not be empty")
		} else {
			category.UpdatedBy = UpdatedBy.ID
			if category.ID == 0 {
				category.CreatedBy = UpdatedBy.ID
				var duplicates Category
				err = db.Conn.First(&duplicates, " `name` = ? && `account_id` = ? ", category.Name, category.AccountID).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&category).Error
				}
			} else {
				err = db.Conn.Save(&category).Error
			}
		}
	}
	return flag, err
}

func UpdateSubcategoryCount(Category_ID uint) {
	err := db.Conn.Exec("UPDATE categories SET categories.subcategories = (SELECT count(subcategories.id) FROM subcategories WHERE subcategories.category_id = categories.id && subcategories.deleted_at IS NULL) WHERE categories.id = ? ", Category_ID).Error
	if err != nil {
		util.Log("UpdateSubcategoryCount Filed", err.Error())
	}
}
