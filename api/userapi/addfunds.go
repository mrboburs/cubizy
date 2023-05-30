package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"errors"
	"net/http"
)

var addfundsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["item"]; ok {
		postTransaction := response.Request["item"].(map[string]interface{})
		item := model.Transaction{
			AccountID: model.SuperAdminAccount.ID,
			CreatedBy: response.User.ID,
		}
		_, err = item.Update(postTransaction, response.User)
		if err != nil {
			response.Message = "failed to add (" + err.Error() + ")"
		} else {
			response.Result["transaction"] = item
			response.Status = apiresponse.SUCCESS
			if item.Accepted {
				response.Message = "Fund added to your wallet"
			} else {
				response.Message = "Request added to add funds , please wait while we verify and accept the request, it may take from 30 minutes to up to 1 day"
			}
		}
	} else {
		response.Message += "Empty request "
		err = errors.New("items not provided")
	}
	return err
}
