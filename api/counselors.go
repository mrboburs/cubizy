package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"fmt"
	"net/http"
)

var supportagentsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if response.Account != nil && response.Account.ID > 0 {
		var supportagents []model.User

		var query = db.Conn.Model(&model.User{}).Select("users.id, users.name, users.photo, users.last_active_on, users.supportagent_quote")
		query.Where(" users.is_supportagent = 1")
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal

		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" users.name Like ? OR users.email Like ? OR users.mobile Like ? OR users.supportagent_quote Like ? ", SearchStringLike, SearchStringLike, SearchStringLike, SearchStringLike)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}

		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&supportagents).Error
		}
		if err == nil {
			response.Data = supportagents
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.User{}
			response.Status = apiresponse.FAILED
		}
	} else {
		response.Message = "empty request"
	}
	return err
}
