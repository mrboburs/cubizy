package adminapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/awstorage"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type themeview struct {
	model.PublishedTheme
	CreatedByName string
}

var publishedthemesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postThemes := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postTheme := range postThemes {
			postTheme := _postTheme.(map[string]interface{})
			message := " Added "
			var item model.PublishedTheme
			id := util.GetID(postTheme)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err == nil {
				_, err = item.Update(postTheme, response.User)
			}
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
				response.Message += strconv.Itoa(value) + "Themes "
			} else {
				response.Message += "Theme "
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
				var theme model.PublishedTheme
				err = db.Conn.First(&theme, "id = ?", itemid).Error
				if err == nil {
					prefix := strings.TrimSpace("published_themes/theme_" + strconv.FormatUint(uint64(theme.ID), 10))
					files, err := awstorage.ListBucketItems(prefix, "")
					if err == nil {
						for _, file := range files {
							err = awstorage.Delete(*file.Key)
							if err != nil {
								break
							}
						}
					}
					if err == nil {
						err = db.Conn.Unscoped().Delete(&theme).Error
					} else {
						errorMessage = "Failed to delete all files, please try again"
					}
					if err == nil {
						susccessMessage = " Theme deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Themes not found"
				}
			} else {
				invalidIDMessage = " Some Theme ids are invalid "
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
		var themes []themeview

		var query = db.Conn.Model(&model.PublishedTheme{}).Select("published_themes.*, users.name AS created_by_name").Joins("left join users on published_themes.updated_by = users.id")
		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["status"]; ok {
				status := util.GetString(fixConditions["status"])
				query.Where(" published_themes.status = ? ", status)
			}
		}
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" published_themes.title Like ? OR users.name Like ? OR published_themes.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&themes).Error
		}
		if err == nil {
			response.Data = themes
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.PublishedTheme{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
