package model

import (
	"cubizy/plugins/db"

	"gorm.io/gorm"
)

// Account will all informetion of accounts
type AdminAccountView struct {
	gorm.Model
	AccountType        string
	Title              string `gorm:"unique"`
	Description        string
	Keywords           string
	Email              string `gorm:"unique"`
	Mobile             string
	AddressID          uint
	Logo               string
	WideLogo           string
	Banner             string
	Youtube            string
	Facebook           string
	Instagram          string
	Pinterest          string
	WhatsApp           string
	IDProof            string
	IDStatus           bool
	AddressProof       string
	AddressStatus      bool
	RegistretionProof  string
	RegistretionStatus bool
	OtherDocument      string
	TestAccount        bool
	Status             int
	StatusComment      string
	CreatedBy          uint
	UpdatedBy          uint
	LastActiveOn       int64
	Subjects           int
	Sessions           int
	Students           int
	Notes              int

	UpdatedByName string

	/* Website Info */
	Subdomain string
	Domain    string
	CanActive bool
	Active    bool

	ThemeColor     string
	ThemeTextColor string

	/* Owner Info */
	Name           string
	Photo          string
	OwnerEmail     string
	EmailVerified  bool
	OwnerMobile    string
	MobileVerified bool
	Wallet         uint
	Joined         bool

	Question1 string
	Answer1   string

	Question2 string
	Answer2   string

	Question3 string
	Answer3   string

	/* Main Address */
	AddressTitle  string
	AddressMobile string
	AddressLine1  string
	AddressLine2  string
	AddressLine3  string
	Longitude     string
	Latitude      string
	Code          string
	SubLocality   string
	Locality      string
	District      string
	Country       string
	LocationID    uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.Exec(`CREATE OR REPLACE VIEW admin_account_views AS
	SELECT 	accounts.*, 

			addresses.title AS address_title, 
			addresses.mobile AS address_mobile, 
			addresses.address_line1, 
			addresses.address_line2, 
			addresses.address_line3, 
			addresses.code, 
			addresses.country, 
			addresses.district, 
			addresses.locality, 
			addresses.sub_locality, 
			addresses.latitude, 
			addresses.longitude, 
			addresses.location_id,

			owner.name,
			owner.photo,
			owner.email AS owner_email,
			owner.email_verified,
			owner.mobile AS owner_mobile,
			owner.mobile_verified,
			owner.joined,

			owner.question1,
			owner.answer1,

			owner.question2,
			owner.answer2,

			owner.question3,
			owner.answer3,

			users.name AS updated_by_name
	FROM accounts 
	LEFT JOIN addresses ON accounts.address_id = addresses.id
	LEFT JOIN users AS owner ON accounts.created_by = owner.id
	LEFT JOIN users ON accounts.updated_by = users.id`)
}
