package model

type Account struct {
	ACID          int16  `gorm:"primaryKey;autoIncrement;column:ac_id"`
	AccountPhone  string `gorm:"unique;<-;size:32;column:account_phone"`
	AccountEmail  string `gorm:"unique;<-;size:128;column:account_email"`
	Password      string `gorm:"column:password"`
	LastLoginTime string `gorm:"column:last_login_time"`
	Location      string `gorm:"column:location"`
	LocationIP    string `gorm:"column:location_ip"`
	Device        string `gorm:"column:device"`
}

type AccountReqSignIn struct {
	Account   string `json:"account" binding:"required" msg:"account field can't be null'"`
	Password  string `json:"password" binding:"required" msg:"password field can't be null"`
	LoginType string `json:"loginType" binding:"oneof=email phone" msg:"the type of login is one of email and phone"`
}

type AccountRepSignIn struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (account *Account) TableName() string {
	return "ldap_accounts"
}
