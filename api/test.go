package api

import (
	"cubizy/apiresponse"

	"net/http"
)

var testAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	response.Result["content"] = "testapi called successfully"
	response.Message = ""
	response.Status = apiresponse.SUCCESS
	return err
}
