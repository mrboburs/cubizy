package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var registerAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var user model.User
	ok, err := user.Update(response.Request, &user)
	if err != nil {
		response.Message = "Failed : " + err.Error()
	} else if !ok {
		response.Message = "Account not created"
	} else {
		response.Message = "Account created successfully Welcome " + user.Name
		user.Token = uuid.New().String()
		user.LastToken = user.Token
		user.LastLoginOn = user.LoginOn
		user.LoginOn = time.Now().Unix()
		user.LastActiveOn = time.Now().Unix()
		err = db.Conn.Save(&user).Error
		if err == nil {
			response.User = &user
			response.Account = model.GetAccount(&user, response.Domain)
			response.Status = apiresponse.SUCCESS
		}
		response.Status = apiresponse.SUCCESS
	}
	return err
}
