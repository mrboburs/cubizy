package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/clicksend"
	"cubizy/plugins/db"
	"cubizy/plugins/srand"
	"cubizy/util"
	"net/http"
)

var sendvirificationmobileAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["Mobile"]; ok {
		Mobile := util.GetString(response.Request["Mobile"])
		if response.User.Mobile != Mobile {
			response.User.Mobile = Mobile
			response.User.MobileCode = ""
			response.User.MobileCodeSet = false
		}
	}

	response.User.MobileCode = srand.Number(8)
	response.User.MobileCodeSet = true
	db.Conn.Save(response.User)
	err = clicksend.SendSMS(response.User.Mobile, "Verification code for registretion on "+util.Settings.AppName+" account panel is :"+response.User.MobileCode)
	if err == nil {
		response.Status = apiresponse.SUCCESS
		response.Message = "SMS sent to you phone, please check"
	} else {
		response.Message = "Failed to send sms(" + err.Error() + ")"
	}
	return err
}
