package sellerapi

import (
	"cubizy/apiresponse"
	"net/http"
)

var sellerapis map[string]apiresponse.Handler

func init() {
	sellerapis = make(map[string]apiresponse.Handler)
	sellerapis["addresses"] = addressesAPI
	sellerapis["dashbord"] = dashbordAPI
	sellerapis["reviews"] = reviewsAPI
}

// Apihandler will handall all request coming to domain/api/anything
func Apihandler(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if handler, exists := sellerapis[response.API]; exists {
		err = handler(w, r, response)
	}
	return err
}

func HaveSellerApi(API string) bool {
	_, exists := sellerapis[API]
	return exists
}
