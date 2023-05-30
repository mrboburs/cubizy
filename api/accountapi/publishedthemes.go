package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

type publishedthemesview struct {
	model.PublishedTheme
	CreatedByName string
}

var publishedthemesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if err == nil {
		var themes []publishedthemesview

		var query = db.Conn.Model(&model.PublishedTheme{})
		query.Select("published_themes.*, users.name AS created_by_name")
		query.Joins("left join users on published_themes.updated_by = users.id")

		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["status"]; ok {
				status := util.GetString(fixConditions["status"])
				query.Where(" published_themes.status = ? ", status)
			} else {
				query.Where(" published_themes.status = 'Published' ")
			}
		} else {
			query.Where(" published_themes.status = 'Published' ")
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
