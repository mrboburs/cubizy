package accountapi

import (
	"cubizy/apiresponse"
	"cubizy/plugins/db"
	"cubizy/util"
	"net/http"
)

var accountbalanceAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	var total uint

	/*
		var wallet int
		err = db.Conn.Table("transactions").Select("sum(amount) as orderpaid").Where(" accepted IS TRUE AND `for`= 'wallet' AND account_id = ? ", response.Account.ID).Scan(&wallet).Error

		if err != nil {
			response.Message = "Fail to get wallet collection"
			return err
		}

		total = uint(wallet)
	*/

	var orderpaid *int
	err = db.Conn.Table("transactions").Select("sum(amount) as orderpaid").Where(" accepted IS TRUE AND `for`= 'Order' AND account_id = ? ", response.Account.ID).Scan(&orderpaid).Error
	if err != nil {
		response.Message += "Fail to get paid order collection. "
		//return err
		util.Log(err)
		err = nil
	} else {
		util.Log("&orderpaid", orderpaid)
		if orderpaid != nil {
			util.Log("orderpaid", *orderpaid)
		} else {
			util.Log("orderpaid", 0)
		}
	}
	if orderpaid != nil {
		total += uint(*orderpaid)
	}

	var ordercanceled *int
	err = db.Conn.Table("transactions").Select("sum(amount) as ordercanceled").Where(" accepted IS TRUE AND `for`= 'Cancel Item' AND account_id = ? ", response.Account.ID).Scan(&ordercanceled).Error

	if err != nil {
		response.Message += "Fail to get cancel order item expense. "
		//return err
		util.Log(err)
		err = nil
	} else {
		util.Log("&ordercanceled", ordercanceled)
		if ordercanceled != nil {
			util.Log("ordercanceled", *ordercanceled)
		} else {
			util.Log("ordercanceled", 0)
		}
	}

	if ordercanceled != nil {
		total -= uint(*ordercanceled)
	}

	var orderdelivered *int
	err = db.Conn.Table("transactions").Select("sum(amount) as orderdelivered").Where(" accepted IS TRUE AND `for`= 'Delivered Item' AND account_id = ? ", response.Account.ID).Scan(&orderdelivered).Error

	if err != nil {
		response.Message += "Fail to get delivered order item expense. "
		//return err
		util.Log(err)
		err = nil
	} else {
		util.Log("&orderdelivered", orderdelivered)
		if orderdelivered != nil {
			util.Log("orderdelivered", *orderdelivered)
		} else {
			util.Log("orderdelivered", 0)
		}
	}

	if orderdelivered != nil {
		total -= uint(*orderdelivered)
	}

	var itemdelivered *int
	err = db.Conn.Table("transactions").Select("sum(amount) as itemdelivered").Where(" accepted IS TRUE AND `for`= 'Delivered Item' AND seller_id = ? ", response.Account.ID).Scan(&itemdelivered).Error

	if err != nil {
		response.Message += "Fail to get delivered order item collection. "
		//return err
		util.Log(err)
		err = nil
	} else {
		util.Log("&itemdelivered", itemdelivered)
		if itemdelivered != nil {
			util.Log("itemdelivered", itemdelivered)
		} else {
			util.Log("itemdelivered", 0)
		}
	}

	if itemdelivered != nil {
		total += uint(*itemdelivered)
	}
	response.Account.Wallet = uint(total)
	err = db.Conn.Save(response.Account).Error
	if err != nil {
		response.Message += "Fail to update account balance. "
		return err
	} else {
		response.Message += "Updated account balance. "
		response.Status = apiresponse.SUCCESS
	}
	return err
}
