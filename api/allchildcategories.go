package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

var allchildcategoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var childcategories []model.Childcategory

	var query = db.Conn.Model(&model.Childcategory{})

	var flag = false
	if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
		fixConditions := response.Request["fix_condition"].(map[string]interface{})
		if _, ok := fixConditions["BaseChildcategoryID"]; ok {
			BaseChildcategoryID := util.GetUint(fixConditions["BaseChildcategoryID"])
			query.Where(" base_childcategory_id = ? ", BaseChildcategoryID)
			flag = true
		}
		if _, ok := fixConditions["SubcategoryID"]; ok {
			SubcategoryID := util.GetUint(fixConditions["SubcategoryID"])
			query.Where(" subcategory_id = ? ", SubcategoryID)
			flag = true
		}
	}

	if !flag {
		query.Where(" account_id = ? ", response.Account.ID)
	}
	query.Count(&response.RecordsTotal)
	response.RecordsFiltered = response.RecordsTotal
	if Search, okSearch := response.Request["search"]; okSearch {
		SearchString := fmt.Sprintf("%v", Search)
		if SearchString != "" {
			SearchStringLike := "%" + SearchString + "%"
			query.Where(" childcategories.name Like ? OR childcategories.description Like ?", SearchStringLike, SearchStringLike)
			//response.RecordsFiltered = 3
			query.Count(&response.RecordsFiltered)
		}
	}

	db.SetUpQuery(response.Request, query)
	if response.RecordsFiltered > 0 {
		err = query.Scan(&childcategories).Error
	}
	if err == nil {
		response.Data = childcategories
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.Childcategory{}
		response.Status = apiresponse.FAILED
	}
	return err
}
