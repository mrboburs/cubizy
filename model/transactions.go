package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Transaction will all informetion of Transaction,
// Status  0 : pending; 1 : successful; 2: failed; 3: cancelled
// For : wallet, order, return, withdrawn
type Transaction struct {
	gorm.Model

	Amount        uint `gorm:"default:0"`
	Method        string
	TransactionID string
	UserID        uint
	AccountID     uint
	SellerID      uint
	For           string
	OrderID       uint
	ItemID        uint
	Status        string
	Note          string
	Accepted      bool
	CreatedBy     uint
	UpdatedBy     uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Transaction{})
	db.Conn.Exec("DELETE FROM `transactions` WHERE amount IS NULL")
}

// Update will update product by given post argumnets
func (transaction *Transaction) Update(TransactionMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool

	if _, ok := TransactionMap["Amount"]; ok {
		Value := util.GetUint(TransactionMap["Amount"])
		if Value != transaction.Amount {
			transaction.Amount = Value
			flag = true
		}
	}

	if _, ok := TransactionMap["Method"]; ok {
		Value := util.GetString(TransactionMap["Method"])
		if Value != transaction.Method {
			transaction.Method = Value
			flag = true
		}
	}

	if _, ok := TransactionMap["TransactionID"]; ok {
		Value := util.GetString(TransactionMap["TransactionID"])
		if Value != transaction.TransactionID {
			transaction.TransactionID = Value
			flag = true
		}
	}

	if _, ok := TransactionMap["Note"]; ok {
		Value := util.GetString(TransactionMap["Note"])
		if Value != transaction.Note {
			transaction.Note = Value
			flag = true
		}
	}

	if _, ok := TransactionMap["For"]; ok {
		Value := util.GetString(TransactionMap["For"])
		if Value != transaction.For {
			transaction.For = Value
			flag = true
		}
	}

	if transaction.For == "" {
		transaction.For = "wallet"
	}

	if transaction.For == "wallet" {
		transaction.AccountID = SuperAdminAccount.ID
	}

	if UpdatedBy != nil && UpdatedBy.IsAdmin {
		if _, ok := TransactionMap["Status"]; ok {
			Value := util.GetString(TransactionMap["Status"])
			if Value != transaction.Status {
				transaction.Status = Value
				flag = true
			}
		}
		if _, ok := TransactionMap["Accepted"]; ok {
			Value := util.GetBool(TransactionMap["Accepted"])
			if Value != transaction.Accepted {
				transaction.Accepted = Value
				flag = true
			}
		}
	}

	if transaction.ID == 0 {
		if _, ok := TransactionMap["UserID"]; ok {
			Value := util.GetUint(TransactionMap["UserID"])
			if Value != transaction.UserID {
				transaction.UserID = Value
				flag = true
			}
		}

		if transaction.UserID == 0 && !transaction.Accepted {
			transaction.UserID = transaction.CreatedBy
		}
	}

	if transaction.Status == "" {
		transaction.Status = "pending"
	}

	if flag {
		if transaction.Status == "successful" {
			transaction.Accepted = true
		}

		if (transaction.Status == "pending" || transaction.Status == "") && transaction.Accepted {
			transaction.Status = "successful"
		}

		if transaction.Method == "" {
			err = errors.New("method can not be empty")
		} else if transaction.TransactionID == "" {
			err = errors.New("transaction id can not be empty")
		} else if transaction.Amount == 0 {
			err = errors.New("amount can not be empty")
		} else {
			if UpdatedBy != nil {
				transaction.UpdatedBy = UpdatedBy.ID
			}
			if transaction.ID == 0 {
				transaction.CreatedBy = UpdatedBy.ID
				var duplicate Transaction
				err = db.Conn.First(&duplicate, " `amount` = ? && `transaction_id` = ? && `method` = ?  ", transaction.Amount, transaction.TransactionID, transaction.Method).Error
				if err == nil {
					if duplicate.UserID == 0 {
						duplicate.UserID = transaction.UserID
						duplicate.UpdatedBy = UpdatedBy.ID
						err = db.Conn.Save(&duplicate).Error
						if err == nil {
							transaction.ID = duplicate.ID
							transaction.Accepted = duplicate.Accepted
							transaction.Status = duplicate.Status
							transaction.Amount = duplicate.Amount
							transaction.Note = duplicate.Note
							transaction.For = duplicate.For
							transaction.CreatedAt = duplicate.CreatedAt
							transaction.UpdatedAt = duplicate.UpdatedAt
							transaction.CreatedBy = duplicate.CreatedBy
						}
					} else {
						err = errors.New("duplicate transaction code")
					}
				} else {
					err = db.Conn.Create(&transaction).Error
				}
			} else if transaction.Amount == 0 {
				err = errors.New(" Amount can not be empty")
			} else {
				err = db.Conn.Save(&transaction).Error
			}
		}
		//  && transaction.Accepted
		if err == nil && transaction.UserID > 0 {
			var balancecount int64
			var balances []map[string]interface{}
			var balance uint = 0
			// AND `accepted` IS TRUE
			var balancequery = db.Conn.Model(&Transaction{}).Where("`for` = 'wallet' AND `accepted` IS TRUE AND user_id = ?", transaction.UserID)
			err = balancequery.Count(&balancecount).Error
			if err != nil {
				util.Log("error in counting balance")
				util.Log(err)
				return flag, err
			}
			if balancecount > 0 {
				err = balancequery.Select("SUM(amount) AS balance").Scan(&balances).Error
				if err != nil {
					util.Log("error in geting balance")
					util.Log(err)
					return flag, err
				}
				balance = util.GetUint(balances[0]["balance"])
			} else {
				balance = 0
			}

			var expenses []map[string]interface{}
			//
			var expensecount int64
			var expense uint
			var expensesquery = db.Conn.Model(&Transaction{}).Where("`for` = 'order%' AND `accepted` IS TRUE AND user_id = ?", transaction.UserID)
			err = expensesquery.Count(&expensecount).Error
			if err != nil {
				util.Log("error in counting balance")
				util.Log(err)
				return flag, err
			}

			if expensecount < 0 {
				err = expensesquery.Select("SUM(amount) AS balance").Scan(&expenses).Error
				if err != nil {
					util.Log("error in geting expenses")
					util.Log(err)
					return flag, err
				}
				expense = util.GetUint(expenses[0]["expense"])
			} else {
				expense = 0
			}
			// get all Accepted wallet transactions for user id
			balance = balance - expense

			var user *User
			if transaction.UserID == UpdatedBy.ID {
				user = UpdatedBy
			} else {
				err = db.Conn.First(&user, transaction.UserID).Error
				if err != nil {
					util.Log("error in geting user")
					util.Log(err)
					return flag, err
				}
			}
			user.Wallet = balance
			err = db.Conn.Save(user).Error
			if err != nil {
				util.Log("error in updating user balance")
				util.Log(err)
				return flag, err
			}
		}
	}
	return flag, err
}
