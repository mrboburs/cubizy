package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

var wishlistproductsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var products []map[string]interface{}

	var query *gorm.DB

	if response.Account.AccountType == "admin" {
		query = db.Conn.Table("productsview")
	} else {
		query = db.Conn.Table("accountproductsview")
		query.Where("account_id = ?", response.Account.ID)
	}

	if response.User == nil {
		if products, okSearch := response.Request["products"]; okSearch {
			postProducts := products.([]interface{})
			if len(postProducts) > 0 {
				query.Where(" id IN ? ", postProducts)
			} else {
				response.Data = []model.Product{}
				response.Status = apiresponse.FAILED
				return errors.New("empty request")
			}
		} else {
			response.Data = []model.Product{}
			response.Status = apiresponse.FAILED
			return errors.New("empty request")
		}
	} else {
		if products, okSearch := response.Request["products"]; okSearch {
			postProducts := products.([]interface{})
			if len(postProducts) > 0 {
				for _, postProduct := range postProducts {
					postWishItem := util.GetUint(postProduct)
					var item model.Wishlist
					if postWishItem > 0 {
						err = db.Conn.First(&item, "product_id = ? AND created_by = ?", postWishItem, response.User.ID).Error
						if err != nil {
							item = model.Wishlist{
								ProductID: postWishItem,
								CreatedBy: response.User.ID,
							}
							err = db.Conn.Create(&item).Error
							if err != nil {
								util.Log("While adding wishlists item ")
								util.Log(err)
								err = nil
							}
						}
					}
				}
				response.User.UpdateWishlistCount()
			}
		}
		query.Where(" id IN (SELECT product_id FROM wishlists WHERE created_by = ?) ", response.User.ID)
	}
	query.Count(&response.RecordsTotal)
	response.RecordsFiltered = response.RecordsTotal

	db.SetUpQuery(response.Request, query)
	if response.RecordsFiltered > 0 {
		err = query.Scan(&products).Error
	}
	if err == nil {
		response.Data = products
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.Product{}
		response.Status = apiresponse.FAILED
	}
	return err
}
