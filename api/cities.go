package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

var citiesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	//time.Sleep(10 * time.Second)

	var results []string
	if _city, ok := response.Request["city"]; ok {
		city := fmt.Sprintf("%v", _city)
		limit := 100
		if _, ok := response.Request["limit"]; ok {
			limit = util.GetInt(response.Request["limit"])
		}
		err = db.Conn.Model(&model.Location{}).Where("locality LIKE ? ", "%"+city+"%").Distinct().Limit(limit).Pluck("locality", &results).Error
		if err == nil {
			response.Data = results
			response.Status = apiresponse.SUCCESS
		}
	}
	return err
}
