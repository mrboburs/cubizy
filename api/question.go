package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
	"strconv"
)

var questionsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if response.Account != nil && response.Account.AccountType == "admin" {
		if _, ok := response.Request["items"]; ok {
			postQuestions := response.Request["items"].([]interface{})
			messages := make(map[string]int)
			for _, _postQuestion := range postQuestions {
				postQuestion := _postQuestion.(map[string]interface{})
				message := " Added "
				var item model.Question
				id := util.GetID(postQuestion)
				if id > 0 {
					err = db.Conn.First(&item, "id = ?", id).Error
					if err == nil {
						message = " Updated "
					}
				}
				if err != nil || id == 0 {
					item = model.Question{}
					err = nil
				}
				_, err = item.Update(postQuestion, response.User)
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
					response.Message += strconv.Itoa(value) + " Questions "
				} else {
					response.Message += "Question "
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
					var question model.Question
					err = db.Conn.First(&question, "id = ?", itemid).Error
					if err == nil {
						err = db.Conn.Delete(&question).Error
						if err == nil {
							susccessMessage = " Question deleted "
							response.Status = apiresponse.SUCCESS
						} else {
							break
						}
					} else {
						errorMessage = "Some Questions not found"
					}
				} else {
					invalidIDMessage = " Some Question ids are invalid "
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
	}

	if err == nil {
		var questions []model.Question

		var query = db.Conn.Model(&model.Question{})
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal

		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" questions.question Like ? ", SearchStringLike)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}

		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&questions).Error
		}
		if err == nil {
			response.Data = questions
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Question{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
