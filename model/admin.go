package model

type Admin struct {
	AID         string `gorm:"primaryKey;size:128;column:a_id" json:"a_id"`
	FirstName   string `gorm:"type:string;size:64;column:first_name" json:"firstName"`
	LastName    string `gorm:"type:string;size:64;column:last_name" json:"lastName"`
	FullName    string `gorm:"type:string;size:128;column:full_name" json:"fullName"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
	Country     string `gorm:"column:country" json:"country"`
	Phone       string `gorm:"type:string;size:32;column:phone" json:"phone"`
	Email       string `gorm:"type:string;size:128;column:email" json:"email"`
	Address     string `gorm:"column:address" json:"address"`
	Company     string `gorm:"column:company" json:"company"`
	Appointment string `gorm:"column:appointment" json:"appointment"`
}

type AdminReqCreate struct {
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	FullName    string `json:"fullName" binding:"required"`
	Avatar      string `json:"avatar" binding:"url"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	Email       string `json:"email" binding:"email"`
	Address     string `json:"address"`
	Company     string `json:"company"`
	Appointment string `json:"appointment"`
}

type AdminReqUpdate struct {
}

func (admin *Admin) TableName() string {
	return "ldap_admins"
}

func AdminAddSQL() error {
	return nil
}

func AdminUpdateSQL() error {
	return nil
}