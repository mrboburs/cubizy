package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// attribute will all informetion of attribute of category
type Attribute struct {
	gorm.Model
	Name            string
	SubcategoryID   uint
	ChildcategoryID uint
	FieldType       string
	Options         string
	ProductColumn   string
	CreatedBy       uint
	UpdatedBy       uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Attribute{})
}

// Update will update product by given post argumnets
func (attribute *Attribute) Update(AttributeMap map[string]interface{}, UpdatedBy *User) (bool, error) {

	var err error
	var flag bool
	var subcategory *Subcategory
	var childcategory *Childcategory

	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := AttributeMap["Name"]; ok {
		Value := util.GetString(AttributeMap["Name"])
		if Value != attribute.Name {
			attribute.Name = Value
			flag = true
		}
	}

	if _, ok := AttributeMap["SubcategoryID"]; ok {
		Value := util.GetUint(AttributeMap["SubcategoryID"])
		if Value > 0 && Value != attribute.SubcategoryID {
			err = db.Conn.First(&subcategory, "id = ?", Value).Error
			if err == nil && subcategory.BaseSubcategoryID == 0 && Value != attribute.SubcategoryID {
				attribute.SubcategoryID = Value
				flag = true
			}
		}
	}

	if _, ok := AttributeMap["ChildcategoryID"]; ok {
		Value := util.GetUint(AttributeMap["ChildcategoryID"])
		if Value > 0 && Value != attribute.ChildcategoryID {
			err = db.Conn.First(&childcategory, "id = ?", Value).Error
			if err == nil && childcategory.BaseChildcategoryID == 0 {
				attribute.ChildcategoryID = Value
				flag = true
			}
		}
	}

	if _, ok := AttributeMap["FieldType"]; ok {
		Value := util.GetString(AttributeMap["FieldType"])
		if Value != attribute.FieldType {
			attribute.FieldType = Value
			flag = true
		}
	}

	if _, ok := AttributeMap["Options"]; ok {
		Value := util.GetString(AttributeMap["Options"])
		if Value != attribute.Options {
			attribute.Options = Value
			flag = true
		}
	}

	if flag {
		if attribute.Name == "" {
			err = errors.New(" Name can not be empty")
		} else if attribute.FieldType == "" {
			err = errors.New(" FieldType can not be empty")
		} else if attribute.SubcategoryID == 0 {
			err = errors.New(" Subcategory can not be empty and it must be cubizy category")
		} else {
			attribute.SetDBColumn()
			attribute.UpdatedBy = UpdatedBy.ID
			if attribute.ID == 0 {
				attribute.CreatedBy = UpdatedBy.ID
				var duplicates Attribute
				err = db.Conn.First(&duplicates, " subcategory_id = ? AND childcategory_id = ? AND `name` = ? ", attribute.SubcategoryID, attribute.ChildcategoryID, attribute.Name).Error
				if err == nil {
					err = errors.New("duplicate attribute name")
				} else {
					err = db.Conn.Create(&attribute).Error
				}
			} else {
				err = db.Conn.Save(&attribute).Error
			}
		}
	}
	return flag, err
}

// GetAttributes will give array of attributes by given SubcategoryID and ChildcategoryID
func GetAttributes(SubcategoryID, ChildcategoryID uint) ([]Attribute, error) {
	var err error
	var attributes []Attribute
	if ChildcategoryID > 0 {
		var childcategory Childcategory
		err = db.Conn.First(&childcategory, ChildcategoryID).Error
		if err == nil {
			if childcategory.BaseChildcategoryID > 0 {
				var basechildcategory Childcategory
				err = db.Conn.First(&basechildcategory, childcategory.BaseChildcategoryID).Error
				if err == nil {
					childcategory = basechildcategory
					ChildcategoryID = childcategory.ID
				} else {
					err = errors.New("cubizy child category not found")
				}
			}
			SubcategoryID = childcategory.SubcategoryID
		} else {
			err = errors.New("Childcategory not found")
		}
	} else if SubcategoryID > 0 {
		var subcategory Subcategory
		err = db.Conn.First(&subcategory, SubcategoryID).Error
		if err == nil {
			if subcategory.BaseSubcategoryID > 0 {
				SubcategoryID = subcategory.BaseSubcategoryID
			}
		} else {
			util.Log("Subcategory not found")
		}
	} else {
		err = errors.New("empty request")
	}
	if err == nil {
		query := db.Conn.Model(&Attribute{})
		if SubcategoryID > 0 && ChildcategoryID > 0 {
			query.Where("(subcategory_id = ? && childcategory_id = 0) OR childcategory_id = ?", SubcategoryID, ChildcategoryID)
		} else if SubcategoryID > 0 {
			query.Where("subcategory_id = ? && childcategory_id = 0", SubcategoryID)
		} else if ChildcategoryID > 0 {
			query.Where("childcategory_id = ?", ChildcategoryID)
		}
		if SubcategoryID > 0 || ChildcategoryID > 0 {
			err = query.Order(" subcategory_id desc, id desc").Find(&attributes).Error
		}
	}
	if err != nil {
		util.Log(err)
	}
	return attributes, err
}

