package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

type ReviewView struct {
	model.Review
	Name  string
	Photo string
}

var reviewAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var flag, added bool

	if _, ok := response.Request["AccountID"]; ok {
		AccountID := util.GetUint(response.Request["AccountID"])
		if AccountID > 0 {
			var ProductID uint = 0
			var review *model.Review
			if _, ok := response.Request["ProductID"]; ok {
				ProductID = util.GetUint(response.Request["ProductID"])
			}
			err = db.Conn.First(&review, " `created_by` = ? AND `account_id` = ? AND `product_id` = ?  ", response.User.ID, AccountID, ProductID).Error
			if err != nil {
				err = nil
				added = true
				review = nil
			}
			if _, ok := response.Request["Review"]; ok {
				if review == nil {
					review = &model.Review{
						CreatedBy: response.User.ID,
						AccountID: AccountID,
						ProductID: ProductID,
					}
				}
				flag, err = review.Update(response.Request, response.User)
			}
			if err == nil {
				response.Result["review"] = review
				response.Status = apiresponse.SUCCESS
				if flag {
					if added {
						response.Message = "Review Added"
					} else {
						response.Message = "Review Updated"
					}
				} else {
					response.Message = ""
				}
			} else {
				response.Message = "Action Failed"
			}
		}
	}
	return err
}
