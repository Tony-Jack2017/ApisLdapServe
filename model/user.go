package model

type User struct {
	UID    string `gorm:"primaryKey;column:u_id"`
	Name   string `gorm:"column:name"`
	Avatar string `gorm:"column:avatar"`
	Phone  string `gorm:"column:phone"`
	Email  string `gorm:"column:email"`
}

type UserReqCreate struct {
}

type UserReqUpdate struct {
}

func (user *User) TableName() string {
	return "ldap_users"
}
