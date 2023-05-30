package userapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var setpasswordAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if response.User.Password == "" {
		if _, ok := response.Request["Password"]; ok {
			password := util.GetString(response.Request["Password"])
			var passwordHash []byte
			passwordHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err == nil {
				password = string(passwordHash)
				response.User.Password = password
				db.Conn.Save(response.User)
				response.Message = " New password set successfully. "
				response.Status = apiresponse.SUCCESS
			} else {
				response.Message = " Failed to set password.  " + err.Error()
			}
		} else {
			response.Message = " Empty request. "
		}
	} else {
		response.Message = " Invalid request. "
	}
	return err
}
