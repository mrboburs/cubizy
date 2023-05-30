package model

import (
	"cubizy/keys"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string
	Photo             string
	Email             string `gorm:"unique"`
	EmailCode         string `json:"-"`
	EmailCodeSet      bool
	EmailVerified     bool
	Mobile            string `gorm:"unique"`
	MobileCode        string `json:"-"`
	MobileCodeSet     bool
	MobileVerified    bool
	SupportagentQuote string
	SupportCount      int
	Password          string `json:"-"`
	PasswordSet       bool
	Token             string
	LastToken         string `json:"-"`
	LoginOn           int64
	LastLoginOn       int64
	LastActiveOn      int64
	Remember          bool
	Wallet            uint
	Wishlist          int64
	Orders            int64
	Status            bool
	CreatedBy         uint
	UpdatedBy         uint
	Online            bool
	IsAdmin           bool
	IsSuperAdmin      bool
	IsSupportagent    bool
	IsStudent         bool
	Joined            bool

	SellerAccountID  uint
	AddedByAccountID uint

	Question1 string
	Answer1   string

	Question2 string
	Answer2   string

	Question3 string
	Answer3   string

	DefaultAddress uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&User{})
	db.Conn.Exec("UPDATE `users` SET `online` = 0;")
	db.Conn.Exec("UPDATE `users` SET `orders` = (SELECT count(id) FROM orders WHERE created_by = users.id);")
	checkSuperAdmin()
}

func checkSuperAdmin() {
	var superAdmin User
	err := db.Conn.First(&superAdmin, User{Email: util.Settings.SuperAdmin}).Error
	if err == nil {
		if !superAdmin.IsSuperAdmin {
			db.Conn.Exec("UPDATE `users` SET `is_super_admin` = '0';")

			superAdmin.IsAdmin = true
			superAdmin.IsSuperAdmin = true

			err = db.Conn.Save(&superAdmin).Error
		}
	} else {
		if strings.Contains(err.Error(), "record not found") {
			err = nil
			var passwordHash []byte
			passwordHash, err = bcrypt.GenerateFromPassword([]byte(util.Settings.Password), bcrypt.DefaultCost)
			if err == nil {
				superAdmin := User{
					Name:         "Super Admin",
					Email:        util.Settings.SuperAdmin,
					IsSuperAdmin: true,
					Password:     string(passwordHash),
				}
				err = db.Conn.Create(&superAdmin).Error
				if err == nil {
					util.Log("Super admin added")
				}
			}
		}
	}
	if err != nil {
		util.Panic(err)
	}
}

// GetUserByToken will return online user by his code
func GetUserByToken(token string) (*User, error) {
	var user *User
	var err error
	user = nil
	if token != "" {
		var testUser User
		err = db.Conn.First(&testUser, " token = ? OR last_token = ? ", token, token).Error
		if err == nil {
			timeOutString := GetSetting(keys.TimeOut, "30")
			timeout, _ := strconv.ParseInt(timeOutString, 10, 64)
			if testUser.Remember {
				timeOutString = GetSetting(keys.TimeOutOnRemember, "30")
				timeout, _ = strconv.ParseInt(timeOutString, 10, 64)
			}
			timeout = timeout * 60
			timenow := time.Now().Unix()
			minimumActiveTime := (timenow - timeout)
			if testUser.LastActiveOn < minimumActiveTime {
				util.Log(" user login timeout by ", timeOutString)
				err = errors.New("logout by timeout, please login again")
				testUser.LastToken = ""
				testUser.Token = ""
				db.Conn.Save(testUser)
			} else {
				if testUser.LastActiveOn < (timenow - (timeout / 3)) {
					testUser.LastToken = testUser.Token
					testUser.Token = uuid.New().String()
				}
				testUser.LastActiveOn = time.Now().Unix()
				util.Log("User LastActiveOn updated")
				db.Conn.Save(testUser)
				user = &testUser
			}
		}
	}
	return user, err
}

