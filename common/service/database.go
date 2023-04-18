package service

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn string

func InitDatabase() {
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"127.0.0.1",
		"postgres",
		"apis1618",
		"LDAP",
		5432,
	)
}

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect to database failed!!!")
	} else {
		fmt.Println("connect to database success >>>")
	}
}
