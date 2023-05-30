package db

import (
	"cubizy/util"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conn  is actual conection to mak db oprations
var Conn *gorm.DB

// Settings containg setting stored in database
type Settings struct {
	gorm.Model
	Name  string
	Value string
}

// Dblog containg log stored in database
type Dblog struct {
	gorm.Model
	Message string
	Extra   string
	Time    int64
}

func init() {
	util.Log("DB module initiating")
	defer util.Log("DB module initiated")
	util.Log(util.Settings.ConectionString)
	var err error
	Conn, err = gorm.Open(mysql.Open(util.Settings.ConectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	Conn.AutoMigrate(&Settings{})
	Conn.AutoMigrate(&Dblog{})

	Conn.Create(&Dblog{Message: "application started", Extra: "", Time: time.Now().Unix()})
}

//SetUpQuery will setup query by request paramenters
func SetUpQuery(request map[string]interface{}, query *gorm.DB) {

	var SortBy = "ID"
	var Order = "asc"
	var Limit = 0
	var Page = 0
	if _, okOrder := request["sort"]; okOrder {
		SortBy = strings.TrimSpace(request["sort"].(string))
		OrderBy := request["sortdesc"].(bool)
		if OrderBy {
			Order = "desc"
		}
		query.Order("`" + util.ToSnakeCase(SortBy) + "` " + Order)
	}

	if _, okLimit := request["limit"]; okLimit {
		Limit = util.GetInt(request["limit"])
	}

	if _, okPage := request["page"]; okPage {
		Page = util.GetInt(request["page"])
	}

	if Limit > 0 {
		query.Limit(Limit)
		if Page > 0 {
			query.Offset(Limit * Page)
		}
	}
}

/*

func test() {
	dsn := "root:Myname@123@tcp(localhost:3306)/mystudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	print("Created\n")
	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	print("Updated\n")
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	print("Updated\n")
	// Delete - delete product
	db.Delete(&product, 1)
	print("Deleted\n")
}
*/
