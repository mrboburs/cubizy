package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/awstorage"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
	"strconv"
)

var themesettingsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if _, ok := response.Request["ThemeSettings"]; ok {
		ThemeSettings := util.GetString(response.Request["ThemeSettings"])
		response.Account.ThemeSettings = ThemeSettings
		err = db.Conn.Save(response.Account).Error
		if err == nil {
			response.Message = "Theme settings updated"
			response.Status = apiresponse.SUCCESS
		} else {
			response.Message = "Failed to update theme settings"
		}
	} else {
		var reset = false
		if _, ok := response.Request["ResetThemeSettings"]; ok {
			reset = util.GetBool(response.Request["ResetThemeSettings"])
		}
		if reset || response.Account.ThemeSettings == "" {
			var file_path = "themes/theme_" + strconv.Itoa(int(response.Account.ThemeID)) + "/settings.json"
			if response.Account.AccountType != "admin" {
				file_path = "account_" + strconv.Itoa(int(response.Account.ID)) + "/" + file_path
			}
			response.Result["ThemeSettings"], err = awstorage.ReadFile(file_path)
		} else {
			response.Result["ThemeSettings"] = response.Account.ThemeSettings
		}
		if err == nil {
			response.Status = apiresponse.SUCCESS
		}

	}
	return err
}
