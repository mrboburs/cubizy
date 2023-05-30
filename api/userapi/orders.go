package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
)

var ordersAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if response.Account != nil && response.Account.ID > 0 {
		var orders []model.Order

		var query = db.Conn.Model(&model.Order{})
		query.Where(" created_by = ?  AND account_id = ? ", response.User.ID, response.Account.ID)
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal

		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&orders).Error
		}
		if err == nil {
			response.Data = orders
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Order{}
			response.Status = apiresponse.FAILED
		}
	} else {
		response.Message = "empty request"
	}
	return err
}
