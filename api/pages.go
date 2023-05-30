package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
)

var pagesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if response.Account != nil && response.Account.ID > 0 {
		var pages []model.Page

		var query = db.Conn.Model(&model.Page{}).Select("id, title")
		query.Where(" status = 1  AND account_id = ? ", response.Account.ID)
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&pages).Error
		}
		if err == nil {
			response.Data = pages
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Page{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
