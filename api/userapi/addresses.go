package userapi

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

type addressview struct {
	model.Address
	UpdatedByName string
}

var addressesAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postAddresses := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postAddress := range postAddresses {
			postAddress := _postAddress.(map[string]interface{})
			message := " Added "
			var item model.Address
			id := util.GetID(postAddress)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Address{}
				if response.Domain == "seller" {
					item.AccountID = response.Account.ID
				}
				err = nil
			}
			_, err = item.Update(postAddress, response.User)
			if err != nil {
				if message == " Added " {
					message = "failed to add (" + err.Error() + ")"
				} else {
					message = "failed to update (" + err.Error() + ")"
				}
			}
			count := messages[message]
			messages[message] = count + 1
		}

		for key, value := range messages {
			if value > 1 {
				response.Message += strconv.Itoa(value) + "Addresses "
			} else {
				response.Message += "Address "
			}
			response.Message += key
		}
	} else if _, okDelete := response.Request["todelete"]; okDelete {

		todelete := response.Request["todelete"].([]interface{})
		for _, item := range todelete {
			itemid := util.GetUint(item)
			invalidIDMessage := ""
			errorMessage := ""
			susccessMessage := ""
			if itemid > 0 {
				var address model.Address
				err = db.Conn.First(&address, "id = ?", itemid).Error
				if err == nil {
					if address.AccountID == response.Account.ID {
						err = db.Conn.Delete(&address).Error
					} else if address.CreatedBy == response.User.ID {
						var query = db.Conn.Model(&model.Address{})
						query.Where(" addresses.account_id = 0 AND addresses.created_by = ? ", response.User.ID)
						var userAddresses int64 = 0
						query.Count(&userAddresses)
						if userAddresses > 1 {
							err = db.Conn.Delete(&address).Error
						} else {
							response.Message = "User must have atlsit one address"
							err = errors.New(response.Message)
						}
					} else {
						response.Message = "Unauthorized Request"
						err = errors.New("record did not belong to this account")
					}

					if err == nil {
						susccessMessage = " Address deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Addresses not found"
				}
			} else {
				invalidIDMessage = " Some Address ids are invalid "
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
		var addresses []addressview

		var query = db.Conn.Model(&model.Address{}).Select("addresses.*, users.name AS updated_by_name")
		query.Joins("left join users on addresses.updated_by = users.id")

		if response.Domain == "seller" {
			query.Where(" addresses.account_id = ? ", response.Account.ID)
		} else {
			query.Where(" addresses.account_id = 0 AND addresses.created_by = ? ", response.User.ID)
		}

		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" addresses.title Like ? OR locations.country Like ? OR locations.district Like ? OR ocations.locality Like ? OR locations.sub_locality Like ? OR locations.code Like ? ", SearchStringLike, SearchStringLike, SearchStringLike, SearchStringLike, SearchStringLike, SearchStringLike)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&addresses).Error
		}
		if err == nil {
			response.Data = addresses
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Address{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
