package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
)

var loginpageslidersAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var loginpagesliders []model.Loginpageslider

	var query = db.Conn.Model(&model.Loginpageslider{})
	query.Limit(5)
	query.Order("RAND()")
	err = query.Scan(&loginpagesliders).Error
	if err == nil {
		response.Data = loginpagesliders
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.Loginpageslider{}
		response.Status = apiresponse.FAILED
	}
	return err
}
