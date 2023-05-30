package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/awstorage"
	"cubizy/util"
	"net/http"
	"strconv"
	"strings"
)

var filesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	response.Status = apiresponse.FAILED

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
	files, err := awstorage.ListBucketItems(prefix, "")
	if err == nil {
		response.Result["files"] = files
		response.Result["accessURL"] = awstorage.GetAccessURL()
		response.Status = apiresponse.SUCCESS
	} else {
		util.Log(err.Error())
	}
	return err
}
