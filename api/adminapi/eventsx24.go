package adminapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
	"strconv"
)

var eventsx24API = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var data []map[string]interface{}

	if _, OK := response.Request["select_eventsx24"]; OK {
		selecteventsx24 := response.Request["select_eventsx24"].(string)
		// db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)
		util.Log(selecteventsx24)
		err = db.Conn.Raw(selecteventsx24).Scan(&data).Error
		if err == nil {
			response.Data = data
			response.Status = apiresponse.SUCCESS
		}
	} else if _, OK := response.Request["execute_eventsx24"]; OK {
		selecteventsx24 := response.Request["execute_eventsx24"].(string)
		// db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)
		util.Log(selecteventsx24)
		var trx = db.Conn.Exec(selecteventsx24)
		err = trx.Error
		if err == nil {
			response.Message = strconv.FormatInt(trx.RowsAffected, 10) + " Rows Affected"
			response.Status = apiresponse.SUCCESS
		}
	}
	return err
}
