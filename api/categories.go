package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
)

var categoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if response.Account != nil {
		accountID := response.Account.ID
		var categories []model.Category
		err = db.Conn.Where("active = true && account_id = ?", accountID).Order("`top` desc").Debug().Find(&categories).Error
		if err == nil {
			response.Result["categories"] = categories

			var subcategories []model.Subcategory
			err = db.Conn.Where("active = true && account_id = ?", accountID).Order("id desc").Find(&subcategories).Error
			if err == nil {
				response.Result["subcategories"] = subcategories

				var childcategories []model.Childcategory
				err = db.Conn.Where("active = true && account_id = ?", accountID).Order("id desc").Find(&childcategories).Error
				if err == nil {
					response.Result["childcategories"] = childcategories
				}
			}
		}

		if err == nil {
			response.Status = apiresponse.SUCCESS
		}
	} else {
		response.Status = apiresponse.FAILED
		response.Message = "We dont have account for subdomain : " + response.Domain
	}
	return err
}
