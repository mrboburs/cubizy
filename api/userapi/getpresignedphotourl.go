package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/awstorage"
	"net/http"
	"strconv"
)

var getpresignedphotourlAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var url string
	var key = "photo.jpg"
	fileType := "image/jpeg"
	preprefix := "user_" + strconv.Itoa(int(response.User.ID))
	key = preprefix + "/" + key
	url, err = awstorage.GetPresignedPutURL(key, fileType)
	if err == nil {
		response.Result["key"] = key
		response.Result["presignedUrl"] = url
		response.Result["accessURL"] = awstorage.GetAccessURL()
		response.Message = ""
		response.Status = apiresponse.SUCCESS
	}
	return err
}
