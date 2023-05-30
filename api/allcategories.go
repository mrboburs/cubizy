package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

var allcategoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var categories []model.Category

	var query = db.Conn.Model(&model.Category{})

	var flag = false
	if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
		fixConditions := response.Request["fix_condition"].(map[string]interface{})
		if _, ok := fixConditions["BaseCategoryID"]; ok {
			BaseCategoryID := util.GetUint(fixConditions["BaseCategoryID"])
			query.Where(" base_category_id = ? ", BaseCategoryID)
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
			query.Where(" categories.name Like ? OR categories.description Like ?", SearchStringLike, SearchStringLike)
			//response.RecordsFiltered = 3
			query.Count(&response.RecordsFiltered)
		}
	}

	db.SetUpQuery(response.Request, query)
	if response.RecordsFiltered > 0 {
		err = query.Scan(&categories).Error
	}
	if err == nil {
		response.Data = categories
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.Category{}
		response.Status = apiresponse.FAILED
	}
	return err
}
