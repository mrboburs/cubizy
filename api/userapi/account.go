package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"errors"
	"net/http"
)

var accountAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["account"]; ok {
		postAccount := response.Request["account"].(map[string]interface{})

		var account model.Account
		var account_type = response.Domain

		if account_type == "seller" || account_type == "admin" {
			if response.Account == nil {
				response.Message = " Created "
				account = model.Account{}
				account.AccountType = account_type
				account.CreatedBy = response.User.ID
			} else {
				response.Message = " Updated "
				account = *response.Account
			}
			_, err = account.Update(postAccount, response.User)
			if err != nil {
				if response.Message == " Created " {
					response.User.SellerAccountID = account.ID
					db.Conn.Save(response.User)
					response.Message = "Failed to create Account (" + err.Error() + ")"
				} else {
					response.Message = "Failed to update Account(" + err.Error() + ")"
				}
			} else {
				response.Account = &account
				response.Message = "Account " + response.Message
				response.Status = apiresponse.SUCCESS
			}
		} else {
			err = errors.New("invalid account type : " + account_type)
			response.Message = "Invalid account type" + account_type
		}
	}
	return err
}
