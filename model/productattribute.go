package model

import "cubizy/plugins/db"

// Product will all informetion of Products
type Productattributes struct {
	ProductID uint `gorm:"primaryKey"`
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Productattributes{})
	db.Conn.Exec(`
	CREATE OR REPLACE VIEW productsview AS
	SELECT 
		products.*, 
		productattributes.*,
		accounts.title as account_name,
		accounts.subdomain,		
		accounts.domain
	FROM products 
	LEFT JOIN productattributes on products.id = productattributes.product_id
	LEFT JOIN accounts on products.account_id = accounts.id 
	WHERE products.deleted_at IS NULL AND 
		  products.status = 1`)
	db.Conn.Exec(`
		  CREATE OR REPLACE VIEW accountproductsview AS
		  SELECT 
			  products.*, 
			  productattributes.*
		  FROM products 
		  LEFT JOIN productattributes on products.id = productattributes.product_id
		  WHERE products.deleted_at IS NULL AND 
				products.status = 1`)
}
