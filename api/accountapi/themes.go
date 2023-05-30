package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/awstorage"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
)

type themeview struct {
	model.Theme
	UpdatedByName string
}

var themesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["GetPublishedThemeID"]; ok {
		var PublishedThemeID = util.GetUint(response.Request["GetPublishedThemeID"])
		var item model.Theme
		if PublishedThemeID > 0 {
			var publishedTheme model.PublishedTheme
			err = db.Conn.First(&publishedTheme, "id = ?", PublishedThemeID).Error
			if err == nil {
				itemerr := db.Conn.First(&item, "title = ? AND account_id = ?", publishedTheme.Title, response.Account.ID).Error

				if itemerr != nil {
					item = model.Theme{
						Title:            publishedTheme.Title,
						PublishedThemeID: publishedTheme.ID,
						Logo:             publishedTheme.Logo,
						Tags:             publishedTheme.Tags,
						Description:      publishedTheme.Description,
						Images:           publishedTheme.Images,
						AccountID:        response.Account.ID,
						CreatedBy:        response.User.ID,
						UpdatedBy:        response.User.ID,
					}
					err = db.Conn.Create(&item).Error
				} else {
					item.Title = publishedTheme.Title
					item.PublishedThemeID = publishedTheme.ID
					item.BasePublishedThemeID = publishedTheme.BasePublishedThemeID
					item.Logo = publishedTheme.Logo
					item.Tags = publishedTheme.Tags
					item.Description = publishedTheme.Description
					item.Images = publishedTheme.Images
					item.AccountID = response.Account.ID
					item.CreatedBy = response.User.ID
					item.UpdatedBy = response.User.ID
					item.IsUpdated = false
					err = db.Conn.Save(&item).Error
				}

				if err == nil {
					new_prefix := strings.TrimSpace("themes/theme_" + strconv.FormatUint(uint64(item.ID), 10))
					prefix := strings.TrimSpace("published_themes/theme_" + strconv.FormatUint(uint64(publishedTheme.ID), 10))

					if item.AccountID > 0 && response.Account.AccountType != "admin" {
						preprefix := "account_" + strconv.Itoa(int(item.AccountID))
						if preprefix != "" {
							new_prefix = preprefix + "/" + new_prefix
						}
					}

					//util.Log("Copying files of theme")
					//util.Log("prefix : ", prefix)
					//util.Log("new_prefix : ", new_prefix)
					var files []*s3.Object
					files, err = awstorage.ListBucketItems(prefix, "")
					if err == nil {
						if len(files) == 0 {
							err = errors.New("no files found")
						} else {
							for _, file := range files {
								var new_key = strings.Replace(*file.Key, prefix, new_prefix, 1)
								err = awstorage.CopyFile(*file.Key, new_key)
								if err != nil {
									break
								}
							}
						}
					}
					if err == nil {
						item.IsUpdated = true
						err = db.Conn.Save(&item).Error
					}
				}
			} else {
				err = errors.New("base theme not found")
			}
		}
	} else if _, ok := response.Request["CopyFromThemeID"]; ok {
		var FromThemeID = util.GetUint(response.Request["CopyFromThemeID"])
		var item model.Theme
		if FromThemeID > 0 {
			var FromTheme model.Theme
			err = db.Conn.First(&FromTheme, "id = ? AND account_id = ?", FromThemeID, response.Account.ID).Error
			if err == nil {
				var NewTitle = "Copy of " + FromTheme.Title
				itemerr := db.Conn.First(&item, "title = ? AND account_id = ?", NewTitle, response.Account.ID).Error

				BasePublishedThemeID := FromTheme.PublishedThemeID
				if BasePublishedThemeID == 0 {
					BasePublishedThemeID = FromTheme.BasePublishedThemeID
				}
				if itemerr != nil {
					item = model.Theme{
						Title:                NewTitle,
						ThemeID:              FromTheme.ID,
						BasePublishedThemeID: BasePublishedThemeID,
						Logo:                 FromTheme.Logo,
						Tags:                 FromTheme.Tags,
						Description:          FromTheme.Description,
						Images:               FromTheme.Images,
						AccountID:            response.Account.ID,
						CreatedBy:            response.User.ID,
						UpdatedBy:            response.User.ID,
					}
					err = db.Conn.Create(&item).Error
				} else {
					response.Message = "You allready have theme of name " + NewTitle + ", change name of that theme and try again."
					err = errors.New("you allready have theme of name")
					return err
				}

				if err == nil {
					new_prefix := strings.TrimSpace("themes/theme_" + strconv.FormatUint(uint64(item.ID), 10))
					prefix := strings.TrimSpace("themes/theme_" + strconv.FormatUint(uint64(FromTheme.ID), 10))

					if item.AccountID > 0 && response.Account.AccountType != "admin" {
						preprefix := "account_" + strconv.Itoa(int(item.AccountID))
						if preprefix != "" {
							new_prefix = preprefix + "/" + new_prefix
							prefix = preprefix + "/" + prefix
						}
					}

					//util.Log("Copying files of theme")
					//util.Log("prefix : ", prefix)
					//util.Log("new_prefix : ", new_prefix)
					var files []*s3.Object
					files, err = awstorage.ListBucketItems(prefix, "")
					if err == nil {
						if len(files) == 0 {
							err = errors.New("no files found")
						} else {
							for _, file := range files {
								var new_key = strings.Replace(*file.Key, prefix, new_prefix, 1)
								err = awstorage.CopyFile(*file.Key, new_key)
								if err != nil {
									break
								}
							}
						}
					}
					if err == nil {
						item.IsUpdated = true
						err = db.Conn.Save(&item).Error
					}
				}
			} else {
				err = errors.New("base theme not found")
			}
		}
	} else if _, ok := response.Request["items"]; ok {
		postThemes := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postTheme := range postThemes {
			postTheme := _postTheme.(map[string]interface{})
			message := " Added "
			var item model.Theme
			id := util.GetID(postTheme)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				if response.User.IsAdmin {
					item = model.Theme{
						AccountID: response.Account.ID,
						CreatedBy: response.User.ID,
					}
					err = nil
				} else {
					response.Message = "theme not found, you can create new theme from any or other theme"
					err = errors.New("only admin can creat new theme directly, others can creat theme from any of allready created theme")
					return err
				}
			}
			_, err = item.Update(postTheme, response.User)
			if err != nil {
				if message == " Added " {
					message = "failed to add (" + err.Error() + ")"
				} else {
					message = "failed to update (" + err.Error() + ")"
				}
			}
			count := messages[message]
			messages[message] = count + 1
		}

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + "Themes "
			} else {
				response.Message += "Theme "
			}
			response.Message += key
		}
		if err == nil {
			response.Status = apiresponse.SUCCESS
		}
	} else if _, okDelete := response.Request["todelete"]; okDelete {

		todelete := response.Request["todelete"].([]interface{})
		susccessMessage := ""
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			if itemid > 0 {
				var theme model.Theme
				err = db.Conn.First(&theme, "id = ? AND account_id = ? ", itemid, response.Account.ID).Error
				if err == nil {
					prefix := strings.TrimSpace("themes/theme_" + strconv.FormatUint(uint64(theme.ID), 10))
					if response.Account != nil && response.Account.AccountType != "admin" && response.Account.ID > 0 {
						preprefix := "account_" + strconv.Itoa(int(response.Account.ID))
						if preprefix != "" {
							prefix = preprefix + "/" + prefix
						}
					}
					files, err := awstorage.ListBucketItems(prefix, "")
					if err == nil {
						for _, file := range files {
							err = awstorage.Delete(*file.Key)
							if err != nil {
								break
							}
						}
					}
					if err == nil {
						err = db.Conn.Unscoped().Delete(&theme).Error
					} else {
						errorMessage = "Failed to delete all files, please try again"
					}
					if err == nil {
						susccessMessage = " Theme deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Themes not found"
				}
			} else {
				invalidIDMessage = " Some Theme ids are invalid "
			}
			response.Message = ""
			if susccessMessage != "" {
				response.Message += susccessMessage
			}
			if errorMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += errorMessage
			}
			if invalidIDMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += invalidIDMessage
			}
		}
		if susccessMessage == "" {
			err = errors.New("faild to delete any thing")
		} else {
			util.Log(susccessMessage)
		}
	}

	if err == nil {
		var themes []themeview

		var query = db.Conn.Model(&model.Theme{}).Select("themes.*, users.name AS updated_by_name").Joins("left join users on themes.updated_by = users.id")
		query.Where(" account_id = ? ", response.Account.ID)
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" themes.title Like ? OR users.name Like ? OR themes.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&themes).Error
		}
		if err == nil {
			response.Data = themes
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Theme{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
