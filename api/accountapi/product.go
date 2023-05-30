package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var productAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var isupdated bool
	if _, ok := response.Request["product"]; ok {
		postProduct := response.Request["product"].(map[string]interface{})
		var item model.Product
		id := util.GetID(postProduct)
		if id > 0 {
			err = db.Conn.First(&item, "id = ?", id).Error
		}
		if err != nil || id == 0 {
			item = model.Product{
				AccountID: response.Account.ID,
			}
			err = nil
		}
		isupdated, err = item.Update(postProduct, response.User)
		if err != nil {
			if id == 0 {
				response.Message = "Failed to add product (" + err.Error() + ")"
			} else {
				response.Message = "Failed to update product (" + err.Error() + ")"
			}
		} else {
			response.Result["product"] = item
			var results []map[string]interface{}
			db.Conn.Table("productattributes").Find(&results, " product_id = ?", item.ID)
			if len(results) > 0 {
				response.Result["ProductDetails"] = results[0]
			}
			if response.Account.MaxPrice < item.MaxPrice {
				response.Account.MaxPrice = item.MaxPrice*10 - item.MaxPrice
				var t uint
				t = 10
				for t < item.MaxPrice {
					t *= 10
				}
				response.Account.MaxPrice = t
				db.Conn.Save(response.Account)
			}
			response.Status = apiresponse.SUCCESS
			if id == 0 {
				response.Message = "Product added"
			} else {
				response.Message = "Product updated"
			}
			if !isupdated {
				response.Message = ""
			}
		}

	}
	return err
}
