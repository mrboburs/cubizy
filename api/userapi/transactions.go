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

type transactionview struct {
	model.Transaction
	UpdatedByName string
}

var transactionsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	if _, ok := response.Request["items"]; ok {
		postTransactions := response.Request["items"].([]interface{})
		messages := make(map[string]int)
		for _, _postTransaction := range postTransactions {
			postTransaction := _postTransaction.(map[string]interface{})
			message := " Added "
			var item model.Transaction
			id := util.GetID(postTransaction)
			if id > 0 {
				err = db.Conn.First(&item, "id = ?", id).Error
				if err == nil {
					message = " Updated "
				}
			}
			if err != nil || id == 0 {
				item = model.Transaction{
					AccountID: response.Account.ID,
					CreatedBy: response.User.ID,
				}
				err = nil
			}
			_, err = item.Update(postTransaction, response.User)
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
				response.Message += strconv.Itoa(value) + "Transactions "
			} else {
				response.Message += "Transaction "
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
				var transaction model.Transaction
				err = db.Conn.First(&transaction, "id = ?", itemid).Error
				if err == nil {
					if !transaction.Accepted {
						err = db.Conn.Unscoped().Delete(&transaction).Error
					} else {
						err = errors.New("can not delete successful transactions")
					}

					if err == nil {
						susccessMessage = " Transaction deleted "
						response.Status = apiresponse.SUCCESS
					} else {
						break
					}
				} else {
					errorMessage = "Some Transactions not found"
				}
			} else {
				invalidIDMessage = " Some Transaction ids are invalid "
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
		var transactions []transactionview

		var query = db.Conn.Model(&model.Transaction{}).Select("transactions.*, users.name AS updated_by_name").Joins("left join users on transactions.updated_by = users.id")
		query.Where("user_id = ?", response.User.ID)
		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal
		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" transactions.title Like ? OR users.name Like ? OR transactions.id Like ? ", SearchStringLike, SearchStringLike, SearchString)
				//response.RecordsFiltered = 3
				query.Count(&response.RecordsFiltered)
			}
		}
		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&transactions).Error
		}
		if err == nil {
			response.Data = transactions
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Transaction{}
			response.Status = apiresponse.FAILED
		}
	}
	return err
}
