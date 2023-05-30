package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var attributesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	util.Log("geting attributes")
	var attributes []model.Attribute
	SubcategoryID := util.GetUint(response.Request["SubcategoryID"])
	ChildcategoryID := util.GetUint(response.Request["ChildcategoryID"])
	attributes, err = model.GetAttributes(SubcategoryID, ChildcategoryID)
	if err != nil {
		util.Log(err)
		err = nil
		attributes = make([]model.Attribute, 0)
	}
	if _, ok := response.Request["ProductID"]; ok {
		ProductID := util.GetUint(response.Request["ProductID"])
		var results []map[string]interface{}
		db.Conn.Table("productattributes").Find(&results, " product_id = ?", ProductID)
		if len(results) > 0 {
			response.Result["ProductDetails"] = results[0]
		}
	}
	response.Data = attributes
	response.Status = apiresponse.SUCCESS
	return err
}