// set column for this attribute in database
func (attribute *Attribute) SetDBColumn() {
	var columnName = "attr"
	switch attribute.FieldType {
	case "Text":
		columnName = "attr_text"
	case "Number":
		columnName = "attr_numb"
	case "Date":
		columnName = "attr_date"
	case "Dropdown":
		columnName = "attr_text"
	}
	var index = 1

	if strings.Contains(attribute.ProductColumn, columnName) {
		util.Log(attribute.ProductColumn, "is allready set for attribute", attribute.Name)
		return
	} else if len(attribute.ProductColumn) > 0 {
		// clear that column
		util.Log(attribute.ProductColumn, "become absolute for attribute", attribute.Name)
	}

	if attribute.ChildcategoryID > 0 {
		columnName = "child_" + columnName
		var attributesWithSameColumnCount int64 = 1
		for attributesWithSameColumnCount > 0 {
			db.Conn.Model(&Attribute{}).Where(" product_column LIKE ? AND childcategory_id = ? ", columnName, attribute.ChildcategoryID).Count(&attributesWithSameColumnCount)
			if attributesWithSameColumnCount == 0 {
				attribute.ProductColumn = columnName
			} else {
				columnName = columnName + strconv.Itoa(index)
				index += index
			}
		}
		attribute.CheckDBColumn()
	} else if attribute.SubcategoryID > 0 {
		columnName = "sub_" + columnName
		var attributesWithSameColumnCount int64 = 1
		for attributesWithSameColumnCount > 0 {
			db.Conn.Model(&Attribute{}).Where(" product_column LIKE ? AND subcategory_id = ? ", columnName, attribute.SubcategoryID).Count(&attributesWithSameColumnCount)
			if attributesWithSameColumnCount == 0 {
				attribute.ProductColumn = columnName
			} else {
				columnName = columnName + strconv.Itoa(index)
				index += index
			}
		}
		attribute.CheckDBColumn()
	}
}

// check column for this attribute in database
func (attribute *Attribute) CheckDBColumn() {
	if len(attribute.ProductColumn) > 0 {
		var query = "SELECT `" + attribute.ProductColumn + "` FROM `productattributes` LIMIT 1"
		err := db.Conn.Exec(query).Error
		if err != nil {
			util.Log("Adding attribute column ")
			switch attribute.FieldType {
			case "Text":
				query = " ALTER TABLE `productattributes` ADD `" + attribute.ProductColumn + "` VARCHAR(150) NOT NULL DEFAULT ''; "
			case "Number":
				query = " ALTER TABLE `productattributes` ADD `" + attribute.ProductColumn + "` INT NOT NULL DEFAULT 0; "
			case "Date":
				query = " ALTER TABLE `productattributes` ADD `" + attribute.ProductColumn + "` INT UNSIGNED NOT NULL DEFAULT 0; "
			case "Dropdown":
				query = " ALTER TABLE `productattributes` ADD `" + attribute.ProductColumn + "` VARCHAR(150) NOT NULL DEFAULT ''; "
			default:
				util.Log("Please set ALTER TABLE query for  :", attribute.FieldType)
			}

			err := db.Conn.Exec(query).Error
			if err != nil {
				util.Log(err.Error())
			}
		}
	}
}
