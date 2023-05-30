package adminapi

import (
	"cubizy/apiresponse"
	"cubizy/keys"
	"cubizy/model"
	"cubizy/plugins/awstorage"
	"cubizy/plugins/db"
	"cubizy/plugins/sendgrid"
	"net/http"
	"strings"
)

var settingsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var setting model.Setting
	if _, ok := response.Request["setting"]; ok {
		postedSetting := response.Request["setting"].(map[string]interface{})
		id := uint(postedSetting["ID"].(float64))
		if id > 0 {
			err = db.Conn.First(&setting, "id = ?", id).Error
			if err == nil {
				setting.Value = strings.TrimSpace(postedSetting["Value"].(string))
				setting.UpdatedBy = response.User.ID
				db.Conn.Save(&setting)

				switch setting.Name {
				case keys.S3Bucket:
					awstorage.ResetS3()
				case keys.S3Region:
					awstorage.ResetS3()
				case keys.S3AwsAccessKeyID:
					awstorage.ResetS3()
				case keys.S3AwsSecretAccessKey:
					awstorage.ResetS3()
				case keys.SendgridAPIKey:
					sendgrid.Reset()
				case keys.SendgridSenderName:
					sendgrid.Reset()
				case keys.SendgridSenderEmail:
					sendgrid.Reset()
				}

				response.Message = " Setting Updated "
				response.Status = apiresponse.SUCCESS
			} else {
				return err
			}
		}
	} else {
		response.Result["settings"], err = model.GetAllSettings()
		response.Status = apiresponse.SUCCESS
	}
	return err
}
