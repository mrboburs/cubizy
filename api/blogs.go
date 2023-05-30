package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

type blogview struct {
	model.Blog
	UpdatedByName string
	Photo         string
}

var blogsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if response.Account != nil && response.Account.ID > 0 {
		var blogs []blogview

		var query = db.Conn.Model(&model.Blog{}).Select("blogs.id, blogs.title, blogs.tags, blogs.category, blogs.image, blogs.updated_at, users.name AS updated_by_name, users.photo")
		query.Joins("left join users on blogs.updated_by = users.id")
		query.Where(" blogs.status = 1")

		if response.Account.AccountType != "admin" {
			query.Where("account_id = ?", response.Account.ID)
		}

		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["category"]; ok {
				category := util.GetUint(fixConditions["category"])
				query.Where("category = ?", category)
			}
		}

		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal

		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" blogs.title Like ? OR blogs.content Like ? OR users.name Like ? OR blogs.id Like ? ", SearchStringLike, SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}

		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&blogs).Error
		}
		if err == nil {
			response.Data = blogs
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []blogview{}
			response.Status = apiresponse.FAILED
		}
	} else {
		response.Message = "empty request"
	}
	return err
}
