package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
	"strconv"
)

var cartitemsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		items := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _item := range items {
			postItem := _item.(map[string]interface{})
			message := " Added "
			item := model.Cartitem{
				AccountID: response.Account.ID,
			}
			_, err = item.Update(postItem, response.User)

			if err != nil {
				message = "failed to add (" + err.Error() + ")"
				util.Log("While adding cartitems item ")
				util.Log(err)
				err = nil
			}
			count := messages[message]
			messages[message] = count + 1
		}

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + " Cartitem items "
			} else {
				response.Message += "Cartitem items"
			}
			response.Message += key
		}
	}

	if err == nil {
		var cartitems []model.Cartitem
		var query = db.Conn.Model(&model.Cartitem{})

		if _, ok := response.Request["order_id"]; ok {
			var order_id = util.GetUint(response.Request["order_id"])
			query.Where("created_by = ? AND order_id = ?", response.User.ID, order_id)
		} else {
			query.Where("created_by = ? AND order_id = 0", response.User.ID)
		}

		err = query.Scan(&cartitems).Error
		if err == nil {
			response.Data = cartitems
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Cartitem{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
