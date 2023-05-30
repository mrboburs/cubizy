package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
)

var accountaddressAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var accountaddress *model.Address
	if response.Account.AddressID > 0 {
		err = db.Conn.First(&accountaddress, response.Account.AddressID).Error
		if err == nil {
			response.Result["address"] = accountaddress
			response.Status = apiresponse.SUCCESS
		}
	}
	if _, ok := response.Request["address"]; ok {
		postAddress := response.Request["address"].(map[string]interface{})

		if accountaddress == nil {
			accountaddress = &model.Address{
				AccountID: response.Account.ID,
			}
		}
		_, err = accountaddress.Update(postAddress, response.User)
		if err == nil {

			if response.Account.AddressID == 0 {
				response.Account.AddressID = accountaddress.ID
				db.Conn.Save(response.Account)
			}

			response.Message = "Account address updated"
			response.Result["address"] = accountaddress
			response.Status = apiresponse.SUCCESS
		} else {
			response.Message = "Failed to update Account address"
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
