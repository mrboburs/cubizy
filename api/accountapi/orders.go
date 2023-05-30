package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var ordersAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var orders []model.Cartitem
	if _, ok := response.Request["items"]; ok {
		postCartitems := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postCartitem := range postCartitems {
			postCartitem := _postCartitem.(map[string]interface{})
			message := " Added "
			var item model.Cartitem
			id := util.GetID(postCartitem)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}

			if err != nil || id == 0 {
				message = " not found " + strconv.FormatUint(uint64(id), 10)
			} else {
				_, err = item.Update(postCartitem, response.User)
				if err != nil {
					if message == " Added " {
						message = "failed to add (" + err.Error() + ")"
					} else {
						message = "failed to update (" + err.Error() + ")"
					}
				} else {
					if _, ok := postCartitem["cancel_order"]; ok {
						Value := util.GetBool(postCartitem["cancel_order"])
						if Value {
							if item.CancelledOn == 0 {
								message, err = item.Cancel("Order Cancled", response.User)
							} else {
								message = " allready cancled"
								err = errors.New(" allready cancled")
							}
						}
					}
				}
			}
			count := messages[message]
			messages[message] = count + 1
		}
		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + " Order items "
			} else {
				response.Message += " Order item "
			}
			response.Message += key
		}
	}
	if err == nil {
		var query = db.Conn.Model(&model.Cartitem{})
		query.Where("order_id > 0")
		if response.Account.AccountType != "admin" {
			query.Where("seler_id = ?", response.Account.ID)
		}
		if _, isFixCondition := response.Request["fix_condition"]; isFixCondition {
			fixConditions := response.Request["fix_condition"].(map[string]interface{})
			if _, ok := fixConditions["ProductID"]; ok {
				ProductID := util.GetUint(fixConditions["ProductID"])
				query.Where("product_id = ?", ProductID)
			}
		}
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" name Like ? OR variation Like ?", SearchStringLike, SearchStringLike)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&orders).Error
		}
		if err == nil {
			response.Data = orders
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Order{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
