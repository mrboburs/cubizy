package api

import (
	"cubizy/api/accountapi"
	"cubizy/api/adminapi"
	"cubizy/api/sellerapi"
	"cubizy/api/userapi"
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

var apis map[string]apiresponse.Handler

func init() {
	apis = make(map[string]apiresponse.Handler)
	apis["test"] = testAPI
	apis["login"] = loginapi
	apis["resetcode"] = resetcodeAPI
	apis["resetpassword"] = resetpasswordAPI
	apis["register"] = registerAPI
	apis["countries"] = countriesAPI
	apis["districts"] = districtsAPI
	apis["localities"] = localitiesAPI
	apis["sublocalities"] = sublocalitiesAPI
	apis["codes"] = codesAPI
	apis["pages"] = pagesAPI
	apis["page"] = pageAPI
	apis["blogcategories"] = blogcategoriesAPI
	apis["blogs"] = blogsAPI
	apis["blog"] = blogAPI
	apis["accounts"] = accountsAPI
	apis["account"] = accountAPI
	apis["locations"] = locationsAPI
	apis["questions"] = questionsAPI
	apis["activethemes"] = themesAPI
	apis["loginpagesliders"] = loginpageslidersAPI
	apis["allcategories"] = allcategoriesAPI
	apis["allsubcategories"] = allsubcategoriesAPI
	apis["allchildcategories"] = allchildcategoriesAPI
	apis["attributes"] = attributesAPI
	apis["categories"] = categoriesAPI
	apis["products"] = productsAPI
	apis["description"] = descriptionAPI
	apis["wishlistproducts"] = wishlistproductsAPI
	apis["reviews"] = reviewsAPI
	apis["productsuggestions"] = productsuggestionsAPI
	apis["cities"] = citiesAPI
}

// JSONApihandler will handall all request coming to websocket/domain/api/anything
func JSONApihandler(response apiresponse.Response, w http.ResponseWriter, r *http.Request) (apiresponse.Response, error) {

	var err error

	found := false
	if response.User != nil {
		switch response.Domain {
		case "admin":
			if !found && response.Account != nil && response.Account.ID > 0 && response.Account.AccountType == "admin" {
				if adminapi.HaveAdminApi(response.API) {
					found = true
					err = adminapi.Apihandler(w, r, &response)
				}
			}
			if !found && response.Account != nil && response.Account.ID > 0 && (response.User.IsAdmin || response.User.IsSuperAdmin) {
				if accountapi.HaveAccountApi(response.API) {
					found = true
					err = accountapi.Apihandler(w, r, &response)
				}
			}
		case "seller":
			if !found && response.Account != nil && response.Account.ID > 0 && response.Account.AccountType == "seller" && (response.Account.Status > 9 || response.API == "dashbord") {
				if sellerapi.HaveSellerApi(response.API) {
					found = true
					err = sellerapi.Apihandler(w, r, &response)
				}
			}
			if !found && response.Account != nil && response.Account.ID > 0 && (response.User.SellerAccountID == response.Account.ID || response.User.IsAdmin || response.User.IsSuperAdmin) {
				if accountapi.HaveAccountApi(response.API) {
					found = true
					err = accountapi.Apihandler(w, r, &response)
				}
			}
		default:
			found = false
		}
		if !found {
			if userapi.HaveUserApi(response.API) {
				found = true
				err = userapi.Apihandler(w, r, &response)
			}
		}
	}
	if !found {
		if handler, exists := apis[response.API]; exists {
			err = handler(w, r, &response)
		} else {
			err = errors.New("no API named " + response.API)
		}
	}
	return response, err
}

func RestApihandler(w http.ResponseWriter, r *http.Request) {

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

	apiCalled := r.URL.Path[len("/api/"):]
	urlparts := strings.Split(apiCalled, "/")
	apiCalled = urlparts[len(urlparts)-1]
	//util.Log(apiCalled)

	response := apiresponse.Response{
		StartTime: time.Now(),
		Message:   "",
		Status:    apiresponse.FAILED,
		Result:    make(map[string]interface{}),
		URLparts:  r.URL.Path[len("/api/"+apiCalled):],
		Domain:    SubDomain,
		API:       apiCalled,
	}

	json.NewDecoder(r.Body).Decode(&response.Request)
	if response.Request == nil {
		response.Request = make(map[string]interface{})
	}
	var err error
	err = SetResponseAuth(&response)

	if err == nil {
		response, err = JSONApihandler(response, w, r)
	}

	if err != nil {
		if response.Message == "" {
			response.Message = "Action " + response.API + " failed, " + err.Error()
		}
		response.Result["Error"] = err.Error()
	}

	if err == nil && response.Message == "" {
		response.Status = apiresponse.SUCCESS
	}

	if response.Account != nil && response.Account.ID == 0 {
		response.Account = nil
	}
	if response.User == nil || response.User.Token == "" || response.User.ID == 0 {
		response.User = nil
	} else {
		if response.User.Password != "" {
			response.User.PasswordSet = true //"set"
		}
	}

	response.EndTime = time.Now()
	response.ResponseTime = response.EndTime.Sub(response.StartTime)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SetResponseAuth(response *apiresponse.Response) error {
	var err error

	if response.Domain != "admin" && response.Domain != "seller" {
		if _, ok := response.Request["account_id"]; ok {
			account_id := util.GetInt(response.Request["account_id"])
			var account *model.Account
			err = db.Conn.First(&account, account_id).Error
			if err == nil {
				response.Account = account
			} else {
				util.Log("api loading", err.Error())
				err = nil
			}
		}
	}

	if _, ok := response.Request["token"]; ok {
		token := response.Request["token"].(string)
		var user *model.User
		if err == nil && user == nil {
			user, err = model.GetUserByToken(token)
		}
		if err == nil && user != nil {

			for _, OnlinUser := range apiresponse.OnlinUsers {
				if OnlinUser.User.ID == user.ID {
					OnlinUser.User = user
				}
			}

			if user.Email == util.Settings.SuperAdmin {
				user.IsSuperAdmin = true
			} else {
				user.IsSuperAdmin = false
			}
			if user.IsSuperAdmin {
				user.IsAdmin = true
			}
			if user.ID > 0 {
				response.User = user
			} else {
				response.User = nil
			}
			if response.Account == nil {
				response.Account = model.GetAccount(user, response.Domain)
				if response.Account != nil {
					response.Account.LastActiveOn = time.Now().Unix()
					db.Conn.Save(response.Account)
				}
			}
		}
	}
	return err
}
