package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var verifyemailAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["Email"]; ok {
		Email := util.GetString(response.Request["Email"])
		if response.User.Email == Email {
			if _, ok := response.Request["EmailCode"]; ok {
				EmailCode := util.GetString(response.Request["EmailCode"])
				if response.User.EmailCode == EmailCode {
					response.User.EmailCode = ""
					response.User.EmailCodeSet = false
					response.User.EmailVerified = true
					db.Conn.Save(response.User)
					response.Message = " Email Verified "
					response.Status = apiresponse.SUCCESS
				} else {
					response.Message = " Wrong Code "
				}
			}
		} else {
			response.Message = " Invalid email "
		}
	}
	return err
}
