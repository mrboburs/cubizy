package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

var districtsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	var locations []model.Location

	var query = db.Conn.Model(&model.Location{})
	query.Distinct("district")
	//.Select("DISTINCT(`district`)")

	if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
		fixConditions := response.Request["fix_condition"].(map[string]interface{})
		if _, ok := fixConditions["code"]; ok {
			code := util.GetString(fixConditions["code"])
			query.Where(" code = ? ", code)
		}
		if _, ok := fixConditions["country"]; ok {
			country := util.GetString(fixConditions["country"])
			query.Where(" country = ? ", country)
		}
	}

	query.Count(&response.RecordsTotal)
	response.RecordsFiltered = response.RecordsTotal
	if Search, okSearch := response.Request["search"]; okSearch {
		SearchString := fmt.Sprintf("%v", Search)
		if SearchString != "" {
			SearchStringLike := "%" + SearchString + "%"
			query.Where(" locations.district Like ? ", SearchStringLike)
			//response.RecordsFiltered = 3
			query.Count(&response.RecordsFiltered)
		}
	}

	db.SetUpQuery(response.Request, query)
	if response.RecordsFiltered > 0 {
		err = query.Scan(&locations).Error
	}
	if err == nil {
		response.Data = locations
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.Location{}
		response.Status = apiresponse.FAILED
	}
	return err
}
