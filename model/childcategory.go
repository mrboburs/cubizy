package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Childcategory will all informetion of Childcategory
type Childcategory struct {
	gorm.Model
	SubcategoryID       uint
	Name                string
	Description         string
	Logo                string
	Active              bool
	Products            int64
	Revenue             int64
	CreatedBy           uint
	UpdatedBy           uint
	AccountID           uint
	BaseChildcategoryID uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Childcategory{})
	db.Conn.Exec("UPDATE childcategories SET childcategories.products = (SELECT count(products.id) FROM products WHERE products.base_childcategory_id = childcategories.id)")
}

// Update will update product by given post argumnets
func (childcategory *Childcategory) Update(ChildcategoryMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := ChildcategoryMap["SubcategoryID"]; ok {
		Value := util.GetUint(ChildcategoryMap["SubcategoryID"])
		if Value != childcategory.SubcategoryID {
			childcategory.SubcategoryID = Value
			flag = true
		}
	}

	if _, ok := ChildcategoryMap["Name"]; ok {
		Value := util.GetString(ChildcategoryMap["Name"])
		if Value != childcategory.Name {
			childcategory.Name = Value
			flag = true
		}
	}

	if _, ok := ChildcategoryMap["Description"]; ok {
		Value := util.GetString(ChildcategoryMap["Description"])
		if Value != childcategory.Description {
			childcategory.Description = Value
			flag = true
		}
	}

	if _, ok := ChildcategoryMap["Logo"]; ok {
		Value := util.GetString(ChildcategoryMap["Logo"])
		if Value != childcategory.Logo {
			childcategory.Logo = Value
			flag = true
		}
	}

	if _, ok := ChildcategoryMap["Active"]; ok {
		Value := util.GetBool(ChildcategoryMap["Active"])
		if Value != childcategory.Active {
			childcategory.Active = Value
			flag = true
		}
	}

	if _, ok := ChildcategoryMap["BaseChildcategoryID"]; ok {
		Value := util.GetUint(ChildcategoryMap["BaseChildcategoryID"])
		if Value != childcategory.BaseChildcategoryID {
			childcategory.BaseChildcategoryID = Value
			flag = true
		}
	}

	if flag {
		if childcategory.Name == "" {
			err = errors.New(" Name can not be empty")
		} else if childcategory.AccountID == 0 {
			err = errors.New(" Account can not be empty")
		} else if childcategory.SubcategoryID == 0 {
			err = errors.New(" Subcategory can not be empty")
		} else if !UpdatedBy.IsAdmin && childcategory.BaseChildcategoryID == 0 {
			err = errors.New(" Cubizy Childcategory can not be empty")
		} else {
			childcategory.UpdatedBy = UpdatedBy.ID
			if childcategory.ID == 0 {
				childcategory.CreatedBy = UpdatedBy.ID
				var duplicates Childcategory
				err = db.Conn.First(&duplicates, " `name` = ? && `account_id` = ? ", childcategory.Name, childcategory.AccountID).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&childcategory).Error
				}
				UpdateChildcategoryCount(childcategory.SubcategoryID)
			} else {
				err = db.Conn.Save(&childcategory).Error
			}
		}
	}
	return flag, err
}
