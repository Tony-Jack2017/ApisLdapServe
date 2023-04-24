package service

import (
	"ApisLdapServe/common/base"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn string

func InitDatabase() {
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		base.Conf.Database.Host,
		base.Conf.Database.User,
		base.Conf.Database.Password,
		base.Conf.Database.DBName,
		base.Conf.Database.Port,
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
