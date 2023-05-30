package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

var productsuggestionsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var results []map[string]interface{}
	if Search, okSearch := response.Request["search"]; okSearch {
		SearchString := fmt.Sprintf("%v", Search)
		query := db.Conn.Model(&model.Product{})
		query.Select("id, name, category_id, subcategory_id, base_category_id, base_subcategory_id")

		if response.Account.AccountType != "admin" {
			query.Where(" account_id = ? ", response.Account.ID)
		}
		if _, ok := response.Request["location"]; ok {
			location := util.GetString(response.Request["location"])
			if len(location) > 0 {
				query.Where("account_id IN (SELECT `addresses`.`account_id` FROM `addresses` WHERE `addresses`.`locality` = ?)", location)
			}
		}

		if response.Account.AccountType == "admin" {
			if _, ok := response.Request["category_id"]; ok {
				category_id := util.GetUint(response.Request["category_id"])
				if category_id > 0 {
					query.Where("base_category_id = ?", category_id)
				}
			}
			if _, ok := response.Request["subcategory_id"]; ok {
				subcategory_id := util.GetUint(response.Request["subcategory_id"])
				if subcategory_id > 0 {
					query.Where("base_subcategory_id = ?", subcategory_id)
				}
			}
			if _, ok := response.Request["childcategory_id"]; ok {
				childcategory_id := util.GetUint(response.Request["childcategory_id"])
				if childcategory_id > 0 {
					query.Where("base_childcategory_id = ?", childcategory_id)
				}
			}
		} else {
			if _, ok := response.Request["category_id"]; ok {
				category_id := util.GetUint(response.Request["category_id"])
				if category_id > 0 {
					query.Where("category_id = ?", category_id)
				}
			}
			if _, ok := response.Request["subcategory_id"]; ok {
				subcategory_id := util.GetUint(response.Request["subcategory_id"])
				if subcategory_id > 0 {
					query.Where("subcategory_id = ?", subcategory_id)
				}
			}
			if _, ok := response.Request["childcategory_id"]; ok {
				childcategory_id := util.GetUint(response.Request["childcategory_id"])
				if childcategory_id > 0 {
					query.Where("childcategory_id = ?", childcategory_id)
				}
			}
		}

		if SearchString != "" {
			SearchStringLike := "%" + SearchString + "%"
			err = query.Where(" deleted_at IS NULL AND status = 1 AND name LIKE ? ", SearchStringLike).Limit(100).Find(&results).Error
		}
	}
	//time.Sleep(30 * time.Second)
	if err == nil {
		response.Data = results
		response.Status = apiresponse.SUCCESS
	}
	return err
}
