package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountType        string
	Title              string `gorm:"unique"`
	Description        string
	Keywords           string
	Email              string `gorm:"unique"`
	Mobile             string
	AddressID          uint
	Logo               string
	WideLogo           string
	Banner             string
	Youtube            string
	Facebook           string
	Instagram          string
	Pinterest          string
	WhatsApp           string
	IDProof            string
	IDStatus           bool
	AddressProof       string
	AddressStatus      bool
	RegistretionProof  string
	RegistretionStatus bool
	OtherDocument      string
	TestAccount        bool
	Status             int
	StatusComment      string

	ThemeID       uint
	Theme         string
	ThemePath     string
	ThemeSettings string `json:"-"`
	Subdomain     string
	Domain        string
	CanActive     bool `gorm:"default:false"`
	Active        bool `gorm:"default:false"`
	CreatedBy     uint
	UpdatedBy     uint
	LastActiveOn  int64

	Wallet uint

	MaxPrice uint

	Rating    int
	Reviews   int
	OneStar   int
	TwoStar   int
	ThreeStar int
	FourStar  int
	FiveStar  int
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Account{})
	checkSuperAdminAccount()
	db.Conn.Exec("Update users set seller_account_id = (SELECT id FROM accounts WHERE accounts.account_type = 'seller' AND accounts.created_by = users.id)")
	db.Conn.Exec("Update accounts SET max_price = (SELECT MAX(max_price) FROM products WHERE products.account_id = accounts.id)")
	db.Conn.Exec("Update accounts SET max_price = (SELECT MAX(max_price) FROM products) WHERE accounts.account_type = 'admin'")
}

// SuperAdminAccount hold application super admin
var SuperAdminAccount Account

func checkSuperAdminAccount() {

	err := db.Conn.First(&SuperAdminAccount, Account{AccountType: "admin"}).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			err = nil
			SuperAdminAccount = Account{
				AccountType: "admin",
				Title:       util.Settings.AppName,
				Domain:      util.Settings.Domain,
				Subdomain:   "priview",
				Email:       util.Settings.SuperAdmin,
				CreatedBy:   1,
				UpdatedBy:   1,
				CanActive:   true,
				Active:      true,
				Status:      10,
			}
			err = db.Conn.Create(&SuperAdminAccount).Error
			if err == nil {
				util.Log("Super admin account added")
			}
		}
	} else {
		if SuperAdminAccount.Domain != util.Settings.Domain {
			SuperAdminAccount.Domain = util.Settings.Domain
		}
		if SuperAdminAccount.Title != util.Settings.AppName {
			SuperAdminAccount.Title = util.Settings.AppName
		}
		if SuperAdminAccount.Email != util.Settings.SuperAdmin {
			SuperAdminAccount.Email = util.Settings.SuperAdmin
		}
		if SuperAdminAccount.Status != 10 {
			SuperAdminAccount.Status = 10
		}
		if !SuperAdminAccount.CanActive {
			SuperAdminAccount.CanActive = true
		}
		db.Conn.Save(SuperAdminAccount)
	}
	if err != nil {
		util.Panic(err)
	}
}

func (account *Account) ToJson() string {
	json_string, err := json.Marshal(account)
	if err != nil {
		util.Log("account ToJson : ", err)
	}
	return string(json_string)
}

