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

type blogview struct {
	model.Blog
	CategoryName  string
	UpdatedByName string
}

var blogsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postBlogs := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postBlog := range postBlogs {
			postBlog := _postBlog.(map[string]interface{})
			message := " Added "
			var item model.Blog
			id := util.GetID(postBlog)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Blog{
					AccountID: response.Account.ID,
				}
				err = nil
			}
			_, err = item.Update(postBlog, response.User)
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
				response.Message += strconv.Itoa(value) + "Blogs "
			} else {
				response.Message += "Blog "
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
				var blog model.Blog
				err = db.Conn.First(&blog, "id = ?", itemid).Error
				if err == nil {
					err = db.Conn.Delete(&blog).Error
					if err == nil {
						susccessMessage = " Blog deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Blogs not found"
				}
			} else {
				invalidIDMessage = " Some Blog ids are invalid "
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
		var blogs []blogview

		var query = db.Conn.Model(&model.Blog{}).Select("blogs.*, users.name AS updated_by_name, blogcategories.name AS category_name").Joins("left join users on blogs.updated_by = users.id left join blogcategories on blogs.category = blogcategories.id")
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" blogs.title Like ? OR users.name Like ? OR blogs.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
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
			response.Data = []model.Blog{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
