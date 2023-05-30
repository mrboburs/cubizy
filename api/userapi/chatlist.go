package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
)

var chatlistAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var chatlist []model.ChatListView
	var chat_columns = "chats.ispublic, chats.created_at, chats.updated_at, chats.deleted_at,  chats.content, chats.`type`, chats.recived_at, chats.seen_at, chats.replay_to, chats.updated_by, chats.created_by, chats.message_id, chats.total_messages, chats.total_users, chats.id AS chat_id, chats.created_for"
	var blank_chat_columns = "false AS ispublic, null AS created_at, null AS updated_at, null AS deleted_at, '' AS content, 0 AS `type`, null AS recived_at, null AS seen_at, 0 AS replay_to, 0 AS updated_by, 0 AS created_by, 0 AS message_id, 0 AS total_messages, 0 AS total_users, 0 AS chat_id"
	var query = db.Conn.Raw(`
	 	Select `+chat_columns+`, users.id, users.name, users.photo, users.online, users.last_active_on, users.is_supportagent
		FROM chats
		Left Join users ON chats.created_by = users.id
		WHERE chats.created_for = ?
		UNION
		Select `+chat_columns+`, users.id, users.name, users.photo, users.online, users.last_active_on, users.is_supportagent
		FROM chats
		Left Join users ON chats.created_for = users.id
		WHERE chats.created_by = ? AND chats.created_for != 0
		UNION		
		Select chats.*, chats.id AS chat_id, chats.title AS name, chats.image AS photo, false AS is_supportagent
		FROM chat_users
		Left Join chats ON chat_users.chat_id = chats.id
		WHERE chat_users.user_id = ?
		UNION
		Select chats.*, chats.id AS chat_id, chats.title AS name, chats.image AS photo, false AS is_supportagent
		FROM chats
		WHERE (chats.created_by = ? OR chats.ispublic ) AND chats.created_for = 0
		UNION
		Select `+blank_chat_columns+`, users.id AS created_for , users.id, users.name, users.photo, users.online, users.last_active_on, users.is_supportagent
		FROM users
		WHERE users.is_supportagent = true`, response.User.ID, response.User.ID, response.User.ID, response.User.ID)

	/*
		    Content       string
			Type          int
			RecivedAt     time.Time
			SeenAt        time.Time
			ReplayTo      uint
			UpdatedBy     uint
			CreatedBy     uint
			CreatedFor    uint
			MessageID     uint
			TotalMessages uint
			TotalUsers    uint





	*/

	err = query.Scan(&chatlist).Error
	if err == nil {
		response.Data = chatlist
		response.Status = apiresponse.SUCCESS
	} else {
		response.Data = []model.User{}
		response.Status = apiresponse.FAILED
	}
	return err
}
