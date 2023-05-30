package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"cubizy/plugins/sendgrid"
	"cubizy/plugins/srand"
	"cubizy/util"
	"net/http"
)

var sendvirificationemailAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["Email"]; ok {
		Email := util.GetString(response.Request["Email"])
		if response.User.Email != Email {
			response.User.Email = Email
			response.User.EmailCode = ""
			response.User.EmailCodeSet = false
		}
	}
	response.User.EmailCode = srand.String(8)
	response.User.EmailCodeSet = true
	db.Conn.Save(response.User)
	resetPasswordLink := r.Header.Get("Origin") + "/auth/login?code=" + response.User.EmailCode + "&email=" + response.User.Email
	emailContent := sendgrid.VerificationEmailTemplate(response.User.Name, response.User.Email, response.User.EmailCode, resetPasswordLink)

	plainTextContent := "Email Verification Code : " + response.User.EmailCode
	err = sendgrid.SendEmail(response.User.Name, response.User.Email, " Code to verify email on "+util.Settings.AppName+" ", plainTextContent, emailContent)
	if err == nil {
		response.Message = " Please check verification code or link sent on your email to verify your email  "
		response.Status = apiresponse.SUCCESS
	} else {
		response.Message = "Failed to send email(" + err.Error() + ")"
	}
	return err
}
