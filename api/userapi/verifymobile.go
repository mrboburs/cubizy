package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var verifymobileAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["Mobile"]; ok {
		Mobile := util.GetString(response.Request["Mobile"])
		if response.User.Mobile == Mobile {
			if _, ok := response.Request["MobileCode"]; ok {
				MobileCode := util.GetString(response.Request["MobileCode"])
				if response.User.MobileCode == MobileCode {
					response.User.MobileCode = ""
					response.User.MobileCodeSet = false
					response.User.MobileVerified = true
					db.Conn.Save(response.User)
					response.Message = " Mobile Verified "
					response.Status = apiresponse.SUCCESS
				} else {
					response.Message = " Wrong Code "
				}
			}
		} else {
			response.Message = " Invalid mobile "
		}
	}
	return err
}
