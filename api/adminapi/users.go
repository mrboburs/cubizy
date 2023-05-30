package adminapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
	"strconv"
)

type userview struct {
	model.User
	MobileCode    string
	EmailCode     string
	UpdatedByName string
}

var usersAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postUsers := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postUser := range postUsers {
			postUser := _postUser.(map[string]interface{})
			message := " Added "
			var item model.User
			id := util.GetID(postUser)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.User{}
				err = nil
			}
			var done bool
			done, err = item.Update(postUser, response.User)
			if err != nil {
				if message == " Added " {
					message = "failed to add (" + err.Error() + ")"
				} else {
					message = "failed to update (" + err.Error() + ")"
				}
			} else if !done {
				message = " unchanged "
			}
			count := messages[message]
			messages[message] = count + 1
		}

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + "Users "
			} else {
				response.Message += "User "
			}
			response.Message += key
		}
	} else if _, okDelete := response.Request["todelete"]; okDelete {

		todelete := response.Request["todelete"].([]interface{})
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			susccessMessage := ""
			if itemid > 0 {
				var user model.User
				err = db.Conn.Unscoped().First(&user, "id = ?", itemid).Error
				if err == nil {
					var account model.Account
					err1 := db.Conn.Unscoped().First(&account, "created_by = ?", user.ID).Error
					if err1 == nil {
						err = db.Conn.Exec("DELETE FROM addresses WHERE account_id = ? OR created_by = ?", account.ID, user.ID).Error
						if err != nil {
							util.Log("Deleting users address ", err)
						}
						err = db.Conn.Exec("DELETE FROM notes WHERE account_id = ?", account.ID).Error
						if err != nil {
							util.Log("Deleting users notes ", err)
						}
						err = db.Conn.Exec("DELETE FROM pages WHERE account_id = ?", account.ID).Error
						if err != nil {
							util.Log("Deleting users pages ", err)
						}
						err = db.Conn.Exec("DELETE FROM session_students WHERE account_id = ?", account.ID).Error
						if err != nil {
							util.Log("Deleting users session_students ", err)
						}
						err = db.Conn.Exec("DELETE FROM sessions WHERE account_id = ?", account.ID).Error
						if err != nil {
							util.Log("Deleting users sessions ", err)
						}
						if err == nil {
							err = db.Conn.Unscoped().Delete(&account).Error
						}
					}
					if user.IsStudent {
						err = db.Conn.Exec("DELETE FROM session_students WHERE student_id = ?", user.ID).Error
						if err != nil {
							util.Log("Deleting users from session_students ", err)
						}
					}
					if err == nil {
						err = db.Conn.Unscoped().Delete(&user).Error
					}
					if err == nil {
						susccessMessage = " User deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						errorMessage = "Failed to delete"
						break
					}
				} else {
					errorMessage = "Some Users not found"
				}
			} else {
				invalidIDMessage = " Some User ids are invalid "
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
	}

	if err == nil {
		var users []userview

		var query = db.Conn.Model(&model.User{}).Unscoped().Select("users.*, users.name AS updated_by_name").Joins("left join users as admin on users.updated_by = admin.id")

		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["user_type"]; ok {
				user_type := util.GetString(fixConditions["user_type"])
				switch user_type {
				case "admins":
					query.Where(" `users`.`is_super_admin` = true || `users`.`is_admin` = true ")
				case "sellers":
					query.Where(" `users`.`seller_account_id` > 0 ")
				}
			}
		}

		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" users.name Like ? OR users.email Like ? OR users.mobile Like ? OR users.id Like ? ", SearchStringLike, SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&users).Error
		}
		if err == nil {
			response.Data = users
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.User{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
