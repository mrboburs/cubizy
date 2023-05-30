package adminapi

import (
	"cubizy/apiresponse"
	"net/http"
)

var dashbordAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	var SMSBalance, SMSSent, EmailBalance, EmailSent int64

	//db.Model(&User{}).Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Count(&count)
	//var time = time.Now().Unix()

	response.Result["SMSBalance"] = SMSBalance
	response.Result["SMSSent"] = SMSSent

	response.Result["EmailBalance"] = EmailBalance
	response.Result["EmailSent"] = EmailSent

	response.Message = ""
	response.Status = apiresponse.SUCCESS
	return err
}
