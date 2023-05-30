package sellerapi

import (
	"cubizy/apiresponse"
	"cubizy/util"
	"net/http"
)

var dashbordAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	var SessionActive, StudentActive int64

	//var time = time.Now().Unix()
	//err = db.Conn.Model(&model.Session{}).Where("account_id = ? AND ends_on > ?", response.Account.ID, time).Count(&SessionActive).Error
	if err != nil {
		util.Log("getting SessionActive", err)
	}

	//err = db.Conn.Model(&model.SessionStudent{}).Distinct("student_id").Where("session_id IN (SELECT sessions.id FROM sessions WHERE account_id = ? AND ends_on > ?)", response.Account.ID, time).Count(&StudentActive).Error
	if err != nil {
		util.Log("getting StudentActive", err)
	}

	response.Result["SessionActive"] = SessionActive
	response.Result["StudentActive"] = StudentActive

	response.Message = ""
	response.Status = apiresponse.SUCCESS
	return err
}
