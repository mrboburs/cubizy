package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"fmt"
	"net/http"
)

var locationsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var locations []model.Location

	var query = db.Conn.Model(&model.Location{})

	if response.Account.AccountType != "admin" && response.User == nil {
		query.Where(" id IN (SELECT location_id FROM addresses WHERE account_id = ? )", response.Account.ID)
	}

	query.Count(&response.RecordsTotal)
	response.RecordsFiltered = response.RecordsTotal
	if Search, okSearch := response.Request["search"]; okSearch {
		SearchString := fmt.Sprintf("%v", Search)
		if SearchString != "" {
			SearchStringLike := "%" + SearchString + "%"
			query.Where(" locations.code Like ? OR locations.sub_locality Like ? OR locations.locality Like ? OR locations.country Like ? ", SearchStringLike, SearchStringLike, SearchStringLike, SearchStringLike)
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
