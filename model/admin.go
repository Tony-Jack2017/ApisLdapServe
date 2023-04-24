package model

type Admin struct {
	AID         string `gorm:"primaryKey;size:128;column:a_id"`
	FirstName   string `gorm:"type:string;size:64;column:first_name"`
	LastName    string `gorm:"type:string;size:64;column:last_name"`
	FullName    string `gorm:"type:string;size:128;column:full_name"`
	Avatar      string `gorm:"column:avatar"`
	Country     string `gorm:"column:country"`
	Phone       string `gorm:"type:string;size:32;column:phone"`
	Email       string `gorm:"type:string;size:128;column:email"`
	Address     string `gorm:"column:address"`
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
