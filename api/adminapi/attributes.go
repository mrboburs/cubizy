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

type attributeview struct {
	model.Attribute
	UpdatedByName string
}

var attributesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postAttributes := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postAttribute := range postAttributes {
			postAttribute := _postAttribute.(map[string]interface{})
			message := " Added "
			var item model.Attribute
			id := util.GetID(postAttribute)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Attribute{}
				err = nil
			}
			_, err = item.Update(postAttribute, response.User)
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
				response.Message += strconv.Itoa(value) + "Attributes "
			} else {
				response.Message += "Attribute "
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
				var attribute model.Attribute
				err = db.Conn.First(&attribute, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&attribute).Error
					if err == nil {
						susccessMessage = " Attribute deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Attributes not found"
				}
			} else {
				invalidIDMessage = " Some Attribute ids are invalid "
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
		var attributes []attributeview

		var query = db.Conn.Model(&model.Attribute{}).Select("attributes.*, users.name AS updated_by_name").Joins("left join users on attributes.updated_by = users.id")

		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			var SubcategoryID, ChildcategoryID uint
			if _, ok := fixConditions["subcategory_id"]; ok {
				SubcategoryID = util.GetUint(fixConditions["subcategory_id"])
			}
			if _, ok := fixConditions["childcategory_id"]; ok {
				ChildcategoryID = util.GetUint(fixConditions["childcategory_id"])
			}

			if SubcategoryID > 0 && ChildcategoryID > 0 {
				query.Where("(subcategory_id = ? && childcategory_id = 0)  OR childcategory_id = ?", SubcategoryID, ChildcategoryID)
			} else if SubcategoryID > 0 {
				query.Where("subcategory_id = ? && childcategory_id = 0", SubcategoryID)
			} else if ChildcategoryID > 0 {
				query.Where("childcategory_id = ?", ChildcategoryID)
			}
		}

		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" attributes.title Like ? OR users.name Like ? OR attributes.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&attributes).Error
		}
		if err == nil {
			response.Data = attributes
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Attribute{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
