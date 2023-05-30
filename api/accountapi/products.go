package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type productview struct {
	model.Product
	UpdatedByName string
}

var productsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postProducts := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postProduct := range postProducts {
			postProduct := _postProduct.(map[string]interface{})
			message := " Added "
			var item model.Product
			id := util.GetID(postProduct)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if (err != nil || id == 0) && response.Account.AccountType != "admin" {
				item = model.Product{
					AccountID: response.Account.ID,
				}
				err = nil
			}
			if err == nil {
				_, err = item.Update(postProduct, response.User)
			}
			if err != nil {
				if message == " Added " {
					message = "failed to add (" + err.Error() + ")"
				} else {
					message = "failed to update (" + err.Error() + ")"
				}
			} else {
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
			}
			count := messages[message]
			messages[message] = count + 1
		}

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + "Products "
			} else {
				response.Message += "Product "
			}
			response.Message += key
		}
	} else if _, okDelete := response.Request["todelete"]; okDelete {

		todelete := response.Request["todelete"].([]interface{})
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			susccessMessage := ""
			if itemid > 0 {
				var product model.Product
				err = db.Conn.First(&product, "id = ?", itemid).Error
				if product.AccountID != response.Account.ID {
					err = errors.New("you are not authorized to delete this record")
				}
				if err == nil {
					err = db.Conn.Delete(&product).Error
					if err == nil {
						susccessMessage = " Product deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Products not found"
				}
			} else {
				invalidIDMessage = " Some Product ids are invalid "
			}
			response.Message = ""
			if susccessMessage != "" {
				response.Message += susccessMessage
			}
			if errorMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += errorMessage
			}
			if invalidIDMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += invalidIDMessage
			}
		}
	}

	if err == nil {
		var products []productview

		var query = db.Conn.Model(&model.Product{}).Select("products.*, users.name AS updated_by_name").Joins("left join users on products.updated_by = users.id")
		if response.Account.AccountType != "admin" {
			query.Where("products.account_id = ?", response.Account.ID)
		}
		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["Service"]; ok {
				Service := util.GetBool(fixConditions["Service"])
				query.Where("products.service = ?", Service)
			}
		}
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" products.name Like ? OR products.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
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
	}
	return err
}
