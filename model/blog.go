package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Blog will all informetion of Blog
type Blog struct {
	gorm.Model

	Title     string `gorm:"size:128,unique,uniqueIndex:unBlog"`
	Image     string
	Content   string `gorm:"type:text"`
	Status    bool
	Tags      string
	Category  uint
	AccountID uint
	CreatedBy uint
	UpdatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Blog{})
}

// Update will update product by given post argumnets
func (blog *Blog) Update(BlogMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := BlogMap["Title"]; ok {
		Value := util.GetString(BlogMap["Title"])
		if Value != blog.Title {
			blog.Title = Value
			flag = true
		}
	}

	if _, ok := BlogMap["Image"]; ok {
		Value := util.GetString(BlogMap["Image"])
		if Value != blog.Image {
			blog.Image = Value
			flag = true
		}
	}

	if _, ok := BlogMap["Tags"]; ok {
		Value := util.GetString(BlogMap["Tags"])
		if Value != blog.Tags {
			blog.Tags = Value
			flag = true
		}
	}

	var oldCategory uint
	if _, ok := BlogMap["Category"]; ok {
		Value := util.GetUint(BlogMap["Category"])
		if Value != blog.Category {
			oldCategory = blog.Category
			blog.Category = Value
			flag = true
		}
	}

	if _, ok := BlogMap["Content"]; ok {
		Value := util.GetString(BlogMap["Content"])
		if Value != blog.Content {
			blog.Content = Value
			flag = true
		}
	}

	if _, ok := BlogMap["Status"]; ok {
		Value := util.GetBool(BlogMap["Status"])
		if Value != blog.Status {
			blog.Status = Value
			flag = true
		}
	}

	if flag {
		if blog.Title == "" {
			err = errors.New(" Title can not be empty")
		} else if blog.Content == "" {
			err = errors.New(" Content can not be empty")
		} else {
			blog.UpdatedBy = UpdatedBy.ID
			if blog.ID == 0 {
				blog.CreatedBy = UpdatedBy.ID
				var duplicates Blog
				err = db.Conn.First(&duplicates, " `title` = ? ", blog.Title).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&blog).Error
				}
			} else {
				err = db.Conn.Save(&blog).Error
			}
		}

		if err == nil {
			if oldCategory != blog.Category {
				if oldCategory > 0 {
					updateBLogsCount(oldCategory)
				}
				updateBLogsCount(blog.Category)
			}
		}
	}
	return flag, err
}

func updateBLogsCount(category_id uint) {
	err := db.Conn.Exec("UPDATE blogcategories SET blogcategories.blogs = (SELECT count(blogs.id) FROM blogs WHERE blogs.category = blogcategories.id ) WHERE blogcategories.id = ?", category_id).Error
	if err != nil {
		util.Log("Unable to update blog count ", err.Error())
	}
}
