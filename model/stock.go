package model

import (
	"cubizy/util"
	"encoding/json"
)

/*
this.Stock.push({
	SKU: "",
	Variation: {},
	Active: true,
	Quantity: 1,
	Sold: 0,
	Price: 0,
	Discount: 0,
	Cost: 0,
	LowStockLimit: 1,
	AddToCart: false,
})
*/

// Stock will all informetion of Stock
type Stock struct {
	SKU           string
	Quantity      uint
	Sold          uint
	Price         uint
	Discount      uint
	Cost          uint
	LowStockLimit uint
}

func (b *Stock) UnmarshalJSON(data []byte) error {

	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	b.SKU = util.GetString(v["SKU"])
	b.Quantity = util.GetUint(v["Quantity"])
	b.Sold = util.GetUint(v["Sold"])
	b.Price = util.GetUint(v["Price"])
	b.Discount = util.GetUint(v["Discount"])
	b.Cost = util.GetUint(v["Cost"])
	b.LowStockLimit = util.GetUint(v["LowStockLimit"])

	return nil
}
