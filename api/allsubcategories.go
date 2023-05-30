package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

var allsubcategoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var subcategories []model.Subcategory

	var query = db.Conn.Model(&model.Subcategory{})
	var flag = false
	if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
		fixConditions := response.Request["fix_condition"].(map[string]interface{})
		if _, ok := fixConditions["BaseSubcategoryID"]; ok {
			BaseSubcategoryID := util.GetUint(fixConditions["BaseSubcategoryID"])
			query.Where(" base_subcategory_id = ? ", BaseSubcategoryID)
			flag = true
		}
		if _, ok := fixConditions["CategoryID"]; ok {
			CategoryID := util.GetUint(fixConditions["CategoryID"])
			query.Where(" category_id = ? ", CategoryID)
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
			query.Where(" subcategories.name Like ? OR subcategories.description Like ?", SearchStringLike, SearchStringLike)
			//response.RecordsFiltered = 3
			query.Count(&response.RecordsFiltered)
		}
	}

	db.SetUpQuery(response.Request, query)
	if response.RecordsFiltered > 0 {
		err = query.Scan(&subcategories).Error
	}
	if err == nil {
		response.Data = subcategories
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.Subcategory{}
		response.Status = apiresponse.FAILED
	}
	return err
}
