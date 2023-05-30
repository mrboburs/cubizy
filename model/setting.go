package model

import (
	"cubizy/keys"
	"cubizy/plugins/db"
	"cubizy/util"

	"gorm.io/gorm"
)

// Setting hold setting values also store and get it in database
type Setting struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);UNIQUE"`
	Value     string `gorm:"type:TEXT;"`
	Details   string `gorm:"type:TEXT;"`
	Type      string `gorm:"type:TEXT;"`
	UpdatedBy uint
}

var allsettings map[string]Setting

func init() {
	// Migrate the schema for Setting obj
	db.Conn.AutoMigrate(&Setting{})
	allsettings = make(map[string]Setting)
	var settings []Setting
	err := db.Conn.Find(&settings).Error
	if err == nil {
		for _, setting := range settings {
			allsettings[setting.Name] = setting
		}

		GetSetting(keys.TimeOut, "30")

		GetSetting(keys.TimeOutOnRemember, "1440")

		GetSetting(keys.SendgridSenderName, "")

		GetSetting(keys.SendgridSenderEmail, "")

		GetSetting(keys.SendgridAPIKey, "")

		GetSetting(keys.S3Bucket, "")

		GetSetting(keys.S3Region, "")

		GetSetting(keys.S3AwsAccessKeyID, "")

		GetSetting(keys.S3AwsSecretAccessKey, "")

	} else {
		util.Log("settings not initiated")
		util.Log(err)
	}
}

// GetSetting will give value of setting
func GetSetting(key string, defaultValue string) string {
	setting, ok := allsettings[key]
	if !ok {
		setting = Setting{
			Name:    key,
			Value:   defaultValue,
			Details: getDetails(key),
		}
		err := db.Conn.Create(&setting).Error
		if err != nil {
			util.Log(err)
		}
		allsettings[key] = setting
	}
	return setting.Value
}

// UpdateSettings will update setting value in database and in application
func UpdateSettings(settingid int, settingValue string) error {
	err := db.Conn.Table("settings").Where("id = ?", settingid).Updates(map[string]interface{}{"value": settingValue}).Error
	if err == nil {
		var setting Setting
		err = db.Conn.Where("id = ?", settingid).First(&setting).Error
		allsettings[setting.Name] = setting
	}
	return err
}

func getDetails(settingKey string) string {
	switch settingKey {
	case keys.TimeOut:
		return "Auto logout after seconds"
	case keys.TimeOutOnRemember:
		return "Auto logout after seconds is user selected remember me"
	case keys.SendgridSenderName:
		return "Default Sender name while using Sendgrid API Key to send email from applicaion"
	case keys.SendgridSenderEmail:
		return "Default Sender email while using Sendgrid API Key to send email from applicaion"
	case keys.SendgridAPIKey:
		return "Sendgrid API Key to send email notifications"
	case keys.S3Bucket:
		return "S3Bucket where all files will get stored"
	case keys.S3Region:
		return "S3Region where all files will get stored"
	case keys.S3AwsAccessKeyID:
		return "S3AwsAccessKeyID to authenticate request on S3Bucket "
	case keys.S3AwsSecretAccessKey:
		return "S3AwsSecretAccessKey to authenticate request on S3Bucket "
	}
	return ""
}

// GetAllSettings will provied array of all settings in database
func GetAllSettings() ([]Setting, error) {
	var settings []Setting
	err := db.Conn.Find(&settings).Error
	return settings, err
}
