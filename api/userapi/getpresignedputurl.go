package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/awstorage"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var getpresignedputurlAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var url string

	if _, ok := response.Request["key"]; !ok || strings.TrimSpace(response.Request["key"].(string)) == "" {
		err = errors.New("invalid request")
		return err
	}
	key := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(response.Request["key"].(string)), " ", ""), "(", ""), ")", "")
	fileType := strings.TrimSpace(response.Request["type"].(string))

	prefix := strings.TrimSpace(response.Request["prefix"].(string))
	if !response.User.IsAdmin {
		preprefix := "user_" + strconv.Itoa(int(response.User.ID))

		if response.Account != nil && response.Account.ID > 0 && response.User.SellerAccountID == response.Account.ID {
			preprefix = "account_" + strconv.Itoa(int(response.Account.ID))
		}
		if preprefix != "" {
			prefix = preprefix + "/" + prefix
		}
	}

	if prefix != "" {
		key = prefix + "/" + key
	}
	url, err = awstorage.GetPresignedPutURL(key, fileType)
	if err == nil {
		response.Result["key"] = key
		response.Result["presignedUrl"] = url
		response.Message = ""
		response.Status = apiresponse.SUCCESS
	}
	return err
}
