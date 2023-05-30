package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Blogcategory will all informetion of category for blogs
type Blogcategory struct {
	gorm.Model
	Name      string
	Status    bool
	Blogs     int64
	CreatedBy uint
	UpdatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Blogcategory{})
	err := db.Conn.Exec("UPDATE blogcategories SET blogcategories.blogs = (SELECT count(blogs.id) FROM blogs WHERE blogs.category = blogcategories.id )").Error
	if err != nil {
		util.Log("Unable to update blog count ", err.Error())
	}
}

// Update will update product by given post argumnets
func (blogcategory *Blogcategory) Update(BlogcategoryMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := BlogcategoryMap["Name"]; ok {
		Value := util.GetString(BlogcategoryMap["Name"])
		if Value != blogcategory.Name {
			blogcategory.Name = Value
			flag = true
		}
	}

	if _, ok := BlogcategoryMap["Status"]; ok {
		Value := util.GetBool(BlogcategoryMap["Status"])
		if Value != blogcategory.Status {
			blogcategory.Status = Value
			flag = true
		}
	}

	if flag {
		if blogcategory.Name == "" {
			err = errors.New(" Name can not be empty")
		} else {
			blogcategory.UpdatedBy = UpdatedBy.ID
			if blogcategory.ID == 0 {
				blogcategory.CreatedBy = UpdatedBy.ID
				var duplicates Blogcategory
				err = db.Conn.First(&duplicates, " `name` = ? ", blogcategory.Name).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&blogcategory).Error
				}
			} else {
				err = db.Conn.Save(&blogcategory).Error
			}
		}
	}
	return flag, err
}
