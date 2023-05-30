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

var accountsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postAccounts := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postAccount := range postAccounts {
			postAccount := _postAccount.(map[string]interface{})
			message := " Added "
			var item model.Account
			id := util.GetID(postAccount)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				continue
			}
			_, err = item.Update(postAccount, response.User)
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
				response.Message += strconv.Itoa(value) + "Accounts "
			} else {
				response.Message += "Account "
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
				var account model.Account
				err = db.Conn.First(&account, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&account).Error
					if err == nil {
						susccessMessage = " Account deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Accounts not found"
				}
			} else {
				invalidIDMessage = " Some Account ids are invalid "
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
		var accounts []model.AdminAccountView

		var query = db.Conn.Model(&model.AdminAccountView{})

		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["status"]; ok {
				status := util.GetInt(fixConditions["status"])
				query.Where(" status = ?", status)
			}
		}
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" title Like ? OR name Like ?", SearchStringLike, SearchStringLike)
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&accounts).Error
		}
		if err == nil {
			response.Data = accounts
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.AdminAccountView{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
