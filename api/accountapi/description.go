package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var descriptionAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var isupdated bool
	if _, ok := response.Request["ProductID"]; ok {
		var item model.Description
		product_id := util.GetUint(response.Request["ProductID"])
		if product_id > 0 {
			err = db.Conn.First(&item, "product_id = ?", product_id).Error
		}
		if err != nil && product_id != 0 {
			item = model.Description{
				AccountID: response.Account.ID,
				ProductID: product_id,
				Content:   "",
			}
			err = nil
		}
		isupdated, err = item.Update(response.Request, response.User)
		if err != nil {
			if product_id == 0 {
				response.Message = "Failed to add description (" + err.Error() + ")"
			} else {
				response.Message = "Failed to update description (" + err.Error() + ")"
			}
		} else {
			response.Result["Content"] = item.Content
			response.Status = apiresponse.SUCCESS
			if product_id == 0 {
				response.Message = "Description added"
			} else {
				response.Message = "Description updated"
			}
			if !isupdated {
				response.Message = ""
			}
		}

	}
	return err
}
