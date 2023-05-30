package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/awstorage"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var publishAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var from_theme_id uint
	var from_theme model.Theme
	var publishing_theme model.PublishedTheme
	if _, ok := response.Request["from_theme_id"]; ok {
		from_theme_id = util.GetUint(response.Request["from_theme_id"])
		if from_theme_id == 0 {
			return err
		}
		err = db.Conn.First(&from_theme, "id = ?", from_theme_id).Error
		if err != nil {
			return err
		}

		if from_theme.CreatedBy != response.User.ID {
			response.Message = "You can publish only your own theme"
			err = errors.New("invalid user")
			return err
		}

		err = db.Conn.First(&publishing_theme, " title = ?", from_theme.Title).Error
		if err == nil {
			if publishing_theme.ThemeID != from_theme.ID {
				err = errors.New("theme name allready used by some one , please try with different theme name")
				return err
			}
			if publishing_theme.Status == "Published" {
				response.Message = "Theme is allready published, you can unpublish it and then update it or publish it with new theme name"
				err = errors.New("theme is allready published")
				return err
			}
			if publishing_theme.Status == "Testing" {
				response.Message = "Theme is allready submited, please wait while completing testing"
				err = errors.New("theme is allready submited")
				return err
			}
		} else {
			publishing_theme = model.PublishedTheme{
				CreatedBy:            response.User.ID,
				ThemeID:              from_theme.ID,
				BasePublishedThemeID: from_theme.BasePublishedThemeID,
				AccountID:            from_theme.AccountID,
			}
		}

		if publishing_theme.ID > 0 && from_theme.UpdatedAt.Before(publishing_theme.CreatedAt) {
			err = errors.New("everything is allready published or in process of publishing")
			return err
		}

		publishing_theme.Title = from_theme.Title
		publishing_theme.Logo = from_theme.Logo
		publishing_theme.Images = from_theme.Images
		publishing_theme.Tags = from_theme.Tags
		publishing_theme.Description = from_theme.Description
		publishing_theme.UpdatedBy = response.User.ID
		publishing_theme.Status = "Submiting"
		publishing_theme.SubmitedOn = time.Now().Unix()

		err = db.Conn.Save(&publishing_theme).Error
		if err != nil {
			return err
		}
		from_theme.PublishedThemeID = publishing_theme.ID
		from_theme.SubmitedOn = publishing_theme.SubmitedOn
		from_theme.Status = publishing_theme.Status
		from_theme.UpdatedBy = response.User.ID
		err = db.Conn.Save(&from_theme).Error
		if err != nil {
			return err
		}
		publishing_theme_prefix := strings.TrimSpace("published_themes/theme_" + strconv.FormatUint(uint64(publishing_theme.ID), 10))
		prefix := strings.TrimSpace("themes/theme_" + strconv.FormatUint(uint64(from_theme.ID), 10))

		if from_theme.AccountID > 0 && !(from_theme.AccountID == response.Account.ID && response.Account.AccountType == "admin") {
			preprefix := "account_" + strconv.Itoa(int(from_theme.AccountID))
			if preprefix != "" {
				prefix = preprefix + "/" + prefix
			}
		}
		files, err := awstorage.ListBucketItems(prefix, "")
		if err == nil {
			if len(files) == 0 {
				err = errors.New("no files found")
			} else {
				for _, file := range files {
					var new_key = strings.Replace(*file.Key, prefix, publishing_theme_prefix, 1)
					err = awstorage.CopyFile(*file.Key, new_key)
					if err != nil {
						break
					}
				}
			}
		}

		if err == nil {
			publishing_theme.Status = "Submitted"
			err = db.Conn.Save(&publishing_theme).Error
		} else {
			publishing_theme.Status = "Fail to Upload"
			err = db.Conn.Save(&publishing_theme).Error
		}
		from_theme.Status = publishing_theme.Status
		_ = db.Conn.Save(&from_theme).Error
		if err == nil {
			response.Message = "Theme submited to review, it may take from few hours to upto the day to get published."
			response.Status = apiresponse.SUCCESS
		} else {
			response.Message = "Faled to submit theme to review for publishing"
		}
	} else {
		err = errors.New("empty or invalid request")
	}
	return err
}
