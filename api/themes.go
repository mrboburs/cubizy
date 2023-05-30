package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"fmt"
	"net/http"
)

var themesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var themes []model.Theme
	var query = db.Conn.Model(&model.Theme{})
	switch response.Account.AccountType {
	case "admin":
		query.Where(" admin = true")
	case "seller":
		query.Where(" seller = true")
	}
	query.Count(&response.RecordsTotal)
	response.RecordsFiltered = response.RecordsTotal
	if Search, okSearch := response.Request["search"]; okSearch {
		SearchString := fmt.Sprintf("%v", Search)
		if SearchString != "" {
			SearchStringLike := "%" + SearchString + "%"
			query.Where(" themes.title Like", SearchStringLike)
			//response.RecordsFiltered = 3
			query.Count(&response.RecordsFiltered)
		}
	}
	db.SetUpQuery(response.Request, query)
	if response.RecordsFiltered > 0 {
		err = query.Scan(&themes).Error
	}
	if err == nil {
		response.Data = themes
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.Theme{}
		response.Status = apiresponse.FAILED
	}
	return err
}
