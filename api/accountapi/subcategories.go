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

type subcategoryview struct {
	model.Subcategory
	UpdatedByName string
}

var subcategoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postSubcategories := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postSubcategory := range postSubcategories {
			postSubcategory := _postSubcategory.(map[string]interface{})
			message := " Added "
			var item model.Subcategory
			id := util.GetID(postSubcategory)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Subcategory{
					AccountID: response.Account.ID,
					CreatedBy: response.User.ID,
				}
				err = nil
			}
			_, err = item.Update(postSubcategory, response.User)
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
				response.Message += strconv.Itoa(value) + "Subcategories "
			} else {
				response.Message += "Subcategory "
			}
			response.Message += key
		}
	} else if _, okDelete := response.Request["todelete"]; okDelete {
		var category_id uint = 0
		todelete := response.Request["todelete"].([]interface{})
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			susccessMessage := ""
			if itemid > 0 {
				var subcategory model.Subcategory
				err = db.Conn.First(&subcategory, "id = ?", itemid).Error
				if err == nil {
					category_id = subcategory.CategoryID
					err = db.Conn.Delete(&subcategory).Error
					if err == nil {
						susccessMessage = " Subcategory deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Subcategories not found"
				}
			} else {
				invalidIDMessage = " Some Subcategory ids are invalid "
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
		model.UpdateSubcategoryCount(category_id)
	}

	if err == nil {
		var subcategories []subcategoryview

		var query = db.Conn.Model(&model.Subcategory{}).Select("subcategories.*, users.name AS updated_by_name").Joins("left join users on subcategories.updated_by = users.id")

		query.Where(" account_id = ? ", response.Account.ID)
		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["category_id"]; ok {
				category_id := util.GetUint(fixConditions["category_id"])
				query.Where(" category_id = ? ", category_id)
			}
		}

		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" subcategories.title Like ? OR users.name Like ? OR subcategories.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}

		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&subcategories).Error
		}
		if err == nil {
			response.Data = subcategories
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Subcategory{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
