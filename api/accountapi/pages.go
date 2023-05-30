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

type pageview struct {
	model.Page
	UpdatedByName string
}

var pagesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if response.Account != nil {
		if _, ok := response.Request["items"]; ok {
			postPages := response.Request["items"].([]interface{})
			messages := make(map[string]int)
			for _, _postPage := range postPages {
				postPage := _postPage.(map[string]interface{})
				message := " Added "
				var item model.Page
				id := util.GetID(postPage)
				if id > 0 {
					err = db.Conn.First(&item, "id = ?", id).Error
					if err == nil {
						message = " Updated "
					}
				}
				if err != nil || id == 0 {
					item = model.Page{
						AccountID: response.Account.ID,
					}
					err = nil
				}
				_, err = item.Update(postPage, response.User)
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
					response.Message += strconv.Itoa(value) + "Pages "
				} else {
					response.Message += "Page "
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
					var page model.Page
					err = db.Conn.First(&page, "id = ?", itemid).Error
					if err == nil {
						err = db.Conn.Delete(&page).Error
						if err == nil {
							susccessMessage = " Page deleted "
							response.Status = apiresponse.SUCCESS
						} else {
							break
						}
					} else {
						errorMessage = "Some Pages not found"
					}
				} else {
					invalidIDMessage = " Some Page ids are invalid "
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
	}
	if err == nil {
		var pages []pageview

		var query = db.Conn.Model(&model.Page{}).Select("pages.*, users.name AS updated_by_name").Joins("left join users on pages.updated_by = users.id")
		query.Where(" pages.account_id = ? ", response.Account.ID)
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				util.Log("search query", SearchString)
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" pages.title Like ? OR users.name Like ? OR pages.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
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
