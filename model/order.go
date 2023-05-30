package model

import (
	"cubizy/plugins/db"
	"cubizy/util"

	"gorm.io/gorm"
)

// Order will all informetion of Order
type Order struct {
	gorm.Model
	UserID               uint
	AccountID            uint
	OrderCode            uint
	Mobile               string
	AddressLine1         string
	AddressLine2         string
	AddressLine3         string
	Longitude            string
	Latitude             string
	Code                 string
	SubLocality          string
	Locality             string
	District             string
	State                string
	Country              string
	BillingAddress       string
	Price                uint
	Cost                 uint
	ShippingCost         uint
	Total                uint
	PaymentMethod        string
	PaymentMethodDetails string
	TransactionID        uint
	Status               string
	PaidOn               int64
	CreatedBy            uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Order{})
	var orders []Order
	err := db.Conn.Find(&orders).Error
	if err != nil {
		util.Log(err)
	} else {
		for _, order := range orders {
			if order.OrderCode == 0 {
				order.UserID = order.CreatedBy
				order.OrderCode = util.CantorFunction(order.ID, order.UserID*100)
				db.Conn.Save(order)
			}
		}
	}
}
