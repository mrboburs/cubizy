package accountapi

import (
	"cubizy/apiresponse"
	"net/http"
)

type accounthandler func(http.ResponseWriter, *http.Request, *apiresponse.Response) error

var accountapis map[string]accounthandler

func init() {
	accountapis = make(map[string]accounthandler)
	accountapis["pages"] = pagesAPI
	accountapis["accountreviews"] = reviewsAPI
	accountapis["accountaddress"] = accountaddressAPI
	accountapis["themesettings"] = themesettingsAPI
	accountapis["categories"] = categoriesAPI
	accountapis["subcategories"] = subcategoriesAPI
	accountapis["childcategories"] = childcategoriesAPI
	accountapis["updatetemplate"] = updatetemplateAPI
	accountapis["themes"] = themesAPI
	accountapis["publish"] = publishAPI
	accountapis["publishedthemes"] = publishedthemesAPI
	accountapis["transactions"] = transactionsAPI
	accountapis["products"] = productsAPI
	accountapis["product"] = productAPI
	accountapis["description"] = descriptionAPI
	accountapis["orders"] = ordersAPI
	accountapis["accountbalance"] = accountbalanceAPI
	accountapis["blogs"] = blogsAPI
	accountapis["export_theme"] = exportThemeAPI
}

// Apihandler will handall all request coming to domain/api/anything
func Apihandler(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if handler, exists := accountapis[response.API]; exists {
		err = handler(w, r, response)
	}
	return err
}

func HaveAccountApi(API string) bool {
	_, exists := accountapis[API]
	return exists
}
