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

type blogcategoryview struct {
	model.Blogcategory
	UpdatedByName string
}

var blogcategoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postBlogcategories := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postBlogcategory := range postBlogcategories {
			postBlogcategory := _postBlogcategory.(map[string]interface{})
			message := " Added "
			var item model.Blogcategory
			id := util.GetID(postBlogcategory)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Blogcategory{
					CreatedBy: response.User.ID,
				}
				err = nil
			}
			_, err = item.Update(postBlogcategory, response.User)
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
				response.Message += strconv.Itoa(value) + "Blogcategories "
			} else {
				response.Message += "Blogcategory "
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
				var blogcategory model.Blogcategory
				err = db.Conn.First(&blogcategory, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&blogcategory).Error
					if err == nil {
						susccessMessage = " Blogcategory deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Blogcategories not found"
				}
			} else {
				invalidIDMessage = " Some Blogcategory ids are invalid "
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
		var blogcategories []blogcategoryview

		var query = db.Conn.Model(&model.Blogcategory{}).Select("blogcategories.*, users.name AS updated_by_name").Joins("left join users on blogcategories.updated_by = users.id")
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" blogcategories.title Like ? OR users.name Like ? OR blogcategories.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&blogcategories).Error
		}
		if err == nil {
			response.Data = blogcategories
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Blogcategory{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
