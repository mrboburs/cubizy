package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

// Address will all informetion of Address
type Address struct {
	gorm.Model

	Title        string `gorm:"size:128,unique,uniqueIndex:unAddress"`
	Mobile       string
	AddressLine1 string
	AddressLine2 string
	AddressLine3 string
	Longitude    string
	Latitude     string
	Code         string
	SubLocality  string
	Locality     string
	District     string
	State        string
	Country      string
	LocationID   uint
	AccountID    uint
	CreatedBy    uint
	UpdatedBy    uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Address{})
}

// Update will update product by given post argumnets
func (address *Address) Update(AddressMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := AddressMap["Title"]; ok {
		Value := util.GetString(AddressMap["Title"])
		if Value != address.Title {
			address.Title = Value
			flag = true
		}
	}

	if _, ok := AddressMap["Mobile"]; ok {
		Value := util.GetString(AddressMap["Mobile"])
		if Value != address.Mobile {
			address.Mobile = Value
			flag = true
		}
	}

	if _, ok := AddressMap["AddressLine1"]; ok {
		Value := util.GetString(AddressMap["AddressLine1"])
		if Value != address.AddressLine1 {
			address.AddressLine1 = Value
			flag = true
		}
	}

	if _, ok := AddressMap["AddressLine2"]; ok {
		Value := util.GetString(AddressMap["AddressLine2"])
		if Value != address.AddressLine2 {
			address.AddressLine2 = Value
			flag = true
		}
	}

	if _, ok := AddressMap["AddressLine3"]; ok {
		Value := util.GetString(AddressMap["AddressLine3"])
		if Value != address.AddressLine3 {
			address.AddressLine3 = Value
			flag = true
		}
	}

	if _, ok := AddressMap["AddressLine3"]; ok {
		Value := util.GetString(AddressMap["AddressLine3"])
		if Value != address.AddressLine3 {
			address.AddressLine3 = Value
			flag = true
		}
	}

	if _, ok := AddressMap["Longitude"]; ok {
		Value := util.GetString(AddressMap["Longitude"])
		if Value != address.Longitude {
			address.Longitude = Value
			flag = true
		}
	}

	if _, ok := AddressMap["Latitude"]; ok {
		Value := util.GetString(AddressMap["Latitude"])
		if Value != address.Latitude {
			address.Latitude = Value
			flag = true
		}
	}

	if _, ok := AddressMap["State"]; ok {
		Value := util.GetString(AddressMap["State"])
		if Value != address.State {
			address.State = Value
			flag = true
		}
	}

	location := Location{}

	if _, ok := AddressMap["Country"]; ok {
		Value := util.GetString(AddressMap["Country"])
		if Value != address.Country {
			address.Country = Value
			location.Country = Value
			flag = true
		}
	}

	if _, ok := AddressMap["District"]; ok {
		Value := util.GetString(AddressMap["District"])
		if Value != address.District {
			address.District = Value
			location.District = Value
			flag = true
		}
	}

	if _, ok := AddressMap["Locality"]; ok {
		Value := util.GetString(AddressMap["Locality"])
		if Value != address.Locality {
			address.Locality = Value
			location.Locality = Value
			flag = true
		}
	}

	if _, ok := AddressMap["SubLocality"]; ok {
		Value := util.GetString(AddressMap["SubLocality"])
		if Value != address.SubLocality {
			address.SubLocality = Value
			location.SubLocality = Value
			flag = true
		}
	}

	if _, ok := AddressMap["Code"]; ok {
		Value := util.GetString(AddressMap["Code"])
		if Value != address.Code {
			address.Code = Value
			location.Code = Value
			flag = true
		}
	}

	var old_location Location
	_ = db.Conn.Where(old_location).First(&location).Error
	if old_location.ID > 0 {
		if old_location.ID != address.LocationID {
			address.LocationID = old_location.ID
			flag = true
		}
	} else {
		_ = db.Conn.Create(&location).Error
	}

	if flag {
		if address.Title == "" {
			err = errors.New(" Title can not be empty")
		} else if address.AddressLine1 == "" {
			err = errors.New(" AddressLine1 can not be empty")
		} else {
			address.UpdatedBy = UpdatedBy.ID
			if address.ID == 0 {
				address.CreatedBy = UpdatedBy.ID
				var duplicates Address
				err = db.Conn.First(&duplicates, " `title` = ? AND `account_id` = ? ", address.Title, address.AccountID).Error
				if err == nil {
					err = errors.New("duplicate code")
				} else {
					err = db.Conn.Create(&address).Error
				}
			} else {
				err = db.Conn.Save(&address).Error
			}
		}
	}
	return flag, err
}

func (address *Address) ToJson() string {
	json_string, err := json.Marshal(address)
	if err != nil {
		util.Log("account ToJson : ", err)
	}
	return string(json_string)
}
