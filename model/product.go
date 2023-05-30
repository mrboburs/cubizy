package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Product will all informetion of Products
type Product struct {
	gorm.Model
	Logo        string
	Name        string
	Summary     string
	Brand       string `gorm:"default:''"`
	Condition   string `gorm:"default:''"`
	QuickPoints string `gorm:"default:''"`
	Comment     string
	Keywords    string
	ImageBy     string
	Images      string
	Options     string

	Status        bool
	Service       bool
	Perhour       bool
	HaveVariation bool

	CategoryID      uint
	SubcategoryID   uint
	ChildcategoryID uint

	BaseCategoryID      uint
	BaseSubcategoryID   uint
	BaseChildcategoryID uint

	Shipping string

	EVRating uint

	Trending  int
	Rating    int
	Reviews   int
	OneStar   int
	TwoStar   int
	ThreeStar int
	FourStar  int
	FiveStar  int

	Variation string
	Stock     string

	SKU           string
	Quantity      uint
	Price         uint
	Discount      uint
	Cost          uint
	LowStockLimit uint

	MaxPrice uint
	MaxCost  uint

	Sold    uint
	Revenue uint

	AccountID uint
	CreatedBy uint
	UpdatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Product{})
	db.Conn.Exec("UPDATE products SET products.base_category_id = (SELECT categories.base_category_id FROM categories WHERE categories.id = products.category_id)")
	db.Conn.Exec("UPDATE products SET products.base_subcategory_id = (SELECT subcategories.base_subcategory_id FROM subcategories WHERE subcategories.id = products.subcategory_id)")
	db.Conn.Exec("UPDATE products SET products.base_childcategory_id = (SELECT childcategories.base_childcategory_id FROM childcategories WHERE childcategories.id = products.childcategory_id)")
	SetTrending(time.Now())
	go util.DoEvery(12*time.Hour, SetTrending)
}

func SetTrending(t time.Time) {
	var two_month_befor = t.AddDate(0, -2, 0)
	err := db.Conn.Exec("UPDATE products SET products.trending = (SELECT count(cartitems.id) FROM cartitems WHERE cartitems.product_id = products.id AND cartitems.order_id > 0 AND cartitems.created_at > ?)", two_month_befor).Error
	if err != nil {
		util.Log("Unable to set trending points", err.Error())
	}
}