func (account *Account) Update(accountMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool

	if account.ID > 0 && !(UpdatedBy.IsAdmin || UpdatedBy.ID == account.CreatedBy) {
		return flag, errors.New("invalid updater" + strconv.Itoa(int(account.ID)))
	}
	if UpdatedBy.IsAdmin {
		if _, ok := accountMap["Status"]; ok {
			Value := util.GetInt(accountMap["Status"])
			if Value != account.Status {
				account.Status = Value
				flag = true
			}
		}
		if _, ok := accountMap["StatusComment"]; ok {
			Value := util.GetString(accountMap["StatusComment"])
			if Value != account.StatusComment {
				account.StatusComment = Value
				flag = true
			}
		}
		if _, ok := accountMap["TestAccount"]; ok {
			Value := util.GetBool(accountMap["TestAccount"])
			if Value != account.TestAccount {
				account.TestAccount = Value
				flag = true
			}
		}
		if _, ok := accountMap["CanActive"]; ok {
			Value := util.GetBool(accountMap["CanActive"])
			if Value != account.CanActive {
				account.CanActive = Value
				flag = true
			}
		}
		if _, ok := accountMap["IDStatus"]; ok {
			Value := util.GetBool(accountMap["IDStatus"])
			if Value != account.IDStatus {
				account.IDStatus = Value
				flag = true
			}
		}
		if _, ok := accountMap["AddressStatus"]; ok {
			Value := util.GetBool(accountMap["AddressStatus"])
			if Value != account.AddressStatus {
				account.AddressStatus = Value
				flag = true
			}
		}
		if _, ok := accountMap["RegistretionStatus"]; ok {
			Value := util.GetBool(accountMap["RegistretionStatus"])
			if Value != account.RegistretionStatus {
				account.RegistretionStatus = Value
				flag = true
			}
		}
	}

	if _, ok := accountMap["Logo"]; ok {
		Value := util.GetString(accountMap["Logo"])
		if Value != account.Logo {
			account.Logo = Value
			flag = true
		}
	}
	if _, ok := accountMap["WideLogo"]; ok {
		Value := util.GetString(accountMap["WideLogo"])
		if Value != account.WideLogo {
			account.WideLogo = Value
			flag = true
		}
	}
	if _, ok := accountMap["Banner"]; ok {
		Value := util.GetString(accountMap["Banner"])
		if Value != account.Banner {
			account.Banner = Value
			flag = true
		}
	}
	if _, ok := accountMap["Title"]; ok {
		Value := util.GetString(accountMap["Title"])
		if Value != account.Title {
			account.Title = Value
			flag = true
		}
	}
	if _, ok := accountMap["Description"]; ok {
		Value := util.GetString(accountMap["Description"])
		if Value != account.Description {
			account.Description = Value
			flag = true
		}
	}
	if _, ok := accountMap["Email"]; ok {
		Value := util.GetString(accountMap["Email"])
		if Value != account.Email {
			account.Email = Value
			flag = true
		}
	}
	if _, ok := accountMap["Mobile"]; ok {
		Value := util.GetString(accountMap["Mobile"])
		if Value != account.Mobile {
			account.Mobile = Value
			flag = true
		}
	}
	if _, ok := accountMap["AddressID"]; ok {
		Value := util.GetUint(accountMap["AddressID"])
		if Value != account.AddressID {
			account.AddressID = Value
			flag = true
		}
	}
	if _, ok := accountMap["ThemeID"]; ok {
		Value := util.GetUint(accountMap["ThemeID"])
		if Value != account.AddressID {
			account.ThemeID = Value
			flag = true
		}
	}
	if _, ok := accountMap["Youtube"]; ok {
		Value := util.GetString(accountMap["Youtube"])
		if Value != account.Youtube {
			account.Youtube = Value
			flag = true
		}
	}
	if _, ok := accountMap["Facebook"]; ok {
		Value := util.GetString(accountMap["Facebook"])
		if Value != account.Facebook {
			account.Facebook = Value
			flag = true
		}
	}
	if _, ok := accountMap["Instagram"]; ok {
		Value := util.GetString(accountMap["Instagram"])
		if Value != account.Instagram {
			account.Instagram = Value
			flag = true
		}
	}
	if _, ok := accountMap["Pinterest"]; ok {
		Value := util.GetString(accountMap["Pinterest"])
		if Value != account.Pinterest {
			account.Pinterest = Value
			flag = true
		}
	}
	if _, ok := accountMap["WhatsApp"]; ok {
		Value := util.GetString(accountMap["WhatsApp"])
		if Value != account.WhatsApp {
			account.WhatsApp = Value
			flag = true
		}
	}

	if _, ok := accountMap["IDProof"]; ok {
		Value := util.GetString(accountMap["IDProof"])
		if Value != account.IDProof {
			account.IDProof = Value
			flag = true
		}
	}

	if _, ok := accountMap["AddressProof"]; ok {
		Value := util.GetString(accountMap["AddressProof"])
		if Value != account.AddressProof {
			account.AddressProof = Value
			if account.AddressProof == "" || account.AddressProof == "0" {
				account.AddressStatus = false
				account.Status = 0
			}
			flag = true
		}
	}

	if _, ok := accountMap["RegistretionProof"]; ok {
		Value := util.GetString(accountMap["RegistretionProof"])
		if Value != account.RegistretionProof {
			account.RegistretionProof = Value
			flag = true
		}
	}

	if _, ok := accountMap["OtherDocument"]; ok {
		Value := util.GetString(accountMap["OtherDocument"])
		if Value != account.OtherDocument {
			account.OtherDocument = Value
			flag = true
		}
	}

	if _, ok := accountMap["Active"]; ok {
		Value := util.GetBool(accountMap["Active"])
		if Value != account.Active {
			account.Active = Value
			flag = true
		}
	}
	if !account.CanActive {
		account.Active = false
	}
	if !account.CanActive && account.Active {
		account.Active = false
		err = errors.New("please contact support make it active")
	}
	if _, ok := accountMap["ThemeID"]; ok {
		Value := util.GetUint(accountMap["ThemeID"])
		if Value != account.ThemeID {
			account.ThemeID = Value
			flag = true
		}
	}

	if _, ok := accountMap["Theme"]; ok {
		Value := util.GetString(accountMap["Theme"])
		if Value != account.Theme {
			account.Theme = Value
			flag = true
		}
	}

	if _, ok := accountMap["ThemePath"]; ok {
		Value := util.GetString(accountMap["ThemePath"])
		if Value != account.ThemePath {
			account.ThemePath = Value
			flag = true
		}
	}

	if _, ok := accountMap["ThemeSettings"]; ok {
		Value := util.GetString(accountMap["ThemeSettings"])
		if Value != account.ThemeSettings {
			account.ThemeSettings = Value
			flag = true
		}
	}
	if _, ok := accountMap["Subdomain"]; ok {
		Value := util.GetString(accountMap["Subdomain"])
		if Value != account.Subdomain {
			account.Subdomain = Value
			flag = true
		}
	}
	if _, ok := accountMap["Domain"]; ok {
		Value := util.GetString(accountMap["Domain"])
		if Value != account.Domain {
			account.Domain = Value
			flag = true
		}
	}
	if _, ok := accountMap["Keywords"]; ok {
		Value := util.GetString(accountMap["Keywords"])
		if Value != account.Keywords {
			account.Keywords = Value
			flag = true
		}
	}

	if flag {
		account.UpdatedBy = UpdatedBy.ID
		if account.ID == 0 {
			account.CreatedBy = UpdatedBy.ID
			var duplicates Account
			err = db.Conn.First(&duplicates, " (created_by = ? AND account_type = ?) OR (account_type = 'admin' AND account_type = ?) OR ( domain != '' AND domain = ?) OR  ( subdomain != '' AND subdomain = ?) ", account.CreatedBy, account.AccountType, account.AccountType, account.Domain, account.Subdomain).Error
			if err == nil {
				err = errors.New("duplicate account")
			} else {
				err = db.Conn.Create(&account).Error
				if account.AccountType == "seller" {
					UpdatedBy.SellerAccountID = account.ID
				}
				db.Conn.Save(UpdatedBy)
			}
		} else {
			var duplicates Account
			err = db.Conn.First(&duplicates, "  id != ? AND (( domain != '' AND domain = ?) OR  ( subdomain != '' AND subdomain = ?)) ", account.ID, account.Domain, account.Subdomain).Error
			if err == nil {
				err = errors.New("duplicate domain or subdomain")
			} else {
				err = db.Conn.Save(&account).Error
			}

		}
	}
	return flag, err
}

