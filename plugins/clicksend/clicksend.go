package clicksend

import (
	"bytes"
	"cubizy/keys"
	"cubizy/model"
	"cubizy/util"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	baseurl  = "https://rest.clicksend.com/v3/sms/send"
	username = "emailvikrammali"
	password = "20A4D822-8808-7D1F-8781-0DAC865648AD"
)

// Message is object for SMS message
type Message struct {
	Body string
	To   string
}

// call test api
func SendSMS(to, message string) error {
	var err error
	baseurl = model.GetSetting(keys.ClicksendBaseURL, baseurl)
	username = model.GetSetting(keys.ClicksendUsername, username)
	password = model.GetSetting(keys.ClicksendPassword, password)

	var messages [1]map[string]string
	messages[0] = map[string]string{
		"body": message,
		"to":   to,
	}
	messages_obj := map[string]interface{}{
		"messages": messages,
	}
	postBody, _ := json.Marshal(messages_obj)
	responseBody := bytes.NewBuffer(postBody)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	//util.Log(baseurl)
	//util.Log(string(postBody))
	var req *http.Request
	req, err = http.NewRequest("POST", baseurl, responseBody)
	if err != nil {
		util.Log("Got error %s", err.Error())
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		util.Log("Got error %s", err.Error())
	}
	defer resp.Body.Close()
	//Read the response body
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	util.Log(sb)
	if strings.Contains(sb, `"response_code":"SUCCESS"`) {
		err = nil
	} else {
		message = "failed to send sms"
		messages := strings.Split(sb, `"response_msg":"`)
		if len(messages) > 1 {
			messages = strings.Split(messages[1], `"`)
			if len(messages) > 0 {
				message = messages[0]
			}
		}
		err = errors.New(message)
	}
	return err
}
