package myws

import (
	"cubizy/api"
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	userid := util.GetUint(r.URL.Path[len("/ws/"):])
	if _, ok := apiresponse.OnlinUsers[userid]; ok {
		if apiresponse.OnlinUsers[userid].WS != nil {
			w.WriteHeader(http.StatusAlreadyReported)
			util.Log("user is conected from anather tab or before this request")
			return
		}
	}
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		util.Log(err)
	}
	reader(ws, w, r)
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn, w http.ResponseWriter, r *http.Request) {

	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = r.Host
	}

	if origin != r.Host {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	// http
	origin = strings.Replace(origin, "http://", "", 1)
	origin = strings.Replace(origin, "https://", "", 1)

	SubDomain := strings.ReplaceAll(origin, util.Settings.Domain, "")
	SubDomain = strings.Trim(SubDomain, ".")
	SubDomain = strings.TrimSpace(SubDomain)

	//if SubDomain != "admin" && SubDomain != "seller" {
	//util.Log("host     : " + r.Host)
	//util.Log("origin   : " + r.Header.Get("Origin"))
	//util.Log("SubDomain: " + SubDomain)
	//}

	response := apiresponse.Response{
		StartTime: time.Now(),
		Message:   "",
		Status:    apiresponse.FAILED,
		Result:    make(map[string]interface{}),
		Domain:    SubDomain,
	}
	var ask = "who are you?"
	conn.WriteMessage(0, []byte(ask))
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			util.Log(err)
			break
		}
		// print out that message for clarity
		//util.Log(string(p))
		//util.Log("messageType", messageType)

		json.Unmarshal(p, &response.Request)
		if response.Request == nil {
			response.Request = make(map[string]interface{})
		}
		response.Result = make(map[string]interface{})
		response.API = util.GetString(response.Request["api"])
		if response.API == "set_user" {
			err = api.SetResponseAuth(&response)
			if err == nil {
				if response.User != nil {
					response.User.Online = true
					response.User.LastActiveOn = time.Now().Unix()
					db.Conn.Model(response.User).Updates(model.User{Online: response.User.Online, LastActiveOn: response.User.LastActiveOn})
					apiresponse.OnlinUsers[response.User.ID] = apiresponse.OnlinUser{
						User: response.User,
						WS:   conn,
					}
				} else {
					util.Log(response.Request)
				}
				if response.User.IsSupportagent {
					for _, online_user := range apiresponse.OnlinUsers {
						message := ` { "online" : ` + strconv.FormatUint(uint64(response.User.ID), 10) + "}"
						online_user.WS.WriteMessage(1, []byte(message))
					}
				}
			}
		} else if response.User == nil {
			var ask = "who are you?"
			err = conn.WriteMessage(messageType, []byte(ask))
		} else {
			response, err = api.JSONApihandler(response, w, r)
		}
		if response.Status == apiresponse.SUCCESS && len(response.Result) > 0 {
			var message_string []byte
			message_string, err = json.Marshal(response.Result)
			apiresponse.OnlinUsers[response.User.ID].WS.WriteMessage(1, message_string)
		}
		if err != nil {
			util.Log(err)
			break
		}
	}

	if response.User != nil && response.User.ID != 0 {
		response.User.Online = false
		response.User.LastActiveOn = time.Now().Unix()
		db.Conn.Model(response.User).Updates(model.User{Online: response.User.Online, LastActiveOn: response.User.LastActiveOn})
		delete(apiresponse.OnlinUsers, response.User.ID)
		if response.User.IsSupportagent {
			for _, online_user := range apiresponse.OnlinUsers {
				message := ` { "offline" : ` + strconv.FormatUint(uint64(response.User.ID), 10) + "}"
				online_user.WS.WriteMessage(1, []byte(message))
			}
		}
	}
}