// GetAccount by user and current subdomin
func GetAccount(user *User, subdomin string) *Account {
	var err error
	var account Account
	if (user.IsAdmin || user.IsSuperAdmin) && subdomin == "admin" {
		err = db.Conn.Where(" account_type = 'admin'").First(&account).Error
	} else if subdomin == "seller" && user.SellerAccountID > 0 {
		err = db.Conn.First(&account, user.SellerAccountID).Error
	} else {
		err = errors.New("invalid subdomin")
	}
	if err != nil {
		util.Log("GetAccount", subdomin, err.Error())
		return nil
	} else {
		return &account
	}
}

func GetAccountByDomin(subdomin string) *Account {
	var err error
	var account Account
	if subdomin == "" || subdomin == "admin" || subdomin == "seller" {
		err = db.Conn.Where(" account_type = 'admin'").First(&account).Error
		if err == nil {
			db.Conn.Exec("Update accounts SET max_price = (SELECT MAX(max_price) FROM products) WHERE accounts.account_type = 'admin'")
		}
	} else {
		err = db.Conn.Where("active = true AND (subdomain = ? OR domain = ?) ", subdomin, subdomin).First(&account).Error
	}
	if err != nil {
		return nil
	} else {
		if account.MaxPrice > 0 && account.MaxPrice%10 != 0 {
			var t uint
			t = 10
			for t < account.MaxPrice {
				t *= 10
			}
			account.MaxPrice = t
			db.Conn.Save(account)
		}
		return &account
	}
}
