package model

import "ApisLdapServe/common/service"

type Account struct {
	ACID          int16  `gorm:"primaryKey;autoIncrement;column:ac_id"`
	Password      string `gorm:"column:password"`
	AccountPhone  string `gorm:"unique;<-;size:32;column:account_phone"`
	AccountEmail  string `gorm:"unique;<-;size:128;column:account_email"`
	PhoneStatus   uint   `gorm:"column:phone_status"`
	EmailStatus   uint   `gorm:"column:email_status"`
	LastLoginTime string `gorm:"column:last_login_time"`
	Location      string `gorm:"column:location"`
	LocationIP    string `gorm:"column:location_ip"`
	Device        string `gorm:"column:device"`
	AdminID       string `gorm:"column:admin_id"`
	Admin         Admin  `gorm:"foreignKey:AdminID"`
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

func AccountCheckSQL(id int16, email string, phone string) (error, *Account) {
	var account Account
	if id != 0 {
		if err := service.DB.Table("ldap_accounts").Where("id = ?", id).First(&account).Error; err != nil {
			return err, nil
		}
	}
	if email != "" {
		if err := service.DB.Table("ldap_accounts").Where("email = ?", email).First(&account).Error; err != nil {
			return err, nil
		}
	}
	if phone != "" {
		if err := service.DB.Table("ldap_accounts").Where("phone = ?", phone).First(&account).Error; err != nil {
			return err, nil
		}
	}
	return nil, nil
}

func AccountAddSQL() {

}

func AccountModifySQL() {

}

func AccountSearchSQL() {}
