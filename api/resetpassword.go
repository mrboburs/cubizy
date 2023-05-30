package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var resetpasswordAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	username := "Invalid"
	if _, ok := response.Request["username"]; ok {
		username = util.GetString(response.Request["username"])
	}
	password := ""
	if _, ok := response.Request["password"]; ok {
		password = util.GetString(response.Request["password"])
	}
	resetcode := ""
	if _, ok := response.Request["resetcode"]; ok {
		resetcode = util.GetString(response.Request["resetcode"])
	}
	if resetcode == "" {
		err = errors.New("invalid reset otp")
		response.Message = "Invalid Reset OTP"
		return err
	}
	var user model.User
	err = db.Conn.First(&user, "email = ? OR mobile = ?", username, username).Error
	if err == nil {
		resetcode_on := "Email"
		if user.Email == username {
			resetcode_on = "Email"
		} else if user.Mobile == username {
			resetcode_on = "Mobile"
		}
		var passwordHash []byte
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err == nil {
			updated_flag := false
			password = string(passwordHash)
			switch resetcode_on {
			case "Email":
				if resetcode == user.EmailCode {
					user.EmailVerified = true
					user.EmailCode = ""
					user.EmailCodeSet = false
					user.Password = password
					updated_flag = true
					db.Conn.Save(user)
					response.Message = " New password set successfully "
					response.Status = apiresponse.SUCCESS
				}
			case "Mobile":
				if resetcode == user.EmailCode {
					user.MobileVerified = true
					user.MobileCode = ""
					user.MobileCodeSet = false
					user.Password = password
					updated_flag = true
				}
			}
			if updated_flag {
				db.Conn.Save(user)
				response.Message = " New password set successfully "
				response.Status = apiresponse.SUCCESS
			} else {
				response.Message = " Reset code not matched "
				response.Status = apiresponse.FAILED
			}
		} else {
			util.Log(err)
			response.Message = "Failed to encrypt the password"
			return err
		}
	} else {
		response.Message = "User Not Found, Please create acount first"
		response.Status = apiresponse.FAILED
	}
	return err
}
