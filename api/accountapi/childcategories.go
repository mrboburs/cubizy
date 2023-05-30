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

type childcategoryview struct {
	model.Childcategory
	UpdatedByName string
}

var childcategoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postChildcategories := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postChildcategory := range postChildcategories {
			postChildcategory := _postChildcategory.(map[string]interface{})
			message := " Added "
			var item model.Childcategory
			id := util.GetID(postChildcategory)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Childcategory{
					AccountID: response.Account.ID,
					CreatedBy: response.User.ID,
				}
				err = nil
			}
			_, err = item.Update(postChildcategory, response.User)
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
				response.Message += strconv.Itoa(value) + "Childcategories "
			} else {
				response.Message += "Childcategory "
			}
			response.Message += key
		}
	} else if _, okDelete := response.Request["todelete"]; okDelete {
		var subcategory_id uint = 0
		todelete := response.Request["todelete"].([]interface{})
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			susccessMessage := ""
			if itemid > 0 {
				var childcategory model.Childcategory
				err = db.Conn.First(&childcategory, "id = ?", itemid).Error
				if err == nil {
					subcategory_id = childcategory.SubcategoryID
					err = db.Conn.Delete(&childcategory).Error
					if err == nil {
						susccessMessage = " Childcategory deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Childcategories not found"
				}
			} else {
				invalidIDMessage = " Some Childcategory ids are invalid "
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
		model.UpdateChildcategoryCount(subcategory_id)
	}

	if err == nil {
		var childcategories []childcategoryview

		var query = db.Conn.Model(&model.Childcategory{}).Select("childcategories.*, users.name AS updated_by_name").Joins("left join users on childcategories.updated_by = users.id")

		query.Where(" account_id = ? ", response.Account.ID)
		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["subcategory_id"]; ok {
				subcategory_id := util.GetUint(fixConditions["subcategory_id"])
				query.Where(" subcategory_id = ? ", subcategory_id)
			}
		}

		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" childcategories.title Like ? OR users.name Like ? OR childcategories.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}

		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&childcategories).Error
		}
		if err == nil {
			response.Data = childcategories
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Childcategory{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
