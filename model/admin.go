package model

type Admin struct {
	AID         int8   `gorm:"primaryKey;"`
	FirstName   string `gorm:"size:64;column:first_name"`
	LastName    string `gorm:"size:64;column:last_name"`
	FullName    string `gorm:"size:128;column:full_name"`
	Country     string `gorm:"column:country"`
	Avatar      string `gorm:"column:avatar"`
	Phone       string `gorm:"column:phone"`
	Email       string `gorm:"column:email"`
	Company     string `gorm:"column:company"`
	Appointment string `gorm:"column:appointment"`
}

type AdminReqCreate struct {
}

type AdminReqUpdate struct {
}

func (admin *Admin) TableName() string {
	return "ldap_admins"
}
