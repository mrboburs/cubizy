package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var pageAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["page_id"]; ok {
		page_id := util.GetString(response.Request["page_id"])
		var page model.Page
		err = db.Conn.First(&page, " id = ? ", page_id).Error
		if err == nil {
			response.Result["content"] = page.Content
			response.Result["title"] = page.Title
			response.Status = apiresponse.SUCCESS
			response.Message = ""
		} else {
			response.Message = err.Error()
		}
	} else {
		response.Message = "Page not found"
	}
	return err
}
