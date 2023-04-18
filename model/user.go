package model

type User struct {
	UUID  string `gorm:"primaryKey"`
	Name  string `gorm:""`
	Phone string `gorm:""`
	Email string `gorm:""`
}

type UserReqCreate struct {
}

type UserReqUpdate struct {
}

func (user *User) TableName() string {
	return "ldap_users"
}
