package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// Location will all informetion of Location
type Location struct {
	gorm.Model

	Country     string `gorm:"size:128,unique,uniqueIndex:unlocation"`
	Code        string `gorm:"size:128,unique,uniqueIndex:unlocation"`
	Locality    string
	SubLocality string
	District    string
	Longitude   uint
	Latitude    uint
	CreatedBy   uint
	UpdatedBy   uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Location{})
}

// Update will update product by given post argumnets
func (location *Location) Update(locationMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}
	if _, ok := locationMap["Code"]; ok {
		location.Code = strings.TrimSpace(locationMap["Code"].(string))
		flag = true
	}

	if _, ok := locationMap["SubLocality"]; ok {
		location.SubLocality = strings.TrimSpace(locationMap["SubLocality"].(string))
		flag = true
	}

	if _, ok := locationMap["Locality"]; ok {
		location.Locality = strings.TrimSpace(locationMap["Locality"].(string))
		flag = true
	}

	if _, ok := locationMap["District"]; ok {
		location.District = strings.TrimSpace(locationMap["District"].(string))
		flag = true
	}

	if _, ok := locationMap["Longitude"]; ok {
		location.Longitude = util.GetUint(locationMap["Longitude"])
		flag = true
	}

	if _, ok := locationMap["Latitude"]; ok {
		location.Latitude = util.GetUint(locationMap["Latitude"])
		flag = true
	}

	if _, ok := locationMap["Country"]; ok {
		location.Country = util.GetString(locationMap["Country"])
		flag = true
	}

	if flag {
		if location.Code == "" {
			err = errors.New(" Code can not be empty")
		} else if location.Country == "" {
			err = errors.New(" Country can not be empty")
		} else if location.Locality == "" {
			err = errors.New(" Locality can not be empty")
		} else if location.SubLocality == "" {
			err = errors.New(" SubLocality can not be empty")
		} else {
			location.UpdatedBy = UpdatedBy.ID
			if location.ID == 0 {
				location.CreatedBy = UpdatedBy.ID
				var duplicates Location
				err = db.Conn.First(&duplicates, " code = ? AND country = ? ", location.Code, location.Country).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&location).Error
				}
			} else {
				err = db.Conn.Save(&location).Error
			}
		}
	}
	return flag, err
}