// Update will update product by given post argumnets
func (product *Product) Update(ProductMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool

	var category_flag bool

	var category Category
	var subcategory Subcategory
	var childcategory Childcategory

	old_category_id := product.CategoryID
	old_subcategory_id := product.SubcategoryID
	old_childcategory_id := product.ChildcategoryID

	old_basecategory_id := product.BaseCategoryID
	old_basesubcategory_id := product.BaseSubcategoryID
	old_basechildcategory_id := product.BaseChildcategoryID

	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := ProductMap["Logo"]; ok {
		Value := util.GetString(ProductMap["Logo"])
		if Value != product.Logo {
			product.Logo = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Name"]; ok {
		Value := util.GetString(ProductMap["Name"])
		if Value != product.Name {
			product.Name = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Summary"]; ok {
		Value := util.GetString(ProductMap["Summary"])
		if Value != product.Summary {
			product.Summary = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Brand"]; ok {
		Value := util.GetString(ProductMap["Brand"])
		if Value != product.Brand {
			product.Brand = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Condition"]; ok {
		Value := util.GetString(ProductMap["Condition"])
		if Value != product.Condition {
			product.Condition = Value
			flag = true
		}
	}

	if _, ok := ProductMap["QuickPoints"]; ok {
		Value := util.GetString(ProductMap["QuickPoints"])
		if Value != product.QuickPoints {
			product.QuickPoints = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Comment"]; ok {
		Value := util.GetString(ProductMap["Comment"])
		if Value != product.Comment {
			product.Comment = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Keywords"]; ok {
		Value := util.GetString(ProductMap["Keywords"])
		if Value != product.Keywords {
			product.Keywords = Value
			flag = true
		}
	}

	if _, ok := ProductMap["ImageBy"]; ok {
		Value := util.GetString(ProductMap["ImageBy"])
		if Value != product.ImageBy {
			product.ImageBy = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Images"]; ok {
		Value := util.GetString(ProductMap["Images"])
		if Value != product.Images {
			product.Images = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Options"]; ok {
		Value := util.GetString(ProductMap["Options"])
		if Value != product.Options {
			product.Options = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Shipping"]; ok {
		Value := util.GetString(ProductMap["Shipping"])
		if Value != product.Shipping {
			product.Shipping = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Variation"]; ok {
		Value := util.GetString(ProductMap["Variation"])
		if Value != product.Variation {
			product.Variation = Value
			flag = true
		}
	}

	if _, ok := ProductMap["SKU"]; ok {
		Value := util.GetString(ProductMap["SKU"])
		if Value != product.SKU {
			product.SKU = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Status"]; ok {
		Value := util.GetBool(ProductMap["Status"])
		if Value != product.Status {
			product.Status = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Service"]; ok {
		Value := util.GetBool(ProductMap["Service"])
		if Value != product.Service {
			product.Service = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Perhour"]; ok {
		Value := util.GetBool(ProductMap["Perhour"])
		if Value != product.Perhour {
			product.Perhour = Value
			flag = true
		}
	}

	if _, ok := ProductMap["HaveVariation"]; ok {
		Value := util.GetBool(ProductMap["HaveVariation"])
		if Value != product.HaveVariation {
			product.HaveVariation = Value
			flag = true
		}
	}

	if _, ok := ProductMap["CategoryID"]; ok {
		Value := util.GetUint(ProductMap["CategoryID"])
		if Value != product.CategoryID {
			product.CategoryID = Value
			flag = true
			category_flag = true
		}
	}

	if category_flag {
		if product.CategoryID > 0 {
			category_err := db.Conn.First(&category, "id = ?", product.CategoryID).Error
			if category_err == nil {
				product.BaseCategoryID = category.BaseCategoryID
			}
		} else {
			product.BaseCategoryID = 0
		}
	}

	if _, ok := ProductMap["SubcategoryID"]; ok {
		Value := util.GetUint(ProductMap["SubcategoryID"])
		if Value != product.SubcategoryID {
			product.SubcategoryID = Value
			flag = true
			category_flag = true
		}
	}

	if category_flag {
		if product.SubcategoryID > 0 {
			subcategory_err := db.Conn.First(&subcategory, "id = ?", product.SubcategoryID).Error
			if subcategory_err == nil {
				product.BaseSubcategoryID = subcategory.BaseSubcategoryID
				product.BaseChildcategoryID = subcategory.BaseChildcategoryID
			}
		} else {
			product.BaseSubcategoryID = 0
		}
	}

	if _, ok := ProductMap["ChildcategoryID"]; ok {
		Value := util.GetUint(ProductMap["ChildcategoryID"])
		if Value != product.ChildcategoryID {
			product.ChildcategoryID = Value
			flag = true
			category_flag = true
		}
	}

	if category_flag {
		if product.ChildcategoryID > 0 {
			childcategory_err := db.Conn.First(&childcategory, "id = ?", product.ChildcategoryID).Error
			if childcategory_err == nil {
				product.BaseChildcategoryID = childcategory.BaseChildcategoryID
			}
		} else {
			product.BaseChildcategoryID = 0
		}
	}

	if _, ok := ProductMap["EVRating"]; ok {
		Value := util.GetUint(ProductMap["EVRating"])
		if Value != product.EVRating {
			product.EVRating = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Stock"]; ok {
		Value := util.GetString(ProductMap["Stock"])
		if Value != product.Stock {
			product.Stock = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Quantity"]; ok {
		Value := util.GetUint(ProductMap["Quantity"])
		if Value != product.Quantity {
			product.Quantity = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Price"]; ok {
		Value := util.GetUint(ProductMap["Price"])
		if Value != product.Price {
			product.Price = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Discount"]; ok {
		Value := util.GetUint(ProductMap["Discount"])
		if Value != product.Discount {
			product.Discount = Value
			flag = true
		}
	}

	if _, ok := ProductMap["Cost"]; ok {
		Value := util.GetUint(ProductMap["Cost"])
		if Value != product.Cost {
			product.Cost = Value
			flag = true
		}
	}

	if _, ok := ProductMap["MaxPrice"]; ok {
		Value := util.GetUint(ProductMap["MaxPrice"])
		if Value != product.MaxPrice {
			product.MaxPrice = Value
			flag = true
		}
	}
	if product.MaxPrice < product.Price {
		product.MaxPrice = product.Price
	}

	if _, ok := ProductMap["MaxCost"]; ok {
		Value := util.GetUint(ProductMap["MaxCost"])
		if Value != product.MaxCost {
			product.MaxCost = Value
			flag = true
		}
	}
	if product.MaxCost < product.Cost {
		product.MaxCost = product.Cost
	}

	if _, ok := ProductMap["LowStockLimit"]; ok {
		Value := util.GetUint(ProductMap["LowStockLimit"])
		if Value != product.LowStockLimit {
			product.LowStockLimit = Value
			flag = true
		}
	}
	if _, ok := ProductMap["ExtraDetails"]; ok {
		flag = true
	}

	if flag {
		if product.Name == "" {
			err = errors.New(" Title can not be empty")
		} else if product.CategoryID == 0 || product.SubcategoryID == 0 {
			err = errors.New(" Product must have Category and Subcategory")
		} else if product.HaveVariation && product.Variation == "" {
			err = errors.New(" If product have variations please add few")
		} else {
			product.UpdatedBy = UpdatedBy.ID
			if product.ID == 0 {
				product.CreatedBy = UpdatedBy.ID
				var duplicates Product
				err = db.Conn.First(&duplicates, " `name` = ? && `account_id` = ? ", product.Name, product.AccountID).Error
				if err == nil {
					err = errors.New("duplicate name")
				} else {
					err = db.Conn.Create(&product).Error
				}
			} else {
				err = db.Conn.Save(&product).Error
			}

			if err == nil {
				if old_category_id != product.CategoryID {
					db.Conn.Exec("UPDATE categories SET categories.products = (SELECT count(products.id) FROM products WHERE products.category_id = categories.id) WHERE categories.id = ?", old_category_id)
					db.Conn.Exec("UPDATE categories SET categories.products = (SELECT count(products.id) FROM products WHERE products.category_id = categories.id) WHERE categories.id = ?", product.CategoryID)
				}

				if old_subcategory_id != product.SubcategoryID {
					db.Conn.Exec("UPDATE subcategories SET subcategories.products = (SELECT count(products.id) FROM products WHERE products.subcategory_id = subcategories.id) WHERE subcategories.id = ?", old_subcategory_id)
					db.Conn.Exec("UPDATE subcategories SET subcategories.products = (SELECT count(products.id) FROM products WHERE products.subcategory_id = subcategories.id) WHERE subcategories.id = ?", product.SubcategoryID)
				}

				if old_childcategory_id != product.ChildcategoryID {
					db.Conn.Exec("UPDATE childcategories SET childcategories.products = (SELECT count(products.id) FROM products WHERE products.childcategory_id = childcategories.id) WHERE childcategories.id = ?", old_childcategory_id)
					db.Conn.Exec("UPDATE childcategories SET childcategories.products = (SELECT count(products.id) FROM products WHERE products.childcategory_id = childcategories.id) WHERE childcategories.id = ?", product.ChildcategoryID)
				}

				if old_basecategory_id != product.BaseCategoryID {
					db.Conn.Exec("UPDATE categories SET categories.products = (SELECT count(products.id) FROM products WHERE products.base_category_id = categories.id) WHERE categories.id = ?", old_basecategory_id)
					db.Conn.Exec("UPDATE categories SET categories.products = (SELECT count(products.id) FROM products WHERE products.base_category_id = categories.id) WHERE categories.id = ?", product.BaseCategoryID)
				}

				if old_basesubcategory_id != product.BaseSubcategoryID {
					db.Conn.Exec("UPDATE subcategories SET subcategories.products = (SELECT count(products.id) FROM products WHERE products.base_subcategory_id = subcategories.id) WHERE subcategories.id = ?", old_basesubcategory_id)
					db.Conn.Exec("UPDATE subcategories SET subcategories.products = (SELECT count(products.id) FROM products WHERE products.base_subcategory_id = subcategories.id) WHERE subcategories.id = ?", product.BaseSubcategoryID)
				}

				if old_basechildcategory_id != product.BaseChildcategoryID {
					db.Conn.Exec("UPDATE childcategories SET childcategories.products = (SELECT count(products.id) FROM products WHERE products.base_childcategory_id = childcategories.id) WHERE childcategories.id = ?", old_basechildcategory_id)
					db.Conn.Exec("UPDATE childcategories SET childcategories.products = (SELECT count(products.id) FROM products WHERE products.base_childcategory_id = childcategories.id) WHERE childcategories.id = ?", product.BaseChildcategoryID)
				}
			}
			if _, ok := ProductMap["ExtraDetails"]; err == nil && ok {
				util.Log("adding ExtraDetails")
				ExtraDetails := ProductMap["ExtraDetails"].(map[string]interface{})
				if len(ExtraDetails) > 0 {
					var extra_details_flag bool
					var results []map[string]interface{}
					err = db.Conn.Table("productattributes").Find(&results, " product_id = ?", product.ID).Error
					if err != nil || len(results) == 0 {
						productattributes := Productattributes{
							ProductID: product.ID,
						}
						err = db.Conn.Create(&productattributes).Error
						if err == nil {
							extra_details_flag = true
						} else {
							util.Log("creating ExtraDetails", err)
						}
					} else {
						var result = results[0]
						for Column, Value := range ExtraDetails {
							if Value != result[Column] {
								extra_details_flag = true
							}
						}
					}

					if extra_details_flag {
						err = db.Conn.Table("productattributes").Where(" product_id = ?", product.ID).Updates(ExtraDetails).Error
						if err == nil {
							util.Log("Updated product extra details for ", product.ID)
						} else {
							util.Log("Unable to update product extra details for ", product.ID)
						}
					} else {
						util.Log("extra_details_flag is false")
					}
				} else {
					util.Log("no ExtraDetails")
				}
			}
		}
	}
	return flag, err
}
