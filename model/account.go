package model

type Account struct {
	ACID          string `gorm:"primaryKey;column:'ac_id'"`
	Account       string `gorm:"unique;<-:create;size:32;not null"`
	Password      string `gorm:"column:password"`
	LastLoginTime string `gorm:"column:last_login_time"`
	Location      string `gorm:"column:location"`
	LocationIP    string `gorm:"column:location_ip"`
	Device        string `gorm:"column:device"`
}

type AccountReqSignIn struct {
	Account  string
	Password string
}

type AccountRepSignIn struct {
}

func (account *Account) TableName() string {
	return "ldap_accounts"
}
