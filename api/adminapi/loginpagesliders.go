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

type loginpagesliderview struct {
	model.Loginpageslider
	UpdatedByName string
}

var loginpageslidersAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postLoginpagesliders := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postLoginpageslider := range postLoginpagesliders {
			postLoginpageslider := _postLoginpageslider.(map[string]interface{})
			message := " Added "
			var item model.Loginpageslider
			id := util.GetID(postLoginpageslider)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Loginpageslider{
					AccountID: response.Account.ID,
				}
				err = nil
			}
			_, err = item.Update(postLoginpageslider, response.User)
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
				response.Message += strconv.Itoa(value) + "Loginpagesliders "
			} else {
				response.Message += "Loginpageslider "
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
				var loginpageslider model.Loginpageslider
				err = db.Conn.First(&loginpageslider, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&loginpageslider).Error
					if err == nil {
						susccessMessage = " Loginpageslider deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Loginpagesliders not found"
				}
			} else {
				invalidIDMessage = " Some Loginpageslider ids are invalid "
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
		var loginpagesliders []loginpagesliderview

		var query = db.Conn.Model(&model.Loginpageslider{}).Select("loginpagesliders.*, users.name AS updated_by_name").Joins("left join users on loginpagesliders.updated_by = users.id")
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" loginpagesliders.title Like ? OR users.name Like ? OR loginpagesliders.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&loginpagesliders).Error
		}
		if err == nil {
			response.Data = loginpagesliders
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Loginpageslider{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
