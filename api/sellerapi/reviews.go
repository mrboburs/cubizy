package sellerapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
	"strconv"
)

type reviewview struct {
	model.Review
	Logo          string
	Name          string
	Photo         string
	CreatedByName string
}

var reviewsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if _, ok := response.Request["items"]; ok {
		postReviews := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postReview := range postReviews {
			postReview := _postReview.(map[string]interface{})
			message := " Added "
			var item model.Review
			id := util.GetID(postReview)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err == nil || id > 0 {
				_, err = item.Update(postReview, response.User)
				if err != nil {
					if message == " Added " {
						message = "failed to add (" + err.Error() + ")"
					} else {
						message = "failed to update (" + err.Error() + ")"
					}
				}
				count := messages[message]
				messages[message] = count + 1
			}
		}

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + "Blogs "
			} else {
				response.Message += "Blog "
			}
			response.Message += key
		}
	}

	if err == nil {
		var reviews []reviewview

		var query = db.Conn.Model(&model.Review{})
		query.Where(" reviews.account_id = ? ", response.Account.ID)
		var product_set = false
		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {

			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["product_id"]; ok {
				product_id := util.GetUint(fixConditions["product_id"])
				if product_id > 0 {
					product_set = true
					query.Where(" product_id = ? ", product_id)
				}
			}
		}
		if !product_set {
			query.Select("reviews.*, products.name, products.logo")
			query.Joins("left join products on reviews.product_id = products.id")
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
	}
	return err
}
