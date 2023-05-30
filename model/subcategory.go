package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Subcategory will all informetion of Subcategory
type Subcategory struct {
	gorm.Model
	CategoryID          uint
	Name                string
	Description         string
	Logo                string
	Active              bool
	Childcategories     int64
	Products            int64
	Revenue             int64
	CreatedBy           uint
	UpdatedBy           uint
	AccountID           uint
	BaseSubcategoryID   uint
	BaseChildcategoryID uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Subcategory{})
	db.Conn.Exec("UPDATE subcategories SET subcategories.childcategories = (SELECT count(childcategories.id) FROM childcategories WHERE childcategories.subcategory_id = subcategories.id && childcategories.deleted_at IS NULL)")
	db.Conn.Exec("UPDATE subcategories SET subcategories.products = (SELECT count(products.id) FROM products WHERE products.base_subcategory_id = subcategories.id)")
}

// Update will update product by given post argumnets
func (subcategory *Subcategory) Update(SubcategoryMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := SubcategoryMap["CategoryID"]; ok {
		Value := util.GetUint(SubcategoryMap["CategoryID"])
		if Value != subcategory.CategoryID {
			subcategory.CategoryID = Value
			flag = true
		}
	}

	if _, ok := SubcategoryMap["Name"]; ok {
		Value := util.GetString(SubcategoryMap["Name"])
		if Value != subcategory.Name {
			subcategory.Name = Value
			flag = true
		}
	}

	if _, ok := SubcategoryMap["Description"]; ok {
		Value := util.GetString(SubcategoryMap["Description"])
		if Value != subcategory.Description {
			subcategory.Description = Value
			flag = true
		}
	}

	if _, ok := SubcategoryMap["Logo"]; ok {
		Value := util.GetString(SubcategoryMap["Logo"])
		if Value != subcategory.Logo {
			subcategory.Logo = Value
			flag = true
		}
	}

	if _, ok := SubcategoryMap["Active"]; ok {
		Value := util.GetBool(SubcategoryMap["Active"])
		if Value != subcategory.Active {
			subcategory.Active = Value
			flag = true
		}
	}

	if _, ok := SubcategoryMap["BaseSubcategoryID"]; ok {
		Value := util.GetUint(SubcategoryMap["BaseSubcategoryID"])
		if Value != subcategory.BaseSubcategoryID {
			subcategory.BaseSubcategoryID = Value
			flag = true
		}
	}

	if flag {
		if subcategory.Name == "" {
			err = errors.New(" Name can not be empty")
		} else if subcategory.AccountID == 0 {
			err = errors.New(" Account can not be empty")
		} else if subcategory.CategoryID == 0 {
			err = errors.New(" Category can not be empty")
		} else if !UpdatedBy.IsAdmin && subcategory.BaseSubcategoryID == 0 {
			err = errors.New(" Base Subcategory can not be empty")
		} else {
			subcategory.UpdatedBy = UpdatedBy.ID
			if subcategory.ID == 0 {
				subcategory.CreatedBy = UpdatedBy.ID
				var duplicates Subcategory
				err = db.Conn.First(&duplicates, " `name` = ? && `account_id` = ? ", subcategory.Name, subcategory.AccountID).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&subcategory).Error
				}
				UpdateSubcategoryCount(subcategory.CategoryID)
			} else {
				err = db.Conn.Save(&subcategory).Error
			}
		}
	}
	return flag, err
}

func UpdateChildcategoryCount(Subcategory_ID uint) {
	err := db.Conn.Exec("UPDATE subcategories SET subcategories.childcategories = (SELECT count(childcategories.id) FROM childcategories WHERE childcategories.subcategory_id = subcategories.id && childcategories.deleted_at IS NULL) WHERE subcategories.id = ? ", Subcategory_ID).Error
	if err != nil {
		util.Log("UpdateChildcategoryCount Filed", err.Error())
	}
}