// Update will update product by given post argumnets
func (user *User) Update(userMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool

	if UpdatedBy.ID == user.ID || (!user.Joined && UpdatedBy.ID == user.CreatedBy) {
		if _, ok := userMap["Photo"]; ok {
			Value := util.GetString(userMap["Photo"])
			if Value != user.Photo {
				user.Photo = Value
				flag = true
			}
		}

		if _, ok := userMap["Name"]; ok {
			Value := util.GetString(userMap["Name"])
			if Value != user.Name {
				user.Name = Value
				flag = true
			}
		}

		if _, ok := userMap["SupportagentQuote"]; ok {
			Value := util.GetString(userMap["SupportagentQuote"])
			if Value != user.SupportagentQuote {
				user.SupportagentQuote = Value
				flag = true
			}
		}

		if _, ok := userMap["Password"]; ok && user.ID == 0 {
			Value := util.GetString(userMap["Password"])
			var passwordHash []byte
			passwordHash, err = bcrypt.GenerateFromPassword([]byte(Value), bcrypt.DefaultCost)
			if err == nil {
				user.Password = string(passwordHash)
				flag = true
			}
		}

		if _, ok := userMap["Email"]; ok {
			Value := util.GetString(userMap["Email"])
			if Value != user.Email {
				user.Email = Value
				user.EmailVerified = false
				flag = true
			}
		}

		if _, ok := userMap["Mobile"]; ok {
			Value := util.GetString(userMap["Mobile"])
			if Value != user.Mobile {
				user.Mobile = Value
				user.MobileVerified = false
				flag = true
			}
		}
	}
	if UpdatedBy.ID == user.ID {
		if _, ok := userMap["Question1"]; ok {
			Value := util.GetString(userMap["Question1"])
			if Value != user.Question1 {
				user.Question1 = Value
				flag = true
			}
		}
		if _, ok := userMap["Answer1"]; ok {
			Value := util.GetString(userMap["Answer1"])
			if Value != user.Answer1 {
				user.Answer1 = Value
				flag = true
			}
		}

		if _, ok := userMap["Question2"]; ok {
			Value := util.GetString(userMap["Question2"])
			if Value != user.Question2 {
				user.Question2 = Value
				flag = true
			}
		}
		if _, ok := userMap["Answer2"]; ok {
			Value := util.GetString(userMap["Answer2"])
			if Value != user.Answer2 {
				user.Answer2 = Value
				flag = true
			}
		}

		if _, ok := userMap["Question3"]; ok {
			Value := util.GetString(userMap["Question3"])
			if Value != user.Question3 {
				user.Question3 = Value
				flag = true
			}
		}
		if _, ok := userMap["Answer3"]; ok {
			Value := util.GetString(userMap["Answer3"])
			if Value != user.Answer3 {
				user.Answer3 = Value
				flag = true
			}
		}

		if _, ok := userMap["DefaultAddress"]; ok {
			Value := util.GetUint(userMap["DefaultAddress"])
			if Value != user.DefaultAddress {
				user.DefaultAddress = Value
				flag = true
			}
		}
	}

	if _, ok := userMap["Joined"]; ok {
		Value := util.GetBool(userMap["Joined"])
		if (Value != user.Joined) && Value {
			user.Joined = Value
			flag = true
		}
	}
	if UpdatedBy.IsAdmin || UpdatedBy.IsSuperAdmin {
		if _, ok := userMap["IsSupportagent"]; ok {
			Value := util.GetBool(userMap["IsSupportagent"])
			if Value != user.IsSupportagent {
				user.IsSupportagent = Value
				flag = true
			}
		}
		if _, ok := userMap["Status"]; ok {
			Value := util.GetBool(userMap["Status"])
			if Value != user.Status {
				user.Status = Value
				flag = true
			}
		}
		if _, ok := userMap["SellerAccountID"]; ok {
			Value := util.GetUint(userMap["SellerAccountID"])
			if Value != user.SellerAccountID {
				user.SellerAccountID = Value
				flag = true
			}
		}
	}
	if UpdatedBy.IsSuperAdmin {
		if _, ok := userMap["IsAdmin"]; ok {
			Value := util.GetBool(userMap["IsAdmin"])
			if Value != user.IsAdmin {
				user.IsAdmin = Value
				flag = true
			}
		}
		if _, ok := userMap["IsSuperAdmin"]; ok {
			Value := util.GetBool(userMap["IsSuperAdmin"])
			if Value != user.IsSuperAdmin {
				user.IsSuperAdmin = Value
				flag = true
			}
		}
		if _, ok := userMap["Wallet"]; ok {
			Value := util.GetUint(userMap["Wallet"])
			if Value != user.Wallet {
				user.Wallet = Value
				flag = true
			}
		}
	}
	if flag {
		if user.Email == util.Settings.SuperAdmin {
			user.IsSuperAdmin = true
			user.IsAdmin = true
		}
		if user.Name == "" {
			err = errors.New(" Name can not be empty")
		} else if user.Email == "" {
			err = errors.New(" Email can not be empty")
		} else if user.Mobile == "" {
			err = errors.New(" Mobile can not be empty")
		} else {
			user.UpdatedBy = UpdatedBy.ID
			if user.ID == 0 {
				user.CreatedBy = UpdatedBy.ID
				var duplicates User
				err = db.Conn.First(&duplicates, " email = ? OR mobile = ? ", user.Email, user.Mobile).Error
				if err == nil {
					err = errors.New("duplicate user")
				} else {
					err = db.Conn.Create(&user).Error
				}
			} else {
				err = db.Conn.Save(&user).Error
			}
		}
	}
	return flag, err
}

// UpdateWishlistCount will recount wishlist
func (user *User) UpdateWishlistCount() {
	var count int64

	err := db.Conn.Model(&Wishlist{}).Where("created_by = ?", user.ID).Count(&count).Error
	if err != nil {
		util.Log("While counting wishlists items ")
		util.Log(err)
		err = nil
	} else {
		user.Wishlist = count
		err = db.Conn.Save(&user).Error
		if err != nil {
			util.Log("While saving wishlist count ")
			util.Log(err)
			err = nil
		}
	}

}

// UpdateOrderCount will recount order
func (user *User) UpdateOrderCount() {
	var count int64

	err := db.Conn.Model(&Order{}).Where("created_by = ?", user.ID).Count(&count).Error
	if err != nil {
		util.Log("While counting orders items ")
		util.Log(err)
		err = nil
	} else {
		user.Orders = count
		err = db.Conn.Save(&user).Error
		if err != nil {
			util.Log("While saving order count ")
			util.Log(err)
			err = nil
		}
	}

}
