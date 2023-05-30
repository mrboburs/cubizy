package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"net/http"
)

var meAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var flag bool
	if _, ok := response.Request["user"]; ok {
		postUser := response.Request["user"].(map[string]interface{})
		flag, err = response.User.Update(postUser, response.User)
		if err == nil && response.Account != nil && response.Account.CreatedBy == response.User.ID && response.Account.Status == 0 {
			response.Account.Status = 1
			err = db.Conn.Save(&response.Account).Error
			if err == nil {
				flag = true
			}
		}
		if err == nil {
			if !flag {
				response.Message = "Nothing changed"
			} else {
				response.Message = " Updated your profile"
				response.Status = apiresponse.SUCCESS
			}
		} else {
			response.Message = " Can not update your profile (" + err.Error() + ")"
		}
	} else {
		response.Message = "Empty request"
	}
	return err
}
