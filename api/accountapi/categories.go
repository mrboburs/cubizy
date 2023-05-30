package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
	"strconv"
)

type categoryview struct {
	model.Category
	UpdatedByName string
}

var categoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postCategories := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postCategory := range postCategories {
			postCategory := _postCategory.(map[string]interface{})
			message := " Added "
			var item model.Category
			id := util.GetID(postCategory)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Category{
					AccountID: response.Account.ID,
					CreatedBy: response.User.ID,
				}
				err = nil
			}
			_, err = item.Update(postCategory, response.User)
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
				response.Message += strconv.Itoa(value) + "Categories "
			} else {
				response.Message += "Category "
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
				var category model.Category
				err = db.Conn.First(&category, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&category).Error
					if err == nil {
						susccessMessage = " Category deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Categories not found"
				}
			} else {
				invalidIDMessage = " Some Category ids are invalid "
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
		var categories []categoryview

		var query = db.Conn.Model(&model.Category{}).Select("categories.*, users.name AS updated_by_name").Joins("left join users on categories.updated_by = users.id")
		query.Where(" account_id = ? ", response.Account.ID)
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" categories.title Like ? OR users.name Like ? OR categories.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&categories).Error
		}
		if err == nil {
			response.Data = categories
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Category{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
