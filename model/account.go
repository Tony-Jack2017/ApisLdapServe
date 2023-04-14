package model

type Account struct {
	ACID          string `gorm:"primaryKey"`
	Account       string `gorm:""`
	Password      string `gorm:""`
	LastLoginTime string `gorm:""`
	Location      string `gorm:""`
}

type AccountReqSignIn struct {
	Account  string
	Password string
}

type AccountRepSignIn struct {
}
