package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ChatHistory struct {
	gorm.Model
	ChatID    uint
	Content   string
	Type      int
	RecivedAt int64
	SeenAt    int64
	ReplayTo  uint
	UpdatedBy uint
	CreatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&ChatHistory{})
}

func (chathistory *ChatHistory) Update(ChatHistoryMap map[string]interface{}, UpdatedBy *User, chat *Chat) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	} else {
		chathistory.UpdatedBy = UpdatedBy.ID
	}

	if chathistory.CreatedBy == UpdatedBy.ID {

		if chathistory.ID == 0 {

			chathistory.ChatID = chat.ID

			if _, ok := ChatHistoryMap["ReplayTo"]; ok {
				Value := util.GetUint(ChatHistoryMap["ReplayTo"])
				if Value != chathistory.ReplayTo {
					chathistory.ReplayTo = Value
					flag = true
				}
			}
		}

		if _, ok := ChatHistoryMap["Content"]; ok {
			Value := util.GetString(ChatHistoryMap["Content"])
			if Value != chathistory.Content {
				chathistory.Content = Value
				flag = true
			}
		}
		if _, ok := ChatHistoryMap["Type"]; ok {
			Value := util.GetInt(ChatHistoryMap["Type"])
			if Value != chathistory.Type {
				chathistory.Type = Value
				flag = true
			}
		}
	} else {
		if chathistory.RecivedAt == 0 {
			chathistory.RecivedAt = time.Now().Unix()
			flag = true
		}
		if _, ok := ChatHistoryMap["Seen"]; ok {
			chathistory.SeenAt = time.Now().Unix()
			flag = true
		}
	}

	if err == nil && flag {
		if chathistory.Content == "" {
			err = errors.New(" Message can not be empty")
		} else if chathistory.ChatID == 0 {
			err = errors.New(" ChatID can not be empty")
		} else {
			var flagAdded bool
			if chathistory.ID == 0 {
				flagAdded = true
				err = db.Conn.Create(&chathistory).Error
			} else {
				flagAdded = false
				err = db.Conn.Save(&chathistory).Error
			}
			if err == nil {
				// find last message and update it

				if flagAdded {
					chat.Content = chathistory.Content
					chat.Type = chathistory.Type
					chat.ReplayTo = chathistory.ReplayTo
					chat.RecivedAt = chathistory.RecivedAt
					chat.SeenAt = chathistory.SeenAt
					chat.TotalMessages = chat.TotalMessages + 1
					chat.MessageID = chathistory.ID
					err = db.Conn.Save(&chat).Error
				}
			}
		}
	}
	return flag, err
}
