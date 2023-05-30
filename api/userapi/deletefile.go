package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/awstorage"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var deletefileAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["key"]; !ok || strings.TrimSpace(response.Request["key"].(string)) == "" {
		err = errors.New("invalid request")
		return err
	}
	key := strings.TrimSpace(response.Request["key"].(string))

	if strings.Contains(key, "user_"+strconv.FormatUint(uint64(response.User.ID), 10)+"/") || strings.Contains(key, "account_"+strconv.FormatUint(uint64(response.User.SellerAccountID), 10)+"/") || response.User.IsAdmin || response.User.IsSuperAdmin {
		err = awstorage.Delete(key)
		if err == nil {
			response.Message = "File Deleted"
			response.Status = apiresponse.SUCCESS
		}
	} else {
		response.Message = "Only file Owner can delete file"
	}
	return err
}
