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

type locationview struct {
	model.Location
	UpdatedByName string
}

var locationsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postLocations := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postLocation := range postLocations {
			postLocation := _postLocation.(map[string]interface{})
			message := " Added "
			var item model.Location
			id := util.GetID(postLocation)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Location{}
				err = nil
			}
			_, err = item.Update(postLocation, response.User)
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
				response.Message += strconv.Itoa(value) + "Locations "
			} else {
				response.Message += "Location "
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
				var location model.Location
				err = db.Conn.First(&location, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&location).Error
					if err == nil {
						susccessMessage = " Location deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Locations not found"
				}
			} else {
				invalidIDMessage = " Some Location ids are invalid "
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
		var locations []locationview

		var query = db.Conn.Model(&model.Location{}).Select("locations.*, users.name AS updated_by_name").Joins("left join users on locations.updated_by = users.id")
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" locations.code Like ? OR locations.sub_locality Like ? OR locations.locality Like ? OR locations.country Like ? ", SearchStringLike, SearchStringLike, SearchStringLike, SearchStringLike)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&locations).Error
		}
		if err == nil {
			response.Data = locations
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Location{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
