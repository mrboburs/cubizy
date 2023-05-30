package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
	"strconv"
)

var wishlistAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["products"]; ok {
		postWishlist := response.Request["products"].([]interface{})
		messages := make(map[string]int)
		for _, _postWishItem := range postWishlist {
			postWishItem := util.GetUint(_postWishItem)
			message := " Added "
			var item model.Wishlist
			if postWishItem > 0 {
				err = db.Conn.First(&item, "product_id = ? AND created_by", postWishItem, response.User.ID).Error
				if err != nil {
					item = model.Wishlist{
						ProductID: postWishItem,
						CreatedBy: response.User.ID,
					}
					err = db.Conn.Create(&item).Error
					if err != nil {
						message = "failed to add (" + err.Error() + ")"
						util.Log("While adding wishlists item ")
						util.Log(err)
						err = nil
					}
				}
			}
			count := messages[message]
			messages[message] = count + 1
		}

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + " Wishlist "
			} else {
				response.Message += "Wishlist "
			}
			response.Message += key
		}
		response.User.UpdateWishlistCount()
	} else if _, okDelete := response.Request["todelete"]; okDelete {

		todelete := response.Request["todelete"].([]interface{})
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			susccessMessage := ""
			if itemid > 0 {
				var wishlist model.Wishlist
				err = db.Conn.First(&wishlist, "product_id = ? AND created_by", itemid, response.User.ID).Error
				if err == nil {
					err = db.Conn.Unscoped().Delete(&wishlist).Error
					if err == nil {
						susccessMessage = " Wishlist deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Wishlist not found"
				}
			} else {
				invalidIDMessage = " Some Wishlist ids are invalid "
			}
			response.Message = ""
			if susccessMessage != "" {
				response.Message += susccessMessage
			}
			if errorMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += errorMessage
			}
			if invalidIDMessage != "" {
				if response.Message != "" {
					response.Message += ", "
				}
				response.Message += invalidIDMessage
			}
		}
	}

	if err == nil {
		var wishlist []uint
		err = db.Conn.Model(&model.Wishlist{}).Pluck("product_id", &wishlist).Error
		if err == nil {
			response.Data = wishlist
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Wishlist{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
