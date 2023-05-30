package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
)

var blogcategoriesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var blogcategories []model.Blogcategory
	err = db.Conn.Model(&model.Blogcategory{}).Where("status = ?", true).Scan(&blogcategories).Error
	if err == nil {
		response.Data = blogcategories
		response.Status = apiresponse.SUCCESS
	}
	return err
}
