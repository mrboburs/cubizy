package adminapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
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
			if err != nil || id == 0 {
				item = model.Review{
					AccountID: response.Account.ID,
				}
				err = nil
			}
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

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + "Reviews "
			} else {
				response.Message += "Review "
			}
			response.Message += key
		}
	} else if _, okDelete := response.Request["todelete"]; okDelete {

		todelete := response.Request["todelete"].([]interface{})
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			susccessMessage := ""
			if itemid > 0 {
				var review model.Review
				err = db.Conn.First(&review, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&review).Error
					if err == nil {
						susccessMessage = " Review deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Reviews not found"
				}
			} else {
				invalidIDMessage = " Some Review ids are invalid "
			}
			response.Message = ""
			if susccessMessage != "" {
				response.Message += susccessMessage
			}
			if errorMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += errorMessage
			}
			if invalidIDMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += invalidIDMessage
			}
		}
	}

	if err == nil {
		var reviews []reviewview

		var query = db.Conn.Model(&model.Review{}).Select("reviews.*, users.name AS created_by_name, users.photo, products.name, products.logo").Joins("left join users on reviews.created_by = users.id left join products on reviews.product_id = products.id")
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" reviews.title Like ? OR users.name Like ? OR reviews.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&reviews).Error
		}
		if err == nil {
			response.Data = reviews
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Review{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
