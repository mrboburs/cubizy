package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// Review will all informetion of Review
type Review struct {
	gorm.Model
	Image     string
	Review    string
	Replay    string
	Pros      string
	Cons      string
	Rating    int
	Likes     int
	Dislikes  int
	AccountID uint
	ProductID uint
	Verefied  bool
	CreatedBy uint
	UpdatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Review{})
}

// Update will update product by given post argumnets
func (review *Review) Update(ReviewMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	var old_rating int
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if UpdatedBy.ID == review.CreatedBy {
		if _, ok := ReviewMap["Rating"]; ok {
			Value := util.GetInt(ReviewMap["Rating"])
			if Value != review.Rating {
				old_rating = review.Rating
				review.Rating = Value
				flag = true
			}
		}

		if _, ok := ReviewMap["Image"]; ok {
			Value := util.GetString(ReviewMap["Image"])
			if Value != review.Image {
				review.Image = Value
				flag = true
			}
		}

		if _, ok := ReviewMap["Review"]; ok {
			Value := util.GetString(ReviewMap["Review"])
			if Value != "" && Value != review.Review {
				if review.Review != "" {
					review.Review = Value + "\n old review : \n" + review.Review
				} else {
					review.Review = Value
				}
				flag = true
			}
		}

		if _, ok := ReviewMap["Pros"]; ok {
			Value := util.GetString(ReviewMap["Pros"])
			if Value != "" && Value != review.Pros {
				review.Pros = Value
				flag = true
			}
		}

		if _, ok := ReviewMap["Cons"]; ok {
			Value := util.GetString(ReviewMap["Cons"])
			if Value != "" && Value != review.Cons {
				review.Cons = Value
				flag = true
			}
		}

		if _, ok := ReviewMap["ProductID"]; ok {
			Value := util.GetUint(ReviewMap["ProductID"])
			if Value != review.ProductID {
				review.ProductID = Value
				flag = true
			}
		}
	}

	if _, ok := ReviewMap["Replay"]; ok {
		Value := util.GetString(ReviewMap["Replay"])
		if Value != review.Replay {
			review.Replay = Value
			flag = true
		}
	}

	if flag {
		if review.ID == 0 {
			review.CreatedBy = UpdatedBy.ID
		}
		if review.AccountID == 0 {
			err = errors.New(" Seller can not be empty")
		} else if review.Review == "" {
			err = errors.New(" Review can not be empty")
		} else {
			if review.CreatedBy == UpdatedBy.ID {
				var cartitem Cartitem
				cartitem_err := db.Conn.First(&cartitem, " order_id > 0 AND product_id = ? AND created_by = ? ", review.ProductID, review.CreatedBy).Error
				if cartitem_err == nil {
					if cartitem.OrderID > 0 && cartitem.DeliveredOn > 0 {
						review.Verefied = true
					}
				}
			}
			var added bool
			review.UpdatedBy = UpdatedBy.ID
			if review.ID == 0 {
				added = true
				err = db.Conn.Create(&review).Error
			} else {
				added = true
				err = db.Conn.Save(&review).Error
			}

			if review.AccountID > 0 {
				if old_rating != review.Rating {
					if added {
						_ = db.Conn.Exec("UPDATE accounts SET rating = (SELECT SUM(reviews.rating) / COUNT(reviews.id) FROM reviews WHERE reviews.account_id = accounts.id) WHERE id = ?", review.AccountID)
						_ = db.Conn.Exec("UPDATE accounts SET reviews = (SELECT COUNT(reviews.id) FROM reviews WHERE reviews.account_id = accounts.id) WHERE id = ?", review.AccountID)
					}
					if old_rating != 0 {
						setRating("accounts", review.AccountID, old_rating)
					}
					setRating("accounts", review.AccountID, review.Rating)
				}
			}

			if review.ProductID > 0 {
				if old_rating != review.Rating {
					if added {
						_ = db.Conn.Exec("UPDATE products SET rating = (SELECT SUM(reviews.rating) / COUNT(reviews.id) FROM reviews WHERE reviews.product_id = products.id) WHERE id = ?", review.ProductID)
						_ = db.Conn.Exec("UPDATE products SET reviews = (SELECT COUNT(reviews.id) FROM reviews WHERE reviews.product_id = products.id) WHERE id = ?", review.ProductID)
					}
					if old_rating != 0 {
						setRating("products", review.ProductID, old_rating)
					}
					setRating("products", review.ProductID, review.Rating)
				}
			}
		}
	}
	return flag, err
}

func setRating(table string, id uint, rating int) {
	var table_column_name = "one_star"
	switch rating {
	case 1:
		table_column_name = "one_star"
	case 2:
		table_column_name = "two_star"
	case 3:
		table_column_name = "three_star"
	case 4:
		table_column_name = "four_star"
	case 5:
		table_column_name = "five_star"
	}
	var column_id = table
	column_id = strings.TrimSuffix(column_id, "s")
	column_id += "_id"
	_ = db.Conn.Exec("UPDATE "+table+" SET "+table_column_name+" = (SELECT COUNT(reviews.id) FROM reviews WHERE reviews."+column_id+" = "+table+".id AND reviews.rating = ? ) WHERE id = ?", rating, id).Error
}
