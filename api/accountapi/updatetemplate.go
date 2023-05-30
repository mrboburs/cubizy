package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/awstorage"
	"cubizy/templates"
	"cubizy/util"
	"html/template"
	"net/http"
	"strconv"
)

var updatetemplateAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	var ThemeID uint = 0
	if response.Account.ThemeID > 0 {
		ThemeID = response.Account.ThemeID
		var index_file string
		util.Log("updating theme file ", response.Account.ThemeID)
		var path = "themes/theme_" + strconv.FormatUint(uint64(response.Account.ThemeID), 10) + "/index.html"
		if response.Account.AccountType != "admin" {
			path = "account_" + strconv.Itoa(int(response.Account.ID)) + "/" + path
		}
		index_file, err = awstorage.ReadFile(path)
		if err == nil && index_file != "" {
			util.Log("got theme file : ", index_file)
			_template := template.New(strconv.FormatUint(uint64(response.Account.ThemeID), 10))
			_template.Parse(index_file)
			templates.IndexTemplate[ThemeID] = _template
		} else if err != nil {
			util.Log(err)
		}
	}
	if err != nil {
		response.Message = "Failed to update Account template"
		response.Result["error"] = err.Error()
		response.Status = apiresponse.FAILED
	} else {
		response.Message = ""
		response.Status = apiresponse.SUCCESS
	}
	return err
}
