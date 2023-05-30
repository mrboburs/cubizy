package userapi

import (
	"cubizy/apiresponse"
	"net/http"
)

var setquestionanswersAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	_, err = response.User.Update(response.Request, response.User)
	if err == nil {
		response.Message = "Two-factor authentication is set"
		response.Status = apiresponse.SUCCESS
	} else {
		response.Message = "Failed to update Secret Questions"
	}
	return err
}
