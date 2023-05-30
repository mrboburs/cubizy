package userapi

import (
	"cubizy/apiresponse"
	"cubizy/model"
	"cubizy/plugins/db"
	"cubizy/util"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var orderAPI = func(w http.ResponseWriter, r *http.Request, response *apiresponse.Response) error {
	var err error
	util.Log(time.Now(), "Placing Order for user", response.User.ID, " account ", response.Account.ID)
	if _, ok := response.Request["order"]; ok {
		postorder := response.Request["order"].(map[string]interface{})
		var order model.Order
		var transactions []model.Transaction
		var items []model.Cartitem
		var order_code uint
		if _, ok := postorder["OrderCode"]; ok {
			order_code = util.GetUint(postorder["OrderCode"])
		}
		if order_code > 0 {
			err = db.Conn.First(&order, "order_code = ?", order_code).Error
			if err != nil {
				response.Result["tab"] = "review"
				response.Message = "Failed to find order"
				return err
			}

			err = db.Conn.Find(&items, "order_id = ?", order.ID).Error
			if err != nil {
				response.Result["tab"] = "review"
				response.Message = "Failed to find order items"
				return err
			}

			if _, ok := response.Request["item_to_cancel"]; ok {
				item_to_cancel := util.GetString(response.Request["item_to_cancel"])
				product_to_cancel := util.GetUint(response.Request["product_to_cancel"])
				trackingDetail := util.GetString(response.Request["TrackingDetail"])
				var message = ""
				for item_index, item := range items {
					if item_to_cancel == "All" || (item_to_cancel == item.SKU && product_to_cancel == item.ProductID) {
						if item.CancelledOn == 0 {
							message, _ = item.Cancel(trackingDetail, response.User)
							response.Message += " item " + strconv.Itoa(item_index+1) + " : " + message
						} else {
							response.Message += " item " + strconv.Itoa(item_index+1) + " allready cancled"
						}
					}
				}
			}

			response.Result["items"] = items
		} else {
			order = model.Order{
				AccountID: model.SuperAdminAccount.ID,
				UserID:    response.User.ID,
				CreatedBy: response.User.ID,
			}
			if _, ok := postorder["Items"]; ok {

				util.Log(time.Now(), " geting order items passed in post")
				var order_items []model.Cartitem
				postItems := postorder["Items"].([]interface{})
				util.Log("Got ", len(postItems), " items in order request")
				for postItemIndex, _postItem := range postItems {
					util.Log(time.Now(), " geting order item at index", postItemIndex)
					postItem := _postItem.(map[string]interface{})
					var item model.Cartitem
					_, err = item.Update(postItem, response.User)
					util.Log(time.Now(), " updated order item at index", postItemIndex)

					if err != nil {
						response.Result["tab"] = "cart"
						response.Message = "Failed to process items in order : " + err.Error()
						return err
					}

					util.Log(time.Now(), " varifying order item at index", postItemIndex)
					response.Message, err = item.Verify()
					if err != nil {
						response.Result["tab"] = "cart"
						return err
					}
					util.Log(time.Now(), " veriefied order item at index", postItemIndex)
					order_items = append(order_items, item)
				}
				util.Log(time.Now(), " veriefied all order items")

				if _, ok := postorder["ShippingAddress"]; ok {
					var addressID = util.GetUint(postorder["ShippingAddress"])
					var address model.Address
					err = db.Conn.First(&address, "id = ?", addressID).Error

					if err != nil {
						response.Result["tab"] = "addresses"
						response.Message = "Shipping address Not fond"
						return err
					}

					order.Mobile = address.Mobile
					order.AddressLine1 = address.AddressLine1
					order.AddressLine2 = address.AddressLine2
					order.AddressLine3 = address.AddressLine3
					order.Longitude = address.Longitude
					order.Latitude = address.Latitude
					order.Code = address.Code
					order.SubLocality = address.SubLocality
					order.Locality = address.Locality
					order.District = address.District
					order.State = address.State
					order.Country = address.Country

				} else {
					response.Message = "Shipping address is not provided"
					err = errors.New("shipping address is not provided")
					return err
				}
				util.Log(time.Now(), "got shipping address , checking billing address")
				if _, ok := postorder["BillingAddress"]; ok {
					var addressID = util.GetUint(postorder["BillingAddress"])
					var address model.Address
					err = db.Conn.First(&address, "id = ?", addressID).Error
					if err != nil {
						response.Result["tab"] = "addresses"
						response.Message = "Billing address Not fond"
						return err
					}
					var billingAddress []byte
					billingAddress, err = json.Marshal(address)
					if err != nil {
						response.Result["tab"] = "addresses"
						response.Message = "Billing address is invalid"
						return err
					}
					order.BillingAddress = string(billingAddress)
				} else {
					response.Message = "Billing address is not provided"
					err = errors.New("billing address is not provided")
					return err
				}
				util.Log(time.Now(), "got billing address , geting other order details")

				if _, ok := postorder["Price"]; ok {
					Value := util.GetUint(postorder["Price"])
					if Value != order.Price {
						order.Price = Value
					}
				}

				if _, ok := postorder["Cost"]; ok {
					Value := util.GetUint(postorder["Cost"])
					if Value != order.Cost {
						order.Cost = Value
					}
				}

				if _, ok := postorder["ShippingCost"]; ok {
					Value := util.GetUint(postorder["ShippingCost"])
					if Value != order.ShippingCost {
						order.ShippingCost = Value
					}
				}

				if _, ok := postorder["Total"]; ok {
					Value := util.GetUint(postorder["Total"])
					if Value != order.Total {
						order.Total = Value
					}
				}

				var TotalCost uint = 0
				var TotalShippingCost uint = 0

				util.Log(time.Now(), "summing serverside total cost")
				for _, order_item := range order_items {
					TotalCost += order_item.CostTotal
					TotalShippingCost += order_item.ShippingCost
				}

				var Total uint = TotalCost + TotalShippingCost

				util.Log(time.Now(), "validating serverside total cost with total cost in request")
				if Total != order.Total {
					response.Result["tab"] = "review"
					response.Message = "Invalid total cost , malicious activity detected "
					return err
				}

				order.Status = "Order Placed"

				tx := db.Conn.Begin()
				util.Log(time.Now(), "transaction started")

				err = tx.Create(&order).Error

				if err != nil {
					tx.Rollback()
					response.Result["tab"] = "review"
					response.Message = "Unable to save order : " + err.Error()
					return err
				}
				util.Log(time.Now(), "order created", order.ID)

				order.OrderCode = util.CantorFunction(order.ID, order.UserID*100)

				_ = db.Conn.Save(order).Error
				for _, order_item := range order_items {
					response.Message, err = order_item.SetOrderID(order.ID, tx)
					if response.Message != "" || err != nil {
						tx.Rollback()
					}
				}

				tx.Commit()
				util.Log(time.Now(), "transaction ended")
			}
		}

		err = db.Conn.Find(&transactions, " order_id = ? ", order.ID).Error

		if order.PaidOn == 0 {
			util.Log(time.Now(), "checking if payment is allreaady done for order id", order.ID)
			for i := 0; i < len(transactions); i++ {
				var transaction = transactions[i]
				if transaction.For == "Order" && transaction.Accepted {
					order.TransactionID = transaction.ID
					order.PaidOn = time.Now().Unix()
					order.Status = "Paid"
					err = db.Conn.Save(order).Error
					if err != nil {
						response.Result["tab"] = "review"
						response.Message = "Failed to set transaction_id for order"
						return err
					} else {
						response.Result["tab"] = "thankyou"
						response.Message = "Payment for this order is received successfully."
					}
					break
				}
			}
		}

		if order.PaidOn == 0 {
			util.Log(time.Now(), "getting payment details in request")

			if _, ok := postorder["PaymentMethod"]; ok {
				Value := util.GetString(postorder["PaymentMethod"])
				if Value != order.PaymentMethod {
					order.PaymentMethod = Value
				}
			}

			if _, ok := postorder["PaymentMethodDetails"]; ok {
				Value := util.GetString(postorder["PaymentMethodDetails"])
				if Value != order.PaymentMethodDetails {
					order.PaymentMethodDetails = Value
				}
			}

			response.User.UpdateOrderCount()
			if order.PaymentMethod == "Wallet" {
				var transaction model.Transaction

				if order.Total > response.User.Wallet {
					response.Result["tab"] = "payment"
					response.Message = "User Wallet don't have sufficient balance, Order is placed and you can try again to complete payment"
					return err
				}

				transaction = model.Transaction{
					AccountID:     response.Account.ID,
					UserID:        response.User.ID,
					CreatedBy:     response.User.ID,
					UpdatedBy:     response.User.ID,
					Amount:        order.Total,
					Method:        order.PaymentMethod,
					TransactionID: uuid.New().String(),
					For:           "Order",
					OrderID:       order.ID,
					Status:        "successful",
					Accepted:      true,
				}
				payment_tx := db.Conn.Begin()
				err = payment_tx.Create(&transaction).Error

				if err != nil {
					payment_tx.Rollback()
					response.Result["tab"] = "review"
					response.Message = "Failed to complete payment transaction, Order is placed and you can try again to complete payment"
					return err
				}

				response.User.Wallet = response.User.Wallet - order.Total
				err = payment_tx.Save(response.User).Error

				if err != nil {
					payment_tx.Rollback()
					response.Result["tab"] = "review"
					response.Message = "Failed to update user balance, Order is placed and you can try again to complete payment"
					return err
				}

				response.Account.Wallet = response.Account.Wallet + order.Total
				err = payment_tx.Save(response.Account).Error

				if err != nil {
					payment_tx.Rollback()
					response.Result["tab"] = "review"
					response.Message = "Failed to update account balance, Order is placed and you can try again to complete payment"
					return err
				}

				order.TransactionID = transaction.ID
				order.PaidOn = time.Now().Unix()
				order.Status = "Paid"
				err = payment_tx.Save(order).Error

				if err != nil {
					payment_tx.Rollback()
					response.Result["tab"] = "review"
					response.Message = "Failed to set paid status on order, But payment is done, Report the issue with order details, we will resolve the issue asap"
					return err
				} else {
					response.Result["tab"] = "thankyou"
					if order.Status == "Paid" {
						response.Message = "Order placed and payment received successfully."
					} else {
						response.Message = "Order placed , please complete the payment"
					}
				}
				payment_tx.Commit()
				_ = db.Conn.Find(&transactions, " order_id = ? ", order.ID).Error
			}
		}
		if order.ID > 0 {
			response.Status = apiresponse.SUCCESS
			response.Result["order"] = order
			response.Result["transactions"] = transactions
		}
	} else {
		util.Log(time.Now(), "order object not present in post")
		response.Message = "Empty request "
	}
	return err
}
