package api

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"fmt"
	"net/http"
)

var accountsAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error

	if response.Account != nil && response.Account.ID > 0 {
		var accounts []model.Account
		var query = db.Conn.Model(&model.Account{})
		var account_type = ""
		if _, isAccountType := response.Request["account_type"]; isAccountType {
			account_type = util.GetString(response.Request["account_type"])
		}
		if account_type != "" {
			query.Where(" account_type = ? ", account_type)
		} else {
			query.Where(" account_type != 'admin'")
		}

		var country = ""
		if _, iscountry := response.Request["country"]; iscountry {
			country = util.GetString(response.Request["country"])
		}

		var district = ""
		if _, isdistrict := response.Request["district"]; isdistrict {
			district = util.GetString(response.Request["district"])
			if district == "_" {
				district = ""
			}
		}

		var locality = ""
		if _, islocality := response.Request["locality"]; islocality {
			locality = util.GetString(response.Request["locality"])
		}

		var sublocality = ""
		if _, issublocality := response.Request["sublocality"]; issublocality {
			sublocality = util.GetString(response.Request["sublocality"])
		}

		if country != "" || district != "" || locality != "" || sublocality != "" {
			query.Where(" id IN (SELECT account_id FROM `addresses` WHERE country LIKE ? AND district LIKE ? AND locality LIKE ? AND sub_locality LIKE ? )", "%"+country, "%"+district, "%"+locality, "%"+sublocality)
		}

		query.Count(&response.RecordsTotal)
		response.RecordsFiltered = response.RecordsTotal

		if Search, okSearch := response.Request["search"]; okSearch {
			SearchString := fmt.Sprintf("%v", Search)
			if SearchString != "" {
				SearchStringLike := "%" + SearchString + "%"
				query.Where(" accounts.title Like ? OR accounts.description Like ? OR accounts.keywords Like ? ", SearchStringLike, SearchStringLike, SearchStringLike)
				query.Count(&response.RecordsFiltered)
			}
		}

		db.SetUpQuery(response.Request, query)
		if response.RecordsFiltered > 0 {
			err = query.Scan(&accounts).Error
		}
		if err == nil {
			response.Data = accounts
			response.Status = apiresponse.SUCCESS
		} else {
			response.Data = []model.Account{}
			response.Status = apiresponse.FAILED
		}
	} else {
		response.Message = "empty request"
	}
	return err
}
