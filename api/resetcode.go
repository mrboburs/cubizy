package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/clicksend"
	"cubizy/plugins/db"
	"cubizy/plugins/sendgrid"
	"cubizy/plugins/srand"
	"cubizy/util"
	"net/http"
)

var resetcodeAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	username := "Invalid"
	if _, ok := response.Request["username"]; ok {
		username = util.GetString(response.Request["username"])
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
		response.Message = "Will send code on  " + resetcode_on
		switch resetcode_on {
		case "Email":
			if user.EmailCode == "" {
				user.EmailCode = srand.String(8)
				user.EmailCodeSet = true
				db.Conn.Save(user)
			}
			resetPasswordLink := r.Header.Get("Origin") + "/auth/recoverpw?code=" + user.EmailCode + "&email=" + user.Email
			emailContent := sendgrid.ResetPasswordTemplate(user.Name, user.Email, user.EmailCode, resetPasswordLink)

			plainTextContent := "Reset Password Code : " + user.EmailCode
			err = sendgrid.SendEmail(user.Name, user.Email, " Code to reset password on "+util.Settings.AppName+" ", plainTextContent, emailContent)
			if err == nil {
				response.Message = " Please check reset code and link sent on your email to reset your password  "
				response.Status = apiresponse.SUCCESS
			} else {
				response.Message = "Failed to send email(" + err.Error() + ")"
			}
		case "Mobile":
			user.MobileCode = srand.Number(8)
			user.MobileCodeSet = true
			db.Conn.Save(user)
			err = clicksend.SendSMS(user.Mobile, "Password reset code for your account on "+util.Settings.AppName+" account panel is :"+user.MobileCode)
			if err == nil {
				response.Status = apiresponse.SUCCESS
				response.Message = "SMS sent to you phone, please check"
			} else {
				response.Message = "Failed to send sms(" + err.Error() + ")"
				response.Status = apiresponse.FAILED
			}
		}
	} else {
		response.Message = "User Not Found, Please create acount first"
		response.Status = apiresponse.FAILED
	}
	return err
}
