package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var logoutAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if response.User != nil {
		if online_user, ok := apiresponse.OnlinUsers[response.User.ID]; ok {
			if online_user.WS != nil {
				err = online_user.WS.Close()
				if err != nil {
					util.Log(err)
				}
			}
			delete(apiresponse.OnlinUsers, response.User.ID)
			response.User.Online = false
		}
		response.User.Token = ""
		response.User.LastToken = ""
		response.User.Remember = false
		err = db.Conn.Save(&response.User).Error
		response.User = nil
		response.Account = nil
		response.Message = "Loged Out"
		response.Status = apiresponse.SUCCESS
	} else {
		response.Message = "Already Loged Out"
	}
	return err
}
