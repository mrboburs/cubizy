package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Cartitem will all informetion of Cartitem
type Cartitem struct {
	gorm.Model
	AccountID uint
	ProductID uint
	Name      string
	Logo      string
	Variation string
	SKU       string
	Quantity  uint
	Price     uint
	Cost      uint
	CreatedBy uint
	UpdatedBy uint

	PriceTotal     uint
	CostTotal      uint
	ShippingImage  string
	ShippingMethod string
	ShippingPrice  uint
	EDTMin         uint
	EDTMax         uint
	ShippingCost   uint
	OrderID        uint

	TrackingID        string
	SelerID           uint
	RequestedToCalcel int64

	TransactionID   uint
	CancelTrxID     uint
	CancelledOn     int64
	DispatchedOn    int64
	DeliveredOn     int64
	ShippingDetails string

	Status string

	product Product
	stocks  []Stock
}

type TrackingDetail struct {
	CreatedAt int64
	Note      string
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Cartitem{})
	db.Conn.Exec("UPDATE cartitems SET account_id = (SELECT orders.account_id FROM orders WHERE orders.id = cartitems.order_id )")
	db.Conn.Exec("UPDATE cartitems SET seler_id = (SELECT products.account_id FROM products WHERE products.id = cartitems.product_id )")
}

