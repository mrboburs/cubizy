package apiresponse

import (
	"cubizy/model"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// NOTFOUND is status when requisted API not found
	NOTFOUND = 0
	// FAILED is statuse for failed ApiRequest
	FAILED = 1
	// SUCCESS is statuse for successful ApiRequest
	SUCCESS = 2
	// NOTCOMPLETED is statuse when not all ApiRequest are successful
	NOTCOMPLETED = 3
	// INPROGRESS is statuse for ApiRequest which could not comleted in time but will keep runing afer responce
	INPROGRESS = 4
	// ONHOLD is statuse for ApiRequest kept on holde and will be continued after some rules are satisfied
	ONHOLD = 5
	// UNOTHARISED is statuse for unothariased ApiRequest
	UNOTHARISED = 6
	// PRIMARY is status for PRIMARY messages like welcome messges, or new Offer messages
	PRIMARY = 7

	// UserTokenCookieName is a cookie name for token of login user
	UserTokenCookieName = "UserToken"
	// UserEmailCookieName is a cookie name for login user email
	UserEmailCookieName = "UserEmail"
	// UserNameCookieName is a cookie name for login user name
	UserNameCookieName = "UserName"
	// UserVarifiedCookieName is a cookie name for login user verification status
	UserVarifiedCookieName = "UserVarified"
	// UserVerificationLinkSentCookieName is a cookie name for login user verification link sent status
	UserVerificationLinkSentCookieName = "UserVerificationLinkSent"
)

// Response is object will provied required data to api calls and return api responces
type Response struct {
	Status          int
	Message         string
	Domain          string
	API             string
	User            *model.User
	Account         *model.Account
	Result          map[string]interface{}
	Data            interface{} `json:"data" `
	Request         map[string]interface{}
	URLparts        string
	StartTime       time.Time
	EndTime         time.Time
	ResponseTime    time.Duration
	RecordsTotal    int64 `json:"recordsTotal" `
	RecordsFiltered int64 `json:"recordsFiltered" `
}

type OnlinUser struct {
	User *model.User
	WS   *websocket.Conn
}

// Handler that calls function with Response.
type Handler func(http.ResponseWriter, *http.Request, *Response) error

var OnlinUsers map[uint]OnlinUser

func init() {
	OnlinUsers = make(map[uint]OnlinUser)
}

/*
// SetUserCookies will set Auth releted cookies for user of current request
func SetUserCookies(w http.ResponseWriter, currentUser *model.User) {
	if currentUser.Email != "" && currentUser.AuthToken != "" {
		SetCookie(w, UserTokenCookieName, currentUser.AuthToken, currentUser.AutoTimeOut, true)
		SetCookie(w, UserEmailCookieName, currentUser.Email, currentUser.AutoTimeOut, false)
		SetCookie(w, UserNameCookieName, currentUser.Name, currentUser.AutoTimeOut, false)
	} else {
		if currentUser.Email == "" {
			util.Log("User email is empty for given cookie auth token")
		} else if currentUser.AuthToken == "" {
			util.Log("Auth code is empty for given cookie auth token")
		}
		util.Log("Deleting Cookies")
		DropUserCookies(w)
	}
}

// DropUserCookies will drop Auth releted cookies for user of current request in case of loged out
func DropUserCookies(w http.ResponseWriter) {
	DeleteCookie(w, UserTokenCookieName)
	DeleteCookie(w, UserEmailCookieName)
	DeleteCookie(w, UserVarifiedCookieName)
}

// SetCookie will set cookie for you in this application
func SetCookie(w http.ResponseWriter, name, value string, timeOut time.Time, secure bool) {
	var cookie *http.Cookie
	if config.Server.Domain == "localhost" {
		cookie = &http.Cookie{
			Name:     name,  //api.UserToken,
			Value:    value, // strconv.FormatInt(token, 10),
			HttpOnly: secure,
			Path:     "/",
			Expires:  timeOut, //user.TimeOut,
		}
	} else {
		cookie = &http.Cookie{
			Name:     name,  //api.UserToken,
			Value:    value, // strconv.FormatInt(token, 10),
			HttpOnly: secure,
			Path:     "/",
			Domain:   "." + config.Server.Domain,
			Expires:  timeOut, //user.TimeOut,
		}
	}
	//util.Log("Setting cookie", name, " to value ", value)
	http.SetCookie(w, cookie)
}

// DeleteCookie will delete cookie for you
func DeleteCookie(w http.ResponseWriter, name string) {
	var cookie *http.Cookie
	if config.Server.Domain == "localhost" {
		cookie = &http.Cookie{
			Name:    name,
			Value:   "",
			Path:    "/",
			Expires: time.Unix(0, 0),
		}
	} else {
		cookie = &http.Cookie{
			Name:    name,
			Value:   "",
			Path:    "/",
			Domain:  "." + config.Server.Domain,
			Expires: time.Unix(0, 0),
		}
	}
	//util.Log("Droping cookie", name)
	http.SetCookie(w, cookie)
}
*/
