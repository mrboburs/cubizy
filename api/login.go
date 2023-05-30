package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

var loginapi = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	username := "Invalid"
	if _, ok := response.Request["username"]; ok {
		username = util.GetString(response.Request["username"])
	}
	password := "Invalid"
	if _, ok := response.Request["password"]; ok {
		password = util.GetString(response.Request["password"])
	}
	remember := false
	if _, ok := response.Request["remember"]; ok {
		remember = util.GetBool(response.Request["remember"])
	}

	var user model.User
	err = db.Conn.First(&user, "email = ? OR mobile = ?", username, username).Error
	if err == nil {
		if user.Email == util.Settings.SuperAdmin {
			user.IsAdmin = true
			user.IsSuperAdmin = true
		} else {
			user.IsSuperAdmin = false
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == nil {
			if response.Domain == "admin" {
				if !(user.IsAdmin || user.IsSuperAdmin) {
					response.Message = "User Not Found"
					return err
				}
			}
			user.Remember = remember
		} else if user.EmailCode != "" && password == user.EmailCode && username == user.Email {
			user.EmailVerified = true
			user.EmailCode = ""
			user.EmailCodeSet = false
		} else if user.MobileCode != "" && password == user.MobileCode && username == user.Mobile {
			user.MobileVerified = true
			user.MobileCode = ""
			user.MobileCodeSet = false
		} else {
			util.Log(err)
			err = errors.New("invalid password")
			response.Message = "Password didn't match, try again or try forgot password"
			return err
		}

		response.Message = "Welcome " + user.Name
		user.Token = uuid.New().String()
		user.LastToken = user.Token
		user.LastLoginOn = user.LoginOn
		user.LoginOn = time.Now().Unix()
		user.LastActiveOn = time.Now().Unix()
		err = db.Conn.Save(&user).Error
		if err == nil {
			response.User = &user
			response.Account = model.GetAccount(&user, response.Domain)
			response.Status = apiresponse.SUCCESS
		}
	} else {
		response.Message = "User Not Found"
	}
	return err
}
