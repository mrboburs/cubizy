package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

var productsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var products []map[string]interface{}

	var query *gorm.DB

	if response.Account.AccountType == "admin" {
		query = db.Conn.Table("productsview")
	} else {
		query = db.Conn.Table("accountproductsview")
		query.Where("account_id = ?", response.Account.ID)
	}

	if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
		fixConditions := response.Request["fix_condition"].(map[string]interface{})
		if _, ok := fixConditions["Service"]; ok {
			Service := util.GetBool(fixConditions["Service"])
			query.Where("service = ?", Service)
		}
		// category_id
		if response.Account.AccountType == "admin" {
			if _, ok := fixConditions["category_id"]; ok {
				category_id := util.GetUint(fixConditions["category_id"])
				query.Where("base_category_id = ?", category_id)
			}
			if _, ok := fixConditions["subcategory_id"]; ok {
				subcategory_id := util.GetUint(fixConditions["subcategory_id"])
				query.Where("base_subcategory_id = ?", subcategory_id)
			}
			if _, ok := fixConditions["childcategory_id"]; ok {
				childcategory_id := util.GetUint(fixConditions["childcategory_id"])
				query.Where("base_childcategory_id = ?", childcategory_id)
			}
		} else {
			if _, ok := fixConditions["category_id"]; ok {
				category_id := util.GetUint(fixConditions["category_id"])
				query.Where("category_id = ?", category_id)
			}
			if _, ok := fixConditions["subcategory_id"]; ok {
				subcategory_id := util.GetUint(fixConditions["subcategory_id"])
				query.Where("subcategory_id = ?", subcategory_id)
			}
			if _, ok := fixConditions["childcategory_id"]; ok {
				childcategory_id := util.GetUint(fixConditions["childcategory_id"])
				query.Where("childcategory_id = ?", childcategory_id)
			}
		}

		if _, ok := fixConditions["max_cost"]; ok {
			max_cost := util.GetUint(fixConditions["max_cost"])
			query.Where("max_cost < ?", max_cost+1)
		}
		if _, ok := fixConditions["min_cost"]; ok {
			min_cost := util.GetUint(fixConditions["min_cost"])
			query.Where("min_cost > ?", min_cost+1)
		}
		if _, ok := fixConditions["product_rating"]; ok {
			product_rating := util.GetInt(fixConditions["product_rating"])
			query.Where("`rating` = ?", product_rating)
		}
		if _, ok := fixConditions["selectedlocation"]; ok {
			selectedlocation := util.GetString(fixConditions["selectedlocation"])
			query.Where("account_id IN (SELECT `addresses`.`account_id` FROM `addresses` WHERE `addresses`.`locality` = ?)", selectedlocation)
		}
		// attributes
		if _, ok := fixConditions["attributes"]; ok {
			attributes := fixConditions["attributes"].(map[string]interface{})
			for k, v := range attributes {
				if strings.Contains(k, " ") {
					break
				}
				valuesInterface := v.([]interface{})
				values := make([]string, len(valuesInterface))
				for i, v := range valuesInterface {
					values[i] = v.(string)
				}
				query.Where("`"+k+"` IN ?", values)
			}
		}

		if _, ok := fixConditions["id"]; ok {
			id := util.GetUint(fixConditions["id"])
			query.Where("id = ?", id)
		}

		if _, ok := fixConditions["ids"]; ok {
			ids := fixConditions["ids"].([]interface{})
			query.Where("id IN ?", ids)
		}
	}

	query.Debug().Count(&response.RecordsTotal)
	response.RecordsFiltered = response.RecordsTotal
	if Search, okSearch := response.Request["search"]; okSearch {
		SearchString := fmt.Sprintf("%v", Search)
		if SearchString != "" {
			SearchStringLike := "%" + SearchString + "%"
			query.Where(" name Like ? OR id Like ? ", SearchStringLike, SearchStringLike, SearchString)
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
	return err
}
