package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/awstorage"
	"cubizy/util"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
)

var exportThemeAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	response.Status = apiresponse.FAILED

	if _, ok := response.Request["prefix"]; ok {
		var prefix = util.GetString(response.Request["prefix"])
		if prefix != "" {
			if response.Account.AccountType != "admin" {
				preprefix := "account_" + strconv.Itoa(int(response.Account.ID))
				if preprefix != "" {
					prefix = preprefix + "/" + prefix
				}
			}

			//util.Log("Copying files of theme")
			//util.Log("prefix : ", prefix)
			//util.Log("new_prefix : ", new_prefix)
			var files []*s3.Object
			var path string
			files, err = awstorage.ListBucketItems(prefix, "")
			if err == nil {
				var file_paths []string

				for _, file := range files {
					file_paths = append(file_paths, *file.Key)
				}
				folder := "static/default/files"
				filename := "zipedthemefiles" + strconv.FormatUint(uint64(response.Account.ID), 10)
				path, err = awstorage.ZipFiles(file_paths, folder, filename)
			}
			if err == nil {
				response.Result["zipfile"] = strings.ReplaceAll(path, "static/default", "")
				go func() {
					time.Sleep(60 * time.Second)
					util.Log("removing file")
					e := os.Remove(path)
					if e != nil {
						util.Log("file not removed")
						util.Log(e)
					} else {
						util.Log("file removed")
					}
				}()
				response.Message = "Exported theme in zip file "
				response.Status = apiresponse.SUCCESS
			}
		}
	}
	return err
}
