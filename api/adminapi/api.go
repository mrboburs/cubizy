package adminapi

import (
	"cubizy/apiresponse"
	"errors"
	"net/http"
)

var adminapis map[string]apiresponse.Handler

func init() {
	adminapis = make(map[string]apiresponse.Handler)
	adminapis["settings"] = settingsAPI
	adminapis["locations"] = locationsAPI
	adminapis["users"] = usersAPI
	adminapis["blogcategories"] = blogcategoriesAPI
	adminapis["accounts"] = accountsAPI
	adminapis["dashbord"] = dashbordAPI
	adminapis["publishedthemes"] = publishedthemesAPI
	adminapis["loginpagesliders"] = loginpageslidersAPI
	adminapis["attributes"] = attributesAPI
	adminapis["reviews"] = reviewsAPI
	adminapis["eventsx24"] = eventsx24API
}

// Apihandler will handall all request coming to domain/api/anything
func Apihandler(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if !response.User.IsAdmin && !response.User.IsSuperAdmin {
		response.Message = "Unauthorized Request"
		err = errors.New("user is not admin")
	} else if handler, exists := adminapis[response.API]; exists {
		err = handler(w, r, response)
	}

	return err
}

func HaveAdminApi(API string) bool {
	_, exists := adminapis[API]
	return exists
}