// Update will update product by given post argumnets
func (cartitem *Cartitem) Update(CartitemMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if cartitem.OrderID > 0 {
		if !UpdatedBy.IsAdmin && !UpdatedBy.IsSuperAdmin && !(UpdatedBy.SellerAccountID == cartitem.SelerID || UpdatedBy.SellerAccountID == cartitem.AccountID) {
			return flag, errors.New("invalid updater")
		}
	}
	cartitem.UpdatedBy = UpdatedBy.ID
	if cartitem.OrderID == 0 {
		if _, ok := CartitemMap["ProductID"]; ok {
			Value := util.GetUint(CartitemMap["ProductID"])
			if Value != cartitem.ProductID {
				cartitem.ProductID = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["Logo"]; ok {
			Value := util.GetString(CartitemMap["Logo"])
			if Value != cartitem.Logo {
				cartitem.Logo = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["Name"]; ok {
			Value := util.GetString(CartitemMap["Name"])
			if Value != cartitem.Name {
				cartitem.Name = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["Variation"]; ok {
			Value := util.GetString(CartitemMap["Variation"])
			if Value != cartitem.Variation {
				cartitem.Variation = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["SKU"]; ok {
			Value := util.GetString(CartitemMap["SKU"])
			if Value != cartitem.SKU {
				cartitem.SKU = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["Quantity"]; ok {
			Value := util.GetUint(CartitemMap["Quantity"])
			if Value != cartitem.Quantity {
				cartitem.Quantity = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["Price"]; ok {
			Value := util.GetUint(CartitemMap["Price"])
			if Value != cartitem.Price {
				cartitem.Price = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["Cost"]; ok {
			Value := util.GetUint(CartitemMap["Cost"])
			if Value != cartitem.Cost {
				cartitem.Cost = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["PriceTotal"]; ok {
			Value := util.GetUint(CartitemMap["PriceTotal"])
			if Value != cartitem.PriceTotal {
				cartitem.PriceTotal = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["CostTotal"]; ok {
			Value := util.GetUint(CartitemMap["CostTotal"])
			if Value != cartitem.CostTotal {
				cartitem.CostTotal = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["ShippingImage"]; ok {
			Value := util.GetString(CartitemMap["ShippingImage"])
			if Value != cartitem.ShippingImage {
				cartitem.ShippingImage = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["ShippingMethod"]; ok {
			Value := util.GetString(CartitemMap["ShippingMethod"])
			if Value != cartitem.ShippingMethod {
				cartitem.ShippingMethod = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["EDTMin"]; ok {
			Value := util.GetUint(CartitemMap["EDTMin"])
			if Value != cartitem.EDTMin {
				cartitem.EDTMin = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["EDTMax"]; ok {
			Value := util.GetUint(CartitemMap["EDTMax"])
			if Value != cartitem.EDTMax {
				cartitem.EDTMax = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["ShippingPrice"]; ok {
			Value := util.GetUint(CartitemMap["ShippingPrice"])
			if Value != cartitem.ShippingPrice {
				cartitem.ShippingPrice = Value
				flag = true
			}
		}

		if _, ok := CartitemMap["ShippingCost"]; ok {
			Value := util.GetUint(CartitemMap["ShippingCost"])
			if Value != cartitem.ShippingCost {
				cartitem.ShippingCost = Value
				flag = true
			}
		}
	}
	if cartitem.OrderID > 0 {
		var autoTrackingDetail string
		if _, ok := CartitemMap["Status"]; ok {
			Value := util.GetString(CartitemMap["Status"])
			if Value != cartitem.Status {
				cartitem.Status = Value
				autoTrackingDetail += Value + ", "
				flag = true
			}
		}

		if _, ok := CartitemMap["TrackingID"]; ok {
			Value := util.GetString(CartitemMap["TrackingID"])
			if Value != cartitem.TrackingID {
				cartitem.TrackingID = Value
				autoTrackingDetail += " Recived TrackingID : " + Value + ", "
				flag = true
			}
		}

		if _, ok := CartitemMap["TrackingDetail"]; ok {
			trackingDetail := autoTrackingDetail + ", " + util.GetString(CartitemMap["TrackingDetail"])
			var ShippingDetails []TrackingDetail
			if len(cartitem.ShippingDetails) > 0 {
				json_err := json.Unmarshal([]byte(cartitem.ShippingDetails), &ShippingDetails)
				if json_err != nil {
					ShippingDetails = make([]TrackingDetail, 0)
				}
			}
			ShippingDetails = append(ShippingDetails, TrackingDetail{
				CreatedAt: time.Now().Unix(),
				Note:      trackingDetail,
			})
			ShippingDetailJson, json_err := json.Marshal(ShippingDetails)
			if json_err == nil {
				cartitem.ShippingDetails = string(ShippingDetailJson)
				flag = true
			}
		}

		if cartitem.Status == "Dispatched" && cartitem.TrackingID != "" && cartitem.DispatchedOn == 0 {
			cartitem.DispatchedOn = time.Now().Unix()
			flag = true
		}

		if cartitem.Status == "Delivered" && cartitem.DeliveredOn == 0 {
			cartitem.DeliveredOn = time.Now().Unix()
			flag = true
		}
	}
	if flag {
		if cartitem.ProductID == 0 {
			err = errors.New(" Product can not be empty")
		} else if cartitem.SKU == "" {
			err = errors.New(" SKU can not be empty")
		} else {
			if cartitem.ID == 0 {
				cartitem.CreatedBy = UpdatedBy.ID
				var duplicate Cartitem
				err = db.Conn.First(&duplicate, " `product_id` = ? AND `sku` = ? AND `created_by` = ? AND order_id = ? ", cartitem.ProductID, cartitem.SKU, cartitem.CreatedBy, cartitem.OrderID).Error
				if err == nil {
					if cartitem.Quantity == 0 {
						err = db.Conn.Unscoped().Delete(&duplicate).Error
					} else {
						cartitem.ID = duplicate.ID
						cartitem.CreatedAt = duplicate.CreatedAt
						cartitem.UpdatedAt = duplicate.UpdatedAt
						err = db.Conn.Save(&cartitem).Error
					}
				} else {
					err = db.Conn.Create(&cartitem).Error
				}
			} else {
				err = db.Conn.Save(&cartitem).Error
			}

			if err == nil && cartitem.AccountID != cartitem.SelerID && cartitem.DeliveredOn > 0 && (cartitem.CancelledOn == 0 || cartitem.DispatchedOn < cartitem.CancelledOn) {

				var transaction Transaction
				err = db.Conn.First(&transaction, " account_id = ? AND seller_id = ? AND item_id  ", cartitem.AccountID, cartitem.SelerID, cartitem.ID).Error

				if err != nil {
					var cost uint = 0
					cost += cartitem.ShippingCost
					if cartitem.CancelledOn == 0 {
						cost += cartitem.CostTotal
					}
					var orderaccount Account

					err = db.Conn.First(&orderaccount, " id =? ", cartitem.AccountID).Error

					if err != nil {
						err = errors.New("Fail to get account of order : " + err.Error())
						return flag, err
					}

					var selleraccount Account

					err = db.Conn.First(&selleraccount, " id =  ?", cartitem.SelerID).Error

					if err != nil {
						err = errors.New("Fail to get seller account of order item : " + err.Error())
						return flag, err
					}

					transaction := Transaction{
						AccountID:     orderaccount.ID,
						SellerID:      selleraccount.ID,
						CreatedBy:     UpdatedBy.ID,
						UpdatedBy:     UpdatedBy.ID,
						Amount:        cost,
						Method:        "Wallet",
						TransactionID: uuid.New().String(),
						For:           "Delivered Item",
						OrderID:       cartitem.OrderID,
						ItemID:        cartitem.ID,
						Status:        "successful",
						Accepted:      true,
					}

					tx := db.Conn.Begin()
					err = tx.Create(&transaction).Error
					if err != nil {
						tx.Rollback()
						err = errors.New("fail to update transaction to seller of order item : " + err.Error())
						return flag, err
					}
					orderaccount.Wallet -= cost
					err = tx.Save(&orderaccount).Error
					if err != nil {
						tx.Rollback()
						err = errors.New("fail to update orderaccount balance for delivered order item : " + err.Error())
						return flag, err
					}
					selleraccount.Wallet += cost
					err = tx.Save(&selleraccount).Error
					if err != nil {
						tx.Rollback()
						err = errors.New("fail to update selleraccount balance for delivered order item : " + err.Error())
						return flag, err
					}
					cartitem.TransactionID = transaction.ID
					err = tx.Save(cartitem).Error
					if err != nil {
						tx.Rollback()
						err = errors.New("fail to update order item transaction ID : " + err.Error())
						return flag, err
					}
					tx.Commit()
				}
			}
		}
	} else {
		util.Log("Nothing to update")
	}
	return flag, err
}

func (cartitem *Cartitem) Cancel(trackingDetail string, UpdatedBy *User) (string, error) {
	var message string
	var err error

	if cartitem.Status == "" || cartitem.UpdatedBy == UpdatedBy.ID {

		if cartitem.DeliveredOn > 0 {
			message = "Can not cancel order after get delivered"
			err = errors.New("can not cancel order after get delivered")
			return message, err
		}

		// will cancel item
		var product Product
		tx := db.Conn.Begin()

		util.Log(time.Now(), "transaction started")
		err = tx.First(&product, "id = ?", cartitem.ProductID).Error

		if err == nil {

			var stocks []Stock
			err = json.Unmarshal([]byte(product.Stock), &stocks)
			if err != nil {
				util.Log(err)
			}
			for i := 0; i < len(stocks); i++ {
				stock := stocks[i]
				if stock.SKU == cartitem.SKU {
					stock.Quantity = stock.Quantity + cartitem.Quantity
					if product.Quantity < stock.Quantity {
						product.Quantity = stock.Quantity
					}
					break
				}
			}

			var stocks_json []byte
			stocks_json, err = json.Marshal(cartitem.stocks)
			if err != nil {
				message = "Unable to update stock in one item"
				return message, err
			}

			cartitem.product.Stock = string(stocks_json)
			err = tx.Save(&cartitem.product).Error

		} else {
			err = nil
		}

		if err != nil {
			tx.Rollback()
			message = "Enable to update products stock in order item"
			return message, err
		}
		var account Account

		err = tx.First(&account, " id = ? ", cartitem.AccountID).Error

		if err != nil {
			tx.Rollback()
			message = "Failed to get order account"
			return message, err
		}

		transaction := Transaction{
			AccountID:     account.ID,
			UserID:        cartitem.CreatedBy,
			CreatedBy:     UpdatedBy.ID,
			UpdatedBy:     UpdatedBy.ID,
			Amount:        cartitem.CostTotal,
			Method:        "Wallet",
			TransactionID: uuid.New().String(),
			For:           "Cancel Item",
			OrderID:       cartitem.OrderID,
			ItemID:        cartitem.ID,
			Status:        "successful",
			Accepted:      true,
		}
		err = tx.Create(&transaction).Error

		if err != nil {
			tx.Rollback()
			message = "Failed to complete refund payment transaction"
		}

		cost := cartitem.CostTotal
		if cartitem.TrackingID == "" {
			cost += cartitem.ShippingCost
		}
		var user User
		if UpdatedBy.ID == cartitem.CreatedBy {
			user = *UpdatedBy
		} else {
			err = tx.First(&user, cartitem.CreatedBy).Error
		}
		if err == nil {
			user.Wallet = user.Wallet + cost
			err = tx.Save(user).Error
		}
		if err != nil {
			tx.Rollback()
			message = "Failed to update user balance"
			return message, err
		}

		account.Wallet = account.Wallet - cost
		err = tx.Save(account).Error

		if err != nil {
			tx.Rollback()
			message = "Failed to update account balance"
			return message, err
		}

		cartitem.Status = "Cancelled"
		cartitem.CancelTrxID = transaction.ID
		cartitem.CancelledOn = time.Now().Unix()
		var ShippingDetails []TrackingDetail
		if len(cartitem.ShippingDetails) > 0 {
			json_err := json.Unmarshal([]byte(cartitem.ShippingDetails), &ShippingDetails)
			if json_err != nil {
				ShippingDetails = make([]TrackingDetail, 0)
			}
		}
		ShippingDetails = append(ShippingDetails, TrackingDetail{
			CreatedAt: time.Now().Unix(),
			Note:      trackingDetail,
		})
		ShippingDetailJson, json_err := json.Marshal(ShippingDetails)
		if json_err == nil {
			cartitem.ShippingDetails = string(ShippingDetailJson)
		}
		err = tx.Save(&cartitem).Error
		if err != nil {
			tx.Rollback()
			message = "Unable to update order item"
			return message, err
		}

		tx.Commit()
		message = "Item cancelled"

	} else if cartitem.CreatedBy == UpdatedBy.ID {
		util.Log(time.Now(), "requesting to cancel order item")
		cartitem.RequestedToCalcel = time.Now().Unix()
		var requestor = "User"
		if UpdatedBy.ID != cartitem.CreatedBy {
			requestor = "Seller"
		}
		trackingDetail := requestor + " want to cancel item. reason: " + trackingDetail
		var ShippingDetails []TrackingDetail
		if len(cartitem.ShippingDetails) > 0 {
			json_err := json.Unmarshal([]byte(cartitem.ShippingDetails), &ShippingDetails)
			if json_err != nil {
				ShippingDetails = make([]TrackingDetail, 0)
			}
		}
		ShippingDetails = append(ShippingDetails, TrackingDetail{
			CreatedAt: time.Now().Unix(),
			Note:      trackingDetail,
		})
		ShippingDetailJson, json_err := json.Marshal(ShippingDetails)
		if json_err == nil {
			cartitem.ShippingDetails = string(ShippingDetailJson)
		}
		err = db.Conn.Save(&cartitem).Error
		if err != nil {
			message = "Unable to update order item"
		} else {
			message = "Requested seller to cancel order"
		}
	} else {
		util.Log(time.Now(), "cart item created by ", cartitem.CreatedBy, "updated by ", cartitem.UpdatedBy, "request by ", UpdatedBy.ID)
	}
	return message, err
}

func (cartitem *Cartitem) Verify() (string, error) {
	var message string
	var err error
	var product Product
	err = db.Conn.First(&product, "id = ?", cartitem.ProductID).Error

	if err != nil {
		message = " One of the item in order is no longer available "
		return message, err
	}

	var stocks []Stock
	err = json.Unmarshal([]byte(product.Stock), &stocks)

	if err != nil {
		message = " Unable to process stock of one of the item in order"
		util.Log(product.Stock)
		return message, err
	}

	var stock_found bool
	for i := 0; i < len(stocks); i++ {
		stock := stocks[i]
		if stock.SKU == cartitem.SKU {
			stock_found = true

			if stock.Quantity == 0 {
				message = "One of the item in order is out of stock, please revise the order "
				return message, err
			}

			if cartitem.Quantity < stock.Quantity {
				message = "One of the item in order is having insufficient stock, please revise the order "
				return message, err
			}

			if cartitem.Cost != stock.Cost {
				message = "Price changed for one of the item in order, please revise the order "
				return message, err
			}

			if cartitem.CostTotal != stock.Cost*cartitem.Quantity {
				message = "Invalid price , malicious activity detected "
				return message, err
			}
			stock.Quantity = stock.Quantity - cartitem.Quantity
			break
		}
	}

	if !stock_found {
		message = " Stock of one of the item in order is no longer available"
		return message, err
	}
	cartitem.SelerID = product.AccountID
	err = db.Conn.Save(&cartitem).Error
	if err != nil {
		message = "fail to update cart item account id"
	}
	cartitem.product = product
	cartitem.stocks = stocks
	return message, err
}

func (cartitem *Cartitem) SetOrderID(OrderID uint, tx *gorm.DB) (string, error) {
	var err error
	var message string
	var stocks_json []byte
	stocks_json, err = json.Marshal(cartitem.stocks)
	if err != nil {
		message = "Unable to update stock in one item"
		return message, err
	}

	cartitem.product.Stock = string(stocks_json)
	err = tx.Save(&cartitem.product).Error

	if err != nil {
		message = "Unable to update product for one item"
		return message, err
	}

	cartitem.SelerID = cartitem.product.AccountID
	cartitem.OrderID = OrderID
	err = tx.Save(&cartitem).Error
	if err != nil {
		message = "Unable to update one cart item"
		return message, err
	}

	return message, err
}
