package userapi

import (
	"cubizy/apiresponse"
	"net/http"
)

var userapis map[string]apiresponse.Handler

func init() {
	userapis = make(map[string]apiresponse.Handler)
	userapis["logout"] = logoutAPI
	userapis["getpresignedputurl"] = getpresignedputurlAPI
	userapis["getpresignedphotourl"] = getpresignedphotourlAPI
	userapis["files"] = filesAPI
	userapis["sendvirificationemail"] = sendvirificationemailAPI
	userapis["verifyemail"] = verifyemailAPI
	userapis["sendvirificationmobile"] = sendvirificationmobileAPI
	userapis["verifymobile"] = verifymobileAPI
	userapis["account"] = accountAPI
	userapis["deletefile"] = deletefileAPI
	userapis["me"] = meAPI
	userapis["setpassword"] = setpasswordAPI
	userapis["setquestionanswers"] = setquestionanswersAPI
	userapis["review"] = reviewAPI
	userapis["reviewresponse"] = reviewresponseAPI
	userapis["chatlist"] = chatlistAPI
	userapis["chathistory"] = chathistoryAPI
	userapis["addresses"] = addressesAPI
	userapis["wishlist"] = wishlistAPI
	userapis["transactions"] = transactionsAPI
	userapis["cartitems"] = cartitemsAPI
	userapis["addfunds"] = addfundsAPI
	userapis["order"] = orderAPI
	userapis["orders"] = ordersAPI
}

// Apihandler will handall all request coming to domain/api/anything
func Apihandler(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if handler, exists := userapis[response.API]; exists {
		err = handler(w, r, response)
	}
	return err
}

func HaveUserApi(API string) bool {
	_, exists := userapis[API]
	return exists
}
