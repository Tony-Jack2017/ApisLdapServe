package model

type User struct {
	Name  string `gorm:""`
	Phone string `gorm:""`
	Email string `gorm:""`
}
