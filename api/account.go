package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var accountAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["id"]; ok {
		account_id := util.GetUint(response.Request["id"])
		var account model.Account
		err = db.Conn.First(&account, " id = ? ", account_id).Error
		if err == nil {
			response.Result["account"] = account

			if err == nil {
				response.Status = apiresponse.SUCCESS
				response.Message = ""
			}
		} else {
			response.Message = err.Error()
		}
	} else {
		response.Message = "Account not found"
	}
	return err
}
