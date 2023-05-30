package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"encoding/json"
	"errors"
	"net/http"
)

var chathistoryAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	//var flag, added bool
	chathistory := make([]model.ChatHistory, 0)
	if _, ok := response.Request["ChatID"]; ok {
		ChatID := util.GetUint(response.Request["ChatID"])
		if ChatID > 0 {
			var chat model.Chat
			var item model.ChatHistory
			err = db.Conn.First(&chat, ChatID).Error
			if err == nil {
				if _, ID_ok := response.Request["ID"]; ID_ok {
					if _, Content_ok := response.Request["Content"]; Content_ok {

						id := util.GetID(response.Request)
						if id > 0 {
							err = db.Conn.First(&item, "id = ?", id).Error
						}
						if err != nil || id == 0 {
							item = model.ChatHistory{
								CreatedBy: response.User.ID,
								ChatID:    ChatID,
							}
							err = nil
						}
						_, err = item.Update(response.Request, response.User, &chat)
						if err == nil {
							response.Message = ""
							response.Status = apiresponse.SUCCESS
							response.Result["message"] = item

							sender := model.ChatListView{}
							sender.ID = response.User.ID
							sender.Name = response.User.Name
							sender.Photo = response.User.Photo
							sender.Online = response.User.Online
							sender.LastActiveOn = response.User.LastActiveOn
							sender.IsSupportagent = response.User.IsSupportagent

							response.Result["sender"] = sender
							var messageTo uint
							if chat.CreatedBy == response.User.ID {
								messageTo = chat.CreatedFor
							}
							if chat.CreatedFor == response.User.ID {
								messageTo = chat.CreatedBy
							}
							if messageTo > 0 {
								if _, online := apiresponse.OnlinUsers[messageTo]; online {
									message_string, err1 := json.Marshal(response.Result)
									if err1 == nil {
										apiresponse.OnlinUsers[messageTo].WS.WriteMessage(1, message_string)
									}
								}
							}
						} else if id == 0 {
							response.Message = "failed to add (" + err.Error() + ")"
						} else {
							response.Message = "failed to update (" + err.Error() + ")"
						}
					}
				} else if _, ID_ok := response.Request["delete_message"]; ID_ok {
					delete_message_id := util.GetUint(response.Request["delete_message"])
					if delete_message_id > 0 {
						err = db.Conn.First(&item, "id = ?", delete_message_id).Error
					}
					if err == nil {
						db.Conn.Delete(&item)
					}
					if err == nil {
						response.Result["message"] = item

						var messageTo uint
						if chat.CreatedBy == response.User.ID {
							messageTo = chat.CreatedFor
						}
						if chat.CreatedFor == response.User.ID {
							messageTo = chat.CreatedBy
						}
						if messageTo > 0 {
							if _, online := apiresponse.OnlinUsers[messageTo]; online {
								message_string, err1 := json.Marshal(response.Result)
								if err1 == nil {
									apiresponse.OnlinUsers[messageTo].WS.WriteMessage(1, message_string)
								}
							}
						}
					}
				}
			}
		}

		var query = db.Conn.Model(&model.ChatHistory{}).Where(" chat_id = ?", ChatID)
		query.Limit(50)
		query.Statement.Order(" created_at desc")

		if _, ok := response.Request["MessageID"]; ok {
			MessageID := util.GetUint(response.Request["MessageID"])
			query.Where(" id = ? ", MessageID)
		}

		if _, ok := response.Request["BeforID"]; ok {
			BeforID := util.GetUint(response.Request["BeforID"])
			query.Where(" id < ? ", BeforID)
		} else if _, ok := response.Request["AfterID"]; ok {
			AfterID := util.GetUint(response.Request["AfterID"])
			query.Where(" id > ? ", AfterID)
		}

		err = query.Unscoped().Scan(&chathistory).Error
		if err == nil {
			response.Result["chathistory"] = chathistory
			response.Status = apiresponse.SUCCESS
		} else {
			response.Result["chathistory"] = []model.ChatHistory{}
			response.Status = apiresponse.FAILED
		}
	} else if _, ok := response.Request["CreatedFor"]; ok {
		CreatedFor := util.GetUint(response.Request["CreatedFor"])
		if CreatedFor > 0 {
			chat := new(model.Chat)
			chat.CreatedBy = response.User.ID
			chat.UpdatedBy = response.User.ID
			chat.CreatedFor = CreatedFor

			var duplicates model.Chat
			err = db.Conn.First(&duplicates, "(`created_by` = ? AND `created_for` = ?) OR (`created_by` = ? AND `created_for` = ?)", chat.CreatedBy, chat.CreatedFor, chat.CreatedFor, chat.CreatedBy).Error
			if err == nil {
				err = errors.New("duplicate code")
			} else {
				err = db.Conn.Create(&chat).Error
			}
			if err == nil {
				response.Result["chat"] = chat
				response.Result["chathistory"] = []model.ChatHistory{}
				response.Status = apiresponse.SUCCESS
			}
		}
	}
	return err
}
