package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var reviewresponseAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	util.Log("I got called")
	if _, ok := response.Request["ReviewID"]; ok {
		ReviewID := util.GetUint(response.Request["ReviewID"])
		if ReviewID > 0 {
			if _, ok := response.Request["Response"]; ok {
				Response := util.GetString(response.Request["Response"])

				var reviewresponse model.ReviewResponse
				err = db.Conn.First(&reviewresponse, " `created_by` = ? AND `review_id` = ? ", response.User.ID, ReviewID).Error
				if err != nil {
					reviewresponse = model.ReviewResponse{
						CreatedBy: response.User.ID,
						ReviewID:  ReviewID,
					}
					err = nil
				}
				if reviewresponse.Response != Response {
					var review model.Review
					err = db.Conn.First(&review, ReviewID).Error
					if err == nil {
						if reviewresponse.Response == "Like" {
							util.Log("Likes reduced")
							review.Likes--
						}
						if reviewresponse.Response == "Dislike" {
							util.Log("Dislikes reduced")
							review.Dislikes--
						}
						reviewresponse.Response = Response
						if reviewresponse.Response == "Like" {
							util.Log("Likes incresed")
							review.Likes++
						}
						if reviewresponse.Response == "Dislike" {
							util.Log("Dislikes incresed")
							review.Dislikes++
						}
						err = db.Conn.Save(&review).Error
						if err == nil {
							err = db.Conn.Save(&reviewresponse).Error
							if err == nil {
								response.Result["review"] = review
								response.Status = apiresponse.SUCCESS
								return err
							}
						}
					}
				}
			}
			if err != nil {
				response.Message = "Faile to submit review response"
			}
		}
	}
	return err
}
