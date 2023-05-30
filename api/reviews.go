package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

type reviewview struct {
	model.Review
	Name  string
	Photo string
}

var reviewsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	var reviews []reviewview

	var query = db.Conn.Model(&model.Review{}).Select("reviews.*, users.name, users.photo").Joins("left join users on reviews.created_by = users.id")

	var valid = false
	if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {

		fixConditions := response.Request["fix_condition"].(map[string]interface{})
		if _, ok := fixConditions["account_id"]; ok {
			account_id := util.GetUint(fixConditions["account_id"])
			if account_id > 0 {
				valid = true
				query.Where(" account_id = ? ", account_id)
			}
		}
		if _, ok := fixConditions["product_id"]; ok {
			product_id := util.GetUint(fixConditions["product_id"])
			if product_id > 0 {
				valid = true
				query.Where(" product_id = ? ", product_id)
			}
		}
	}

	if !valid {
		response.Data = []reviewview{}
		response.Status = apiresponse.SUCCESS
		return err
	}

	query.Count(&response.RecordsTotal)
	response.RecordsFiltered = response.RecordsTotal

	response.Request["sort"] = "created_at"
	response.Request["sortdesc"] = true
	db.SetUpQuery(response.Request, query)
	if response.RecordsFiltered > 0 {
		err = query.Scan(&reviews).Error
	}
	if err == nil {
		response.Data = reviews
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []reviewview{}
		response.Status = apiresponse.FAILED
	}
	return err
}
