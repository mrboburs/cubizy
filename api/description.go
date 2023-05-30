package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var descriptionAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["ProductID"]; ok {
		var item model.Description
		product_id := util.GetUint(response.Request["ProductID"])
		if product_id > 0 {
			err = db.Conn.First(&item, "product_id = ?", product_id).Error
		}
		if err == nil {
			response.Result["Content"] = item.Content
			response.Status = apiresponse.SUCCESS
			response.Message = ""
		}

	}
	return err
}
