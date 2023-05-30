package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var blogAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["blog_id"]; ok {
		blog_id := util.GetUint(response.Request["blog_id"])
		var blog model.Blog
		err = db.Conn.First(&blog, " id = ? ", blog_id).Error
		if err == nil {
			response.Result["blog"] = blog
			response.Status = apiresponse.SUCCESS
			response.Message = ""
		} else {
			response.Message = err.Error()
		}

		var author model.User
		err_author := db.Conn.Model(&model.User{}).Where(" id = ? ", blog.UpdatedBy).First(&author).Error
		if err == nil {
			response.Result["author"] = author
		} else {
			util.Log(err_author.Error())
		}

		var previous model.Blog
		err_previous := db.Conn.Model(&model.Blog{}).Select("blogs.id, blogs.title, blogs.image, blogs.updated_at").Where(" id < ? ", blog_id).Order("id desc").First(&previous).Error
		if err_previous == nil {
			response.Result["previous"] = previous
		} else {
			util.Log(err_previous.Error())
		}

		var next model.Blog
		err_next := db.Conn.Model(&model.Blog{}).Select("blogs.id, blogs.title, blogs.image, blogs.updated_at").Where(" id > ? ", blog_id).Order("id asc").First(&next).Error
		if err_next == nil {
			response.Result["next"] = next
		} else {
			util.Log(err_next.Error())
		}
	} else {
		response.Message = "Blog not found"
	}
	return err
}
