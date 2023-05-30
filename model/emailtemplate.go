package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Emailtemplate will all informetion of Emailtemplate
type Emailtemplate struct {
	gorm.Model
	LanguageCode string
	Subject      string
	Template     string `gorm:"type:text"`
	CreatedBy    uint
	UpdatedBy    uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Emailtemplate{})
}

// Update will update product by given post argumnets
func (emailtemplate *Emailtemplate) Update(emailtemplateMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := emailtemplateMap["LanguageCode"]; ok {
		Value := util.GetString(emailtemplateMap["LanguageCode"])
		if Value != emailtemplate.LanguageCode {
			emailtemplate.LanguageCode = Value
			flag = true
		}
	}

	if _, ok := emailtemplateMap["Subject"]; ok {
		Value := util.GetString(emailtemplateMap["Subject"])
		if Value != emailtemplate.Subject {
			emailtemplate.Subject = Value
			flag = true
		}
	}

	if _, ok := emailtemplateMap["Template"]; ok {
		Value := util.GetString(emailtemplateMap["Template"])
		if Value != emailtemplate.Template {
			emailtemplate.Template = Value
			flag = true
		}
	}

	if flag {
		if emailtemplate.LanguageCode == "" {
			err = errors.New(" LanguageCode can not be empty")
		} else if emailtemplate.Subject == "" {
			err = errors.New(" Subject can not be empty")
		} else if emailtemplate.Template == "" {
			err = errors.New(" Template can not be empty")
		} else {
			emailtemplate.UpdatedBy = UpdatedBy.ID
			if emailtemplate.ID == 0 {
				emailtemplate.CreatedBy = UpdatedBy.ID
				var duplicates Emailtemplate
				err = db.Conn.First(&duplicates, " subject = ? AND language_code = ? ", emailtemplate.Subject, emailtemplate.LanguageCode).Error
				if err == nil {
					err = errors.New("duplicate key")
				} else {
					err = db.Conn.Create(&emailtemplate).Error
				}
			} else {
				err = db.Conn.Save(&emailtemplate).Error
			}
		}
	}

	return flag, err
}

func GetEmailTemplate(id uint) (Emailtemplate, error) {
	var emailtemplate Emailtemplate
	err := db.Conn.First(&emailtemplate, id).Error
	return emailtemplate, err
}
